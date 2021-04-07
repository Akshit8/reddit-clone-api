// Package db impls all function and queries to interact with DB.
package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/reiver/go-pqerror"
)

// Store is implemented to extend transaction functionalities on Queries
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates new instance of Store.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// UpvoteTxParams defines params required for upvote transaction
type UpvoteTxParams struct {
	UserID  int64
	PostID  int64
	Value   int32
}

// UpvoteTx runs a transcation to create upvote on a post
func (s *Store) UpvoteTx(ctx context.Context, arg UpvoteTxParams) error {
	err := s.execTx(ctx, func(q *Queries) error {
		_, err := s.CreateUpvote(ctx, CreateUpvoteParams{
			UserId: arg.UserID,
			PostId: arg.PostID,
			Value:  arg.Value,
		})
		if err != nil {
			pqError := err.(*pq.Error)
			switch pqError.Code {
			case pqerror.CodeIntegrityConstraintViolationUniqueViolation:
				return errors.New("unique violation")
			default:
				return err
			}
		}

		err = s.UpdatePostUpvotes(ctx, UpdatePostUpvotesParams{
			ID:      arg.PostID,
			Upvotes: int64(arg.Value),
		})
		if err != nil {
			return err
		}
		
		return nil
	})

	return err
}
