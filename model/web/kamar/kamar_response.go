package kamar

type KamarResponse struct {
	Id            int    `json:"id"`
	Nama          string `json:"nama"`
	Tipe          string `json:"tipe"`
	HargaPerMalam string `json:"harga_per_malam"`
	Deskripsi     string `json:"deskripsi"`
}
