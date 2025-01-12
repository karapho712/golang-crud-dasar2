package barang

type BarangCreateRequest struct {
	Nama     string `validate:"required,max=150,min=3" json:"nama"`
	Kategori string `validate:"required,max=100,min=3" json:"kategori"`
}
