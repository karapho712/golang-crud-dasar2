package kamar

type KamarUpdateRequest struct {
	Id            int    `validate:"required" json:"id"`
	Nama          string `validate:"max=150,min=3" json:"nama"`
	Tipe          string `validate:"max=100,min=3" json:"tipe"`
	HargaPerMalam string `validate:"numeric, max=100,min=3" json:"harga_per_malam"`
	Deskripsi     string `validate:"max=500,min=3" json:"deskripsi"`
}
