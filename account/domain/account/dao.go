package account

import (
	"context"
	"database/sql"
)

const (
	selectQuery    = `SELECT ID, NAME FROM accounts WHERE ID=($1);`
	insertQuery    = `INSERT INTO accounts (ID, NAME) VALUES ($1, $2);`
	selectAllQuery = `SELECT ID, NAME FROM accounts ORDER BY ID OFFSET ($1) LIMIT ($2);`
)

// Repository interface
type Repository interface {
	Close()
	Select(context.Context, string) (*Account, error)
	Insert(context.Context, Account) error
	SelectAll(context.Context, uint64, uint64) ([]Account, error)
}

type accountRepo struct {
	db *sql.DB
}

// NewAccountRepo returns a new account repository
func NewAccountRepo(db *sql.DB) Repository {
	return &accountRepo{
		db: db,
	}
}

func (ar *accountRepo) getConn(ctx context.Context) *sql.Conn {
	conn, err := ar.db.Conn(ctx)
	if err != nil {
		return nil
	}
	return conn
}

func (ar *accountRepo) Close() {
	ar.db.Close()
}

func (ar *accountRepo) Insert(ctx context.Context, a Account) error {
	conn := ar.getConn(ctx)
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, insertQuery)
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(ctx, a.ID, a.Name); err != nil {
		return err
	}

	return nil
}

func (ar *accountRepo) Select(ctx context.Context, id string) (*Account, error) {
	conn := ar.getConn(ctx)
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, selectQuery)
	if err != nil {
		return nil, err
	}
	a := &Account{}
	row := stmt.QueryRowContext(ctx, id)
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return a, nil
}

func (ar *accountRepo) SelectAll(ctx context.Context, skip, take uint64) ([]Account, error) {
	conn := ar.getConn(ctx)
	defer conn.Close()

	stmt, err := conn.PrepareContext(ctx, selectAllQuery)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	accs := []Account{}

	for rows.Next() {
		a := &Account{}
		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		accs = append(accs, *a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return accs, nil
}
