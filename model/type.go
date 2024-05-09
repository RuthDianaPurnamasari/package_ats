package package_ats

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Peneliti struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	TanggalLahir string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Institusi    string             `bson:"institusi,omitempty" json:"institusi,omitempty"`
	Bidang_Studi string             `bson:"bidang_studi,omitempty" json:"bidang_studi,omitempty"`
	Publikasi    Publikasi          `bson:"publikasi,omitempty" json:"publikasi,omitempty"`
}

// Publikasi adalah struktur data untuk menyimpan informasi tentang publikasi.
type Publikasi struct {
	Judul    string `bson:"judul,omitempty" json:"judul,omitempty"`
	Tanggal  time.Time `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Penulis  string `bson:"penulis,omitempty" json:"penulis,omitempty"`
	Category string `bson:"category,omitempty" json:"category,omitempty"`
}

// penelitian adalah struktur data untuk menyimpan informasi tentang penelitian.
type penelitian struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Judul     string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Institusi string             `bson:"institusi,omitempty" json:"institusi,omitempty"`
	Penulis   string             `bson:"penulis,omitempty" json:"penulis,omitempty"`
	Datetime  primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Ringkasan string             `bson:"ringkasan,omitempty" json:"ringkasan,omitempty"`
	Biodata   Peneliti           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}