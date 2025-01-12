package barang

type BarangResponse struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
}
