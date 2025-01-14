package entity

import "time"

type Kamar struct {
	Id            int
	Nama          string
	Tipe          string
	HargaPerMalam int
	Deskripsi     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Barang        []BarangItem
}

type BarangItem struct {
	Id int
}
