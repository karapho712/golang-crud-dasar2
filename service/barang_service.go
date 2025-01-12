package service

import (
	"context"
	"crud-dasar-go-2/model/web/barang"
)

type BarangService interface {
	Create(ctx context.Context, request barang.BarangCreateRequest) barang.BarangResponse
	Update(ctx context.Context, request barang.BarangUpdateRequest) barang.BarangResponse
	Delete(ctx context.Context, barangId int)
	FindById(ctx context.Context, barangId int) barang.BarangResponse
	FindAll(ctx context.Context) []barang.BarangResponse
}
