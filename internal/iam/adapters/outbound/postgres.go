package outbound

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/internal/iam/domain/iam"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type PostgresAdapter struct {
	Handler IQuerierIam
}

func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (a PostgresAdapter) CreateUser(ctx context.Context, user *iam.User) error {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.create_user")
	defer span.End()

	_, err := a.Handler.CreateUser(ctx, CreateUserParams{
		ID:           pgtype.UUID{Bytes: user.ID, Valid: true},
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    pgtype.Timestamp{Time: user.CreatedAt, Valid: true},
		LastLogin:    pgtype.Timestamp{Time: time.Time{}, Valid: false},
	})

	return err
}

func (p PostgresAdapter) UpdateUser(ctx context.Context, userId string, updateFn func(*iam.User) (*iam.User, error)) error {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.update_user")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	user, err := p.GetUser(ctx, userId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedUser, err := updateFn(user)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	_, err = p.Handler.UpdateUser(ctx, UpdateUserParams{
		ID:           pgtype.UUID{Bytes: updatedUser.ID, Valid: true},
		Email:        updatedUser.Email,
		PasswordHash: updatedUser.PasswordHash,
		LastLogin:    pgtype.Timestamp{Time: time.Time{}, Valid: false},
		CreatedAt:    pgtype.Timestamp{Time: updatedUser.CreatedAt, Valid: true},
	})

	return err
}

func (p PostgresAdapter) FindUserByEmail(ctx context.Context, email string) (IamUser, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.find_user_by_email")
	defer span.End()

	return p.Handler.FindUserByEmail(ctx, FindUserByEmailParams{
		Email: email,
	})
}

func (p PostgresAdapter) FindUserByID(ctx context.Context, id string) (IamUser, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.find_user_by_id")
	defer span.End()

	return p.Handler.FindUserByID(ctx, FindUserByIDParams{
		ID: pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true},
	})
}

func (p PostgresAdapter) GetUser(ctx context.Context, id string) (*iam.User, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.get_user")
	defer span.End()

	user, err := p.Handler.GetUser(ctx, GetUserParams{
		ID: pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true},
	})
	if err != nil {
		telemetree.RecordError(ctx, err, "failed to get user by id")

		return nil, err
	}

	var lastLogin *time.Time
	if user.LastLogin.Valid {
		lastLogin = &user.LastLogin.Time
	}

	return &iam.User{
		ID:           user.ID.Bytes,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Time,
		LastLogin:    lastLogin,
	}, nil
}

func (p PostgresAdapter) CreateRefreshToken(ctx context.Context, token *iam.RefreshToken) error {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.create_refresh_token")
	defer span.End()

	_, err := p.Handler.CreateRefreshToken(ctx, CreateRefreshTokenParams{
		ID:        pgtype.UUID{Bytes: token.ID, Valid: true},
		UserID:    pgtype.UUID{Bytes: token.UserID, Valid: true},
		Token:     token.Token,
		ExpiresAt: pgtype.Timestamp{Time: token.ExpiresAt, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: token.CreatedAt, Valid: true},
		Revoked:   token.Revoked,
	})

	return err
}

func (p PostgresAdapter) GetRefreshToken(ctx context.Context, id string) (*iam.RefreshToken, error) {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.get_refresh_token")
	defer span.End()

	tokenRow, err := p.Handler.GetRefreshToken(ctx, GetRefreshTokenParams{
		ID: pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &iam.RefreshToken{
		ID:        tokenRow.ID.Bytes,
		UserID:    tokenRow.UserID.Bytes,
		Token:     tokenRow.Token,
		ExpiresAt: tokenRow.ExpiresAt.Time,
		CreatedAt: tokenRow.CreatedAt.Time,
		Revoked:   tokenRow.Revoked,
	}, nil
}

func (p PostgresAdapter) UpdateRefreshToken(ctx context.Context, tokenId string, updateFn func(*iam.RefreshToken) (*iam.RefreshToken, error)) error {
	ctx, span := telemetree.AddSpan(ctx, "iam.adapters.outbound.postgres.update_refresh_token")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	token, err := p.GetRefreshToken(ctx, tokenId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedToken, err := updateFn(token)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	_, err = p.Handler.UpdateRefreshToken(ctx, UpdateRefreshTokenParams{
		ID:        pgtype.UUID{Bytes: updatedToken.ID, Valid: true},
		UserID:    pgtype.UUID{Bytes: updatedToken.UserID, Valid: true},
		Token:     updatedToken.Token,
		ExpiresAt: pgtype.Timestamp{Time: updatedToken.ExpiresAt, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: updatedToken.CreatedAt, Valid: true},
		Revoked:   updatedToken.Revoked,
	})
	if err != nil {
		telemetree.RecordError(ctx, err)
	}

	return err
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
