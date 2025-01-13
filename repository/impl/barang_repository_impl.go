package impl

import (
	"context"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/repository"
	"database/sql"
	"errors"
)

type BarangRepositoryImpl struct {
}

// Delete implements repository.BarangRepository.
func (b *BarangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, barang entity.Barang) {
	SQL := "DELETE FROM barang where id = ?"
	_, err := tx.ExecContext(ctx, SQL, barang.Id)

	helper.PanicIfError(err)
}

// FindAll implements repository.BarangRepository.
func (b *BarangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Barang {
	SQL := "SELECT id, nama, kategori FROM barang"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var barangs []entity.Barang
	for rows.Next() {
		barang := entity.Barang{}
		err := rows.Scan(&barang.Id, &barang.Nama, &barang.Kategori)
		helper.PanicIfError(err)

		barangs = append(barangs, barang)
	}
	return barangs
}

// FindById implements repository.BarangRepository.
func (b *BarangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, barangId int) (entity.Barang, error) {
	SQL := "SELECT id, nama, kategori FROM barang WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, barangId)
	helper.PanicIfError(err)

	defer rows.Close()

	barang := entity.Barang{}
	if rows.Next() {
		err := rows.Scan(&barang.Id, &barang.Nama, &barang.Kategori)
		helper.PanicIfError(err)
		return barang, nil
	} else {
		return barang, errors.New("barang not found")
	}
}

// Save implements repository.BarangRepository.
func (b *BarangRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, barang entity.Barang) entity.Barang {
	SQL := "INSERT INTO barang(nama, kategori) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, barang.Nama, barang.Kategori)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	barang.Id = int(id)

	return barang
}

// Update implements repository.BarangRepository.
func (b *BarangRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, barang entity.Barang) entity.Barang {
	SQL := "UPDATE barang SET nama = ?, kategori = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, barang.Nama, barang.Kategori, barang.Id)
	helper.PanicIfError(err)

	return barang
}

func NewBarangRepository() repository.BarangRepository {
	return &BarangRepositoryImpl{}
}
