package repository

import (
	"context"
	"github.com/amzdll/backend_2_go/src/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (repository ClientRepository) GetAll(ctx context.Context, pagination model.Pagination) ([]model.Client, error) {
	const query = `select * from client limit $1 offset $2`
	var clients []model.Client
	if err := pgxscan.Select(ctx, repository.db, &clients, query, pagination.Limit, pagination.Offset); err != nil {
		return nil, err
	}
	return clients, nil
}

func (repository ClientRepository) GetByNameSurname(
	ctx context.Context, clientInfo model.ClientInfo,
) ([]model.Client, error) {
	var clients []model.Client
	query := `select * from client where client_name = $1 and client_surname = $2`
	err := pgxscan.Select(
		ctx, repository.db, &clients, query, clientInfo.ClientName, clientInfo.ClientSurname)
	if err != nil {
		return nil, err
	}
	return clients, nil
}