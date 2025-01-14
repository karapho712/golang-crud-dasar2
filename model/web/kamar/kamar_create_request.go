package kamar

import "crud-dasar-go-2/model/entity"

type KamarCreateRequest struct {
	Nama          string              `validate:"required,max=150,min=3" json:"nama"`
	Tipe          string              `validate:"required,max=100,min=3" json:"tipe"`
	HargaPerMalam int                 `validate:"required,numeric,max=2000,min=3" json:"harga_per_malam"`
	Deskripsi     string              `validate:"required,max=500,min=3" json:"deskripsi"`
	BarangItems   []entity.BarangItem `validate:"dive,numeric" json:"barang_items"`
}
