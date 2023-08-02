package models

type Inventory struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Nama     string `json:"nama" bson:"nama" validate:"required"`
	Kode     string `json:"kode" bson:"kode" validate:"required"`
	Harga    int    `json:"harga" bson:"harga" validate:"required"`
	Stock    int    `json:"stock" bson:"stock" validate:"required"`
	Kategori string `json:"kategori" bson:"kategori" validate:"required"`
}

type UpsertInventory struct {
	Nama     string `json:"nama,omitempty" bson:"nama,omitempty"`
	Kode     string `json:"kode,omitempty" bson:"kode,omitempty"`
	Harga    int    `json:"harga,omitempty" bson:"harga,omitempty"`
	Stock    int    `json:"stock,omitempty" bson:"stock,omitempty"`
	Kategori string `json:"kategori,omitempty" bson:"kategori,omitempty"`
}

func (u Inventory) UpsertInventory() UpsertInventory {
	return UpsertInventory{
		Nama:     u.Nama,
		Kode:     u.Kode,
		Harga:    u.Harga,
		Stock:    u.Stock,
		Kategori: u.Kategori,
	}
}

type CreateResponse struct {
	Nama string `json:"nama"`
	Kode string `json:"kode"`
}

type GetInventoryResponse struct {
	Id       string `json:"id,omitempty"`
	Nama     string `json:"nama"`
	Kode     string `json:"kode"`
	Harga    int    `json:"harga"`
	Stock    int    `json:"sStock"`
	Kategori string `json:"kategori,omitempty"`
}
