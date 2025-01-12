package impl

import (
	"context"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/repository"
	"database/sql"
	"errors"
)

type KamarRepositoryImpl struct {
}

// Delete implements repository.KamarRepository.
func (k *KamarRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, kamar entity.Kamar) {
	SQL := "DELETE FROM kamar where id = ?"
	_, err := tx.ExecContext(ctx, SQL, kamar.Id)

	helper.PanicIfError(err)
}

// FindAll implements repository.KamarRepository.
func (k *KamarRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Kamar {
	SQL := "SELECT id, nama, tipe, harga_per_malam, deskripsi FROM kamar"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var kamars []entity.Kamar
	for rows.Next() {
		kamar := entity.Kamar{}
		err := rows.Scan(&kamar.Id, &kamar.Tipe, &kamar.HargaPerMalam, &kamar.Deskripsi)
		helper.PanicIfError(err)

		kamars = append(kamars, kamar)
	}
	return kamars
}

// FindById implements repository.KamarRepository.
func (k *KamarRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kamarId int) (entity.Kamar, error) {
	SQL := "SELECT nama, tipe, harga_per_malam, deskripsi FROM kamar WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, kamarId)
	helper.PanicIfError(err)

	defer rows.Close()

	kamar := entity.Kamar{}
	if rows.Next() {
		err := rows.Scan(&kamar.Id, &kamar.Tipe, &kamar.HargaPerMalam, &kamar.Deskripsi)
		helper.PanicIfError(err)
		return kamar, nil
	} else {
		return kamar, errors.New("kamar not found")
	}
}

// Save implements repository.KamarRepository.
func (k *KamarRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, kamar entity.Kamar) entity.Kamar {
	SQL := "INSERT INTO kamar(nama, tipe, harga_per_malam, deskripsi) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, kamar.Nama, kamar.Tipe, kamar.HargaPerMalam, kamar.Deskripsi)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	kamar.Id = int(id)

	return kamar
}

// Update implements repository.KamarRepository.
func (k *KamarRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, kamar entity.Kamar) entity.Kamar {
	SQL := "UPDATE kamar SET nama = ? , tipe = ?, harga_per_malam = ?, deskripsi = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, kamar.Nama, kamar.Tipe, kamar.HargaPerMalam, kamar.Deskripsi, kamar.Id)
	helper.PanicIfError(err)

	return kamar
}

func NewKamarRepository() repository.KamarRepository {
	return &KamarRepositoryImpl{}
}
