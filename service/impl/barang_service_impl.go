package impl

import (
	"context"
	"crud-dasar-go-2/exception"
	"crud-dasar-go-2/helper"
	"crud-dasar-go-2/model/entity"
	"crud-dasar-go-2/model/web/barang"
	"crud-dasar-go-2/repository"
	"crud-dasar-go-2/service"
	"database/sql"

	"github.com/go-playground/validator"
)

type BarangServiceImpl struct {
	BarangRepository repository.BarangRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

// Create implements service.BarangService.
func (barangService *BarangServiceImpl) Create(ctx context.Context, request barang.BarangCreateRequest) barang.BarangResponse {
	err := barangService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := barangService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang := entity.Barang{
		Nama:     request.Nama,
		Kategori: request.Kategori,
	}

	barang = barangService.BarangRepository.Save(ctx, tx, barang)

	return helper.ToBarangResponse(barang)
}

// Delete implements service.BarangService.
func (barangService *BarangServiceImpl) Delete(ctx context.Context, barangId int) {
	tx, err := barangService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang, err := barangService.BarangRepository.FindById(ctx, tx, barangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	barangService.BarangRepository.Delete(ctx, tx, barang)
}

// FindAll implements service.BarangService.
func (barangService *BarangServiceImpl) FindAll(ctx context.Context) []barang.BarangResponse {
	tx, err := barangService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barangs := barangService.BarangRepository.FindAll(ctx, tx)

	return helper.ToBarangResponses(barangs)
}

// FindById implements service.BarangService.
func (barangService *BarangServiceImpl) FindById(ctx context.Context, barangId int) barang.BarangResponse {
	tx, err := barangService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang, err := barangService.BarangRepository.FindById(ctx, tx, barangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBarangResponse(barang)
}

// Update implements service.BarangService.
func (barangService *BarangServiceImpl) Update(ctx context.Context, request barang.BarangUpdateRequest) barang.BarangResponse {
	err := barangService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := barangService.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	barang, err := barangService.BarangRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	barang.Nama = request.Nama
	barang.Kategori = request.Kategori

	barang = barangService.BarangRepository.Update(ctx, tx, barang)

	return helper.ToBarangResponse(barang)
}

func NewBarangService(
	barangRepository repository.BarangRepository,
	DB *sql.DB,
	Validate *validator.Validate,
) service.BarangService {
	return &BarangServiceImpl{
		BarangRepository: barangRepository,
		DB:               DB,
		Validate:         Validate,
	}
}
