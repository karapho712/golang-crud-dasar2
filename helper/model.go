package helper

import (
	"crud-dasar-go-2/model/entity"
	barangResponse "crud-dasar-go-2/model/web/barang"
	kamarResponse "crud-dasar-go-2/model/web/kamar"
)

func ToKamarResponse(kamar entity.Kamar) kamarResponse.KamarResponse {
	return kamarResponse.KamarResponse{
		Id:            kamar.Id,
		Nama:          kamar.Nama,
		Tipe:          kamar.Tipe,
		HargaPerMalam: kamar.HargaPerMalam,
		Deskripsi:     kamar.Deskripsi,
	}
}

func ToKamarResponses(kamars []entity.Kamar) []kamarResponse.KamarResponse {
	var kamarResponses []kamarResponse.KamarResponse
	for _, kamar := range kamars {
		kamarResponses = append(kamarResponses, ToKamarResponse(kamar))
	}

	return kamarResponses
}

func ToBarangResponse(barang entity.Barang) barangResponse.BarangResponse {
	return barangResponse.BarangResponse{
		Id:       barang.Id,
		Nama:     barang.Nama,
		Kategori: barang.Kategori,
	}
}

func ToBarangResponses(barangs []entity.Barang) []barangResponse.BarangResponse {
	var barangResponses []barangResponse.BarangResponse
	for _, barang := range barangs {
		barangResponses = append(barangResponses, ToBarangResponse(barang))
	}

	return barangResponses
}
