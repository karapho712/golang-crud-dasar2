package kamar

type KamarCreateRequest struct {
	Nama          string `validate:"required,max=150,min=3" json:"nama"`
	Tipe          string `validate:"required,max=100,min=3" json:"tipe"`
	HargaPerMalam string `validate:"required,numeric, max=100,min=3" json:"harga_per_malam"`
	Deskripsi     string `validate:"required,max=500,min=3" json:"deskripsi"`
}
