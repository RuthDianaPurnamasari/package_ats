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
		"Dampak dari Krisis Air Bersih",
		"Universitas Ganjil Genap",
		"Markonah",
		primitive.NewDateTimeFromTime(time.Now()),
		"Dengan menggunakan pendekatan interdisipliner, menyadarkan masyarakat akan isu air bersih dan perlunya tindakan yang cepat dan efektif untuk menjaga sumber daya air yang vital bagi kehidupan dan keberlanjutan lingkungan",
		model.Peneliti{ 
			Nama :        "Bambang Sugiono",
		TanggalLahir:  "20 10 1990",
		Alamat:         "Zimbabwe",
		Institusi:      "Universitas Ganjil Genap",
		Bidang_Studi:  "Kajian Lingkungan atau Kajian Sumber Daya Air",
		Publikasi:      model.Publikasi{
			Judul:         "Dampak dari Krisis Air Bersih",
			Penulis:       "Markonah",
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
	// if err != nil {
	// 	t.Errorf("Error getting all Penelitian: %v", err)
	// 	return
	// }

	// fmt.Println("All Penelitian:")
	// for _, penelitian := range penelitians {
	// 	fmt.Println(penelitian)
	// }
	fmt.Println(penelitians)
}

