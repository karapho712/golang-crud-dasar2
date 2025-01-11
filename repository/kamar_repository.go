package repository

import (
	"context"
	"crud-dasar-go-2/model/entity"
	"database/sql"
)

type KamarRepository interface {
	Save(ctx context.Context, tx *sql.Tx, kamar entity.Kamar) entity.Kamar
	Update(ctx context.Context, tx *sql.Tx, kamar entity.Kamar) entity.Kamar
	Delete(ctx context.Context, tx *sql.Tx, kamar entity.Kamar)
	FindById(ctx context.Context, tx *sql.Tx, kamarId int) (entity.Kamar, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Kamar
}
