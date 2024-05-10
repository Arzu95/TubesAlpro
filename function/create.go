package function

import "fmt"

func CreatData(pasien *Patient) {
	var i int
	i = 0
	for &pasien[i].nama != "stop" || &pasien[i].nama != "STOP" {
		fmt.Scan(&pasien[i].nama, &pasien[i].umur, &pasien[i].asal, &pasien[i].tanggalLahir, &pasien[i].gologan_riwayat_penyakit)
		i++
	}

	switch medicalHistory {
	case "ringan":
		patient.Priority = 2
	case "sedang":
		patient.Priority = 1
	case "berat":
		patient.Priority = 0
	}
}
