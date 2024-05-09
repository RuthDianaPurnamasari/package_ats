package package_ats_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/RuthDianaPurnamasari/package_ats/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPenelitian(t *testing.T) {
	biodata := map[string]interface{}{
		"nama":          "Bambang Sugiono",
		"tanggal_lahir": "20 10 1990",
		"alamat":        "Zimbabwe",
		"institusi": 	 "Universitas Ganjil Genap",
		"bidang_studi":	 "Kajian Lingkungan atau Kajian Sumber Daya Air",
		"publikasi":	 "Publikasi Peneliti",
		"judul":	     "Dampak dari Krisis Air Bersih",
		"tanggal":		 "15 10 2002",
		"penulis":	     "Markonah",
		"category":	     "Jurnal",
	}

	insertedID := package_ats.InsertPenelitian(
		"Dampak dari Krisis Air Bersih",
		"Universitas Ganjil Genap",
		"Markonah",
		primitive.NewDateTimeFromTime(time.Now()),
		"Dengan menggunakan pendekatan interdisipliner,menyadarkan masyarakat akan isu air bersih dan perlunya tindakan yang cepat dan efektif untuk menjaga sumber daya air yang vital bagi kehidupan dan keberlanjutan lingkungan",
		biodata,
	)
	fmt.Println("Inserted Penelitian ID:", insertedID)
}

func TestGetAllPenelitian(t *testing.T) {
	penelitians := package_ats.GetAllPenelitian()
	fmt.Println("All Penelitian:", penelitians)
}
