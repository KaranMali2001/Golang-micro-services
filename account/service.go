package account

import (
	"context"
	"time"
)

type Account struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}
type accountService struct {
	repository Repository
}

func newService(r Repository) Service {
	return &accountService{r}
}
