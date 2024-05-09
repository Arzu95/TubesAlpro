package main

import (
	"fmt"
	"sort"
)

// Struct untuk merepresentasikan data pasien
type Patient struct {
	Name            string
	Age             int
	Origin          string
	DateOfBirth     string
	MedicalHistory  string
	Priority        int
}

// Slice untuk menyimpan daftar pasien
var patients []Patient

// Fungsi untuk menambahkan pasien ke dalam antrian
func AddPatient(name string, age int, origin string, dob string, medicalHistory string) {
	patient := Patient{
		Name:           name,
		Age:            age,
		Origin:         origin,
		DateOfBirth:    dob,
		MedicalHistory: medicalHistory,
		Priority:       0,
	}

	// Menetapkan prioritas berdasarkan riwayat penyakit
	switch medicalHistory {
	case "ringan":
		patient.Priority = 2
	case "sedang":
		patient.Priority = 1
	case "berat":
		patient.Priority = 0
	}

	// Menambahkan pasien ke dalam slice
	patients = append(patients, patient)

	// Mengurutkan antrian berdasarkan prioritas
	sort.Slice(patients, func(i, j int) bool {
		return patients[i].Priority < patients[j].Priority
	})
}

// Fungsi untuk mencari data pasien berdasarkan nama
func FindPatientByName(name string) *Patient {
	for _, patient := range patients {
		if patient.Name == name {
			return &patient
		}
	}
	return nil
}

// Fungsi untuk mencari data pasien berdasarkan riwayat penyakit
func FindPatientByMedicalHistory(history string) []Patient {
	var result []Patient
	for _, patient := range patients {
		if patient.MedicalHistory == history {
			result = append(result, patient)
		}
	}
	return result
}

// Fungsi untuk mengubah data pasien
func EditPatient(name string, newAge int, newOrigin string, newDob string, newMedicalHistory string) {
	for i, patient := range patients {
		if patient.Name == name {
			patients[i].Age = newAge
			patients[i].Origin = newOrigin
			patients[i].DateOfBirth = newDob
			patients[i].MedicalHistory = newMedicalHistory
		}
	}
}

// Fungsi untuk menghapus data pasien
func DeletePatient(name string) {
	for i, patient := range patients {
		if patient.Name == name {
			patients = append(patients[:i], patients[i+1:]...)
		}
	}
}

// Fungsi untuk menampilkan data antrian yang terurut
func DisplayQueue() {
	fmt.Println("Antrian Pasien:")
	for _, patient := range patients {
		fmt.Printf("- %s (Prioritas: %d)\n", patient.Name, patient.Priority)
	}
}

func main() {
	// Menambahkan beberapa contoh pasien ke dalam antrian
	AddPatient("John Doe", 35, "Jakarta", "1989-05-20", "sedang")
	AddPatient("Jane Smith", 40, "Surabaya", "1984-10-12", "ringan")
	AddPatient("Michael Johnson", 50, "Bandung", "1974-08-30", "berat")

	// Menampilkan antrian pasien
	DisplayQueue()

	// Mencari pasien berdasarkan nama
	fmt.Println("\nCari Pasien:")
	if patient := FindPatientByName("John Doe"); patient != nil {
		fmt.Printf("Nama: %s, Umur: %d, Riwayat Penyakit: %s\n", patient.Name, patient.Age, patient.MedicalHistory)
	} else {
		fmt.Println("Pasien tidak ditemukan.")
	}

	// Mencari pasien berdasarkan riwayat penyakit
	fmt.Println("\nCari Pasien berdasarkan Riwayat Penyakit:")
	sedangPatients := FindPatientByMedicalHistory("sedang")
	for _, patient := range sedangPatients {
		fmt.Printf("Nama: %s, Umur: %d\n", patient.Name, patient.Age)
	}

	// Mengubah data pasien
	fmt.Println("\nEdit Data Pasien:")
	EditPatient("Jane Smith", 41, "Surabaya", "1984-10-12", "sedang")
	fmt.Println("Data pasien berhasil diubah.")

	// Menghapus data pasien
	fmt.Println("\nHapus Data Pasien:")
	DeletePatient("Michael Johnson")
	fmt.Println("Data pasien berhasil dihapus.")

	// Menampilkan antrian pasien setelah perubahan
	fmt.Println("\nAntrian Pasien setelah Perubahan:")
	DisplayQueue()
}






// package main

// import "fmt"

// type pasien struct {
// 	nama, umur, asal, penyakit ,gologan_riwayat_penyakit string
// 	tanggalLahir int
// }

// func cetakDataTerurut() {

// }

// func pilihRole() {

// }

// func daftarPasien(data pasien, t string) {

// }

// func cariData() {

// }

// func main(){
// 	// masukan data pasien
// 	var data pasien
// 	var 
// 	daftarPasien(&data, )
// }