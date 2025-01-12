package barang

type BarangUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Nama     string `validate:"max=150,min=3" json:"nama"`
	Kategori string `validate:"max=100,min=3" json:"kategori"`
}
