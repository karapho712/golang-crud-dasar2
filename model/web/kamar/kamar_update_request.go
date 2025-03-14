package kamar

import "crud-dasar-go-2/model/entity"

type KamarUpdateRequest struct {
	Id            int                 `validate:"required" json:"id"`
	Nama          string              `validate:"max=150,min=3" json:"nama"`
	Tipe          string              `validate:"max=100,min=3" json:"tipe"`
	HargaPerMalam int                 `validate:"numeric,max=2000,min=3" json:"harga_per_malam"`
	Deskripsi     string              `validate:"max=500,min=3" json:"deskripsi"`
	BarangItems   []entity.BarangItem `validate:"dive,numeric" json:"barang_items"`
}
