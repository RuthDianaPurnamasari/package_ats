package package_ats_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RuthDianaPurnamasari/package_ats/model"
	"github.com/RuthDianaPurnamasari/package_ats/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPenelitian(t *testing.T) {

	insertedID := module.InsertPenelitian(
		"Sistem Irigasi Pintar Berbasis Internet of Things (IoT) untuk Pertanian Modern",
		"Universitas Gadjah Mada",
		"Changsub",
		primitive.NewDateTimeFromTime(time.Now()),
		"Pengembangan sistem irigasi otomatis menggunakan sensor kelembapan tanah dan kontroler IoT untuk meningkatkan efisiensi penggunaan air di lahan pertanian.",
		model.Peneliti{ 
			Nama :        "Changsub",
		TanggalLahir:  "21 11 2001",
		Alamat:         "Jalan Bulaksumur, Yogyakarta",
		Institusi:      "Universitas Gadjah Mada",
		Bidang_Studi:  "Teknik Pertanian",
		Publikasi:      model.Publikasi{
			Judul:         "Sistem Irigasi Pintar Berbasis Internet of Things (IoT) untuk Pertanian Modern",
			Penulis:       "Changsub",
			Category:      "Jurnal",	
		},
		},
	)
	// if err != nil {
	// 	t.Errorf("Error inserting Penelitian: %v", err)
	// 	return
	// }
	fmt.Println(insertedID)
}

func TestGetAllPenelitian(t *testing.T) {
	penelitians := module.GetAllPenelitian()
	fmt.Println(penelitians)
}

