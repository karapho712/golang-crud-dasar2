package service

import (
	"context"
	"crud-dasar-go-2/model/web/kamar"
)

type KamarService interface {
	Create(ctx context.Context, request kamar.KamarCreateRequest) kamar.KamarResponse
	Update(ctx context.Context, request kamar.KamarUpdateRequest) kamar.KamarResponse
	Delete(ctx context.Context, kamarId int)
	FindById(ctx context.Context, kamarId int) kamar.KamarResponse
	FindAll(ctx context.Context) []kamar.KamarResponse
}
