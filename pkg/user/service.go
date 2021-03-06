// Package user contains all functions and methods that impl user business logic
package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/pkg/mail"
	"github.com/Akshit8/reddit-clone-api/pkg/password"
	"github.com/Akshit8/reddit-clone-api/pkg/redis"
	"github.com/Akshit8/reddit-clone-api/pkg/token"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/reiver/go-pqerror"
)

const (
	tokenDuration          = 24 * time.Hour
	forgotPasswordPrefix   = "forgot_password"
	forgotPasswordDuration = time.Duration(30 * time.Minute)
)

// Service defines functions available on entity user
type Service interface {
	RegisterUser(ctx context.Context, username, password, email string) (entity.User, error)
	LoginUser(ctx context.Context, usernameOrEmail, password string) (string, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	ForgotPassword(ctx context.Context, email string) (bool, error)
	ChangePassword(ctx context.Context, token, newPassword string) (string, error)
}

type userService struct {
	repo       *db.Store
	hasher     password.Hasher
	tokenMaker token.Maker
	redis      redis.CacheOperations
	mailer     mail.GoMailer
}

// NewUserService creates new instance of postService
func NewUserService(
	repo *db.Store,
	tokenMaker token.Maker,
	hasher password.Hasher,
	redis redis.CacheOperations,
	mailer mail.GoMailer,
) Service {
	return &userService{
		repo:       repo,
		hasher:     hasher,
		tokenMaker: tokenMaker,
		redis:      redis,
		mailer:     mailer,
	}
}

func (u *userService) RegisterUser(ctx context.Context, username, password, email string) (entity.User, error) {
	hashedPassword, err := u.hasher.HashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	newUser := db.CreateUserParams{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
	}

	user, err := u.repo.CreateUser(ctx, newUser)
	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code == pqerror.CodeIntegrityConstraintViolationUniqueViolation {
			return entity.User{}, errors.New("username/email already exists")
		}
		return entity.User{}, err
	}

	result := entity.User{
		ID:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return result, nil
}

func (u *userService) LoginUser(ctx context.Context, usernameOrEmail, password string) (string, error) {
	var user db.User
	var err error
	if strings.Contains(usernameOrEmail, "@") {
		user, err = u.repo.GetUserByEmail(ctx, usernameOrEmail)
	} else {
		user, err = u.repo.GetUserByUsername(ctx, usernameOrEmail)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("username/email doesn't exists")
		}
		return "", err
	}

	err = u.hasher.CheckPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	accessToken, err := u.tokenMaker.CreateToken(user.Username, tokenDuration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user, err := u.repo.GetUserByID(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, err
	}

	result := entity.User{
		ID:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return result, nil
}

func (u *userService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username doesn't exists")
		}
		return nil, err
	}

	result := &entity.User{
		ID:        int(user.ID),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return result, nil
}

func (u *userService) ForgotPassword(ctx context.Context, email string) (bool, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return false, err
	}

	err = u.redis.Set(fmt.Sprintf("%s:%s", forgotPasswordPrefix, token.String()), user.ID, forgotPasswordDuration)
	if err != nil {
		return false, err
	}

	mailBody := fmt.Sprintf(`<a href="http://localhost:3000/password-reset/%s">reset-password</a>"`, token)
	// mail sending is done inside a new goroutine.
	go u.mailer.SendMail([]string{user.Email}, mailBody)

	return true, nil
}

func (u *userService) ChangePassword(ctx context.Context, token, newPassword string) (string, error) {
	key := fmt.Sprintf("%s:%s", forgotPasswordPrefix, token)
	userID, err := u.redis.GetString(key)
	if err != nil {
		return "", err
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return "", err
	}

	hashedPassword, err := u.hasher.HashPassword(newPassword)
	if err != nil {
		return "", err
	}

	user, err := u.repo.GetUserByID(ctx, int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user no longer exists")
		}
	}

	err = u.repo.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		ID: user.ID,
		Password: hashedPassword,
	})
	if err != nil {
		return "", errors.New("update failed")
	}

	err = u.redis.Delete(key)
	if err != nil {
		return "", errors.New("ISR")
	}

	return u.tokenMaker.CreateToken(user.Username, tokenDuration)
}
