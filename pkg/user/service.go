// Package user contains all functions and methods that impl user business logic
package user

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/pkg/password"
	"github.com/Akshit8/reddit-clone-api/pkg/token"
	"github.com/lib/pq"
	"github.com/reiver/go-pqerror"
)

const tokenDuration = 1 * time.Hour

// Service defines functions available on entity user
type Service interface {
	RegisterUser(ctx context.Context, username, password, email string) (entity.User, error)
	LoginUser(ctx context.Context, usernameOrEmail, password string) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userService struct {
	repo       *db.Queries
	hasher     password.Hasher
	tokenMaker token.Maker
}

// NewUserService creates new instance of postService
func NewUserService(repo *db.Queries, tokenMaker token.Maker) Service {
	return &userService{
		repo:       repo,
		hasher:     password.NewNativeHasher(),
		tokenMaker: tokenMaker,
	}
}

func (u *userService) RegisterUser(ctx context.Context, username, password, email string) (entity.User, error) {
	hashedPassword, err := u.hasher.HashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	createTimestamp := time.Now()

	newUser := db.CreateUserParams{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: createTimestamp,
		UpdatedAt: createTimestamp,
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
	if strings.Contains(usernameOrEmail,  "@") {
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
