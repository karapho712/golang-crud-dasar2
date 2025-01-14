package impl

import (
	"context"
	"crud-dasar-go-2/exception"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/model/web/kamar"
	"crud-dasar-go-2/repository"
	"crud-dasar-go-2/service"
	"database/sql"

	"github.com/go-playground/validator"
)

type KamarServiceImpl struct {
	KamarRepository repository.KamarRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

// Create implements service.KamarService.
func (kamarService *KamarServiceImpl) Create(ctx context.Context, request kamar.KamarCreateRequest) kamar.KamarResponse {
	err := kamarService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := kamarService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kamar := entity.Kamar{
		Nama:          request.Nama,
		Tipe:          request.Tipe,
		HargaPerMalam: request.HargaPerMalam,
		Deskripsi:     request.Deskripsi,
		Barang:        request.BarangItems,
	}

	kamar = kamarService.KamarRepository.Save(ctx, tx, kamar)

	return helper.ToKamarResponse(kamar)

}

// Delete implements service.KamarService.
func (kamarService *KamarServiceImpl) Delete(ctx context.Context, kamarId int) {
	tx, err := kamarService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kamar, err := kamarService.KamarRepository.FindById(ctx, tx, kamarId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	kamarService.KamarRepository.Delete(ctx, tx, kamar)
}

// FindAll implements service.KamarService.
func (kamarService *KamarServiceImpl) FindAll(ctx context.Context) []kamar.KamarResponse {
	tx, err := kamarService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kamars := kamarService.KamarRepository.FindAll(ctx, tx)

	return helper.ToKamarResponses(kamars)
}

// FindById implements service.KamarService.
func (kamarService *KamarServiceImpl) FindById(ctx context.Context, kamarId int) kamar.KamarResponse {
	tx, err := kamarService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kamar, err := kamarService.KamarRepository.FindById(ctx, tx, kamarId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToKamarResponse(kamar)
}

// Update implements service.KamarService.
func (kamarService *KamarServiceImpl) Update(ctx context.Context, request kamar.KamarUpdateRequest) kamar.KamarResponse {
	err := kamarService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := kamarService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	kamar, err := kamarService.KamarRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	kamar.Nama = request.Nama
	kamar.Tipe = request.Tipe
	kamar.HargaPerMalam = request.HargaPerMalam
	kamar.Deskripsi = request.Deskripsi
	kamar.Barang = request.BarangItems

	kamar = kamarService.KamarRepository.Update(ctx, tx, kamar)

	return helper.ToKamarResponse(kamar)
}

func NewKamarService(
	kamarRepository repository.KamarRepository,
	DB *sql.DB,
	Validate *validator.Validate,
) service.KamarService {
	return &KamarServiceImpl{
		KamarRepository: kamarRepository,
		DB:              DB,
		Validate:        Validate,
	}
}
