package account

import (
	"context"
	"database/sql"

	"github.com/micro-services-golang/account/db/sqlc-generated"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error)
}
type postgresRepository struct {
	db *sql.DB
	q  *sqlc.Queries
}

func (p *postgresRepository) Close() {
	p.db.Close()
}
func (p *postgresRepository) PutAccount(ctx context.Context, a Account) error {
	return p.q.CreateAccount(ctx, sqlc.CreateAccountParams{
		ID:   a.ID,
		Name: a.Name,
	})

}
func (p *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	account, err := p.q.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
func (p *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error) {
	accounts, err := p.q.ListAccounts(ctx, sqlc.ListAccountsParams{
		Limit:  int32(skip),
		Offset: int32(take),
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
func (p *postgresRepository) NewPostgresRepo(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	q := sqlc.New(db)
	return &postgresRepository{
		db: db,
		q:  q,
	}, nil
}
