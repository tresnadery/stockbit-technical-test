package repository

import (
	"context"
	"database/sql"
	"stockbit-technical-test/test2/internal/domain"
)

type repository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) domain.MovieRepository {
	return &repository{
		Conn: conn,
	}
}

func (r *repository) StoreLog(ctx context.Context, payload *domain.Log) error {
	stmt, err := r.Conn.PrepareContext(ctx, "INSERT logs SET request=? , response=? , created_at=?")
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, payload.Request, payload.Response, payload.CreatedAt)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	payload.ID = lastID
	return nil
}
