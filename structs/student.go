package structs

type Student struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
	ClassID      int64  `json:"class_id"`
}
