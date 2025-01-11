package helper

import (
	"crud-dasar-go-2/model/entity"
	response "crud-dasar-go-2/model/web/kamar"
)

func ToKamarResponse(kamar entity.Kamar) response.KamarResponse {
	return response.KamarResponse{
		Id:            kamar.Id,
		Nama:          kamar.Nama,
		Tipe:          kamar.Tipe,
		HargaPerMalam: kamar.HargaPerMalam,
		Deskripsi:     kamar.Deskripsi,
	}
}

func ToKamarResponses(kamars []entity.Kamar) []response.KamarResponse {
	var kamarResponses []response.KamarResponse
	for _, kamar := range kamars {
		kamarResponses = append(kamarResponses, ToKamarResponse(kamar))
	}

	return kamarResponses
}
