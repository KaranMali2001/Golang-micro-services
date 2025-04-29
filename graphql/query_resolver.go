package main

import "context"

type queryResolver struct {
	server *Server
}

func (q *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {

}
func (q *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {

}
func (q *queryResolver) Orders(ctx context.Context, pagination *PaginationInput, accountId string) ([]*Order, error) {

}
func (q *queryResolver) Account(ctx context.Context, id string) (*Account, error) {

}
func (q *queryResolver) Product(ctx context.Context, id string) (*Product, error) {

}
func (q *queryResolver) Order(ctx context.Context, id string) (*Order, error) {

}
