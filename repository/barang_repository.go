package repository

import (
	"context"
	"crud-dasar-go-2/model/entity"
	"database/sql"
)

type BarangRepository interface {
	Save(ctx context.Context, tx *sql.Tx, barang entity.Barang) entity.Barang
	Update(ctx context.Context, tx *sql.Tx, barang entity.Barang) entity.Barang
	Delete(ctx context.Context, tx *sql.Tx, barang entity.Barang)
	FindById(ctx context.Context, tx *sql.Tx, barangId int) (entity.Barang, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Barang
}
