package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		variabel untuk menyimpan waktu yang ingin dikonversi
	*/
	var waktu, format_waktu string

	// contoh format waktu yang dimasukkan 12:00 AM
	fmt.Print("Masukkan waktu : ")
	fmt.Scanln(&waktu, &format_waktu)

	// ini proses pengecekan format waktunya
	if format_waktu == "AM" {
		// ini proses pengecekan waktu
		jam := strings.Split(waktu, ":")[0]
		menit := strings.Split(waktu, ":")[1]

		var num, err = strconv.Atoi(jam)
		var mnt, erro = strconv.Atoi(menit)

		if err == nil || erro == nil {
			for i := 0; i <= 12; i++ {
				if num > 12 || mnt > 59 {
					fmt.Println("🔘 jam atau menit yang anda masukkan melebihi batas, silahkan cek kembali")
					return
				} else if num == i {
					if num < 10 {
						fmt.Println("\nResult Convertion :")
						fmt.Println("🆗", "0"+strconv.Itoa(num), ":", menit)
					} else if num == 12 {
						fmt.Println("\nResult Convertion :")
						fmt.Println("🆗", "00", ":", menit)
					} else if num >= 10 && num <= 11 {
						fmt.Println("\nResult Convertion :")
						fmt.Println("🆗", num, ":", menit)
					}
				}
			}
		}
	} else if format_waktu == "PM" {
		// ini proses pengecekan waktu
		jam := strings.Split(waktu, ":")[0]
		menit := strings.Split(waktu, ":")[1]

		var num, err = strconv.Atoi(jam)
		var mnt, erro = strconv.Atoi(menit)

		if err == nil || erro == nil {
			for i := 0; i <= 12; i++ {
				if num > 12 || mnt > 59 {
					fmt.Println("🔘 jam atau menit yang anda masukkan melebihi batas, silahkan cek kembali")
					return
				} else if num == i {
					if num <= 11 {
						fmt.Println("\nResult Convertion :")
						fmt.Println("🆗", 12+num, ":", menit)
					} else if num == 12 {
						fmt.Println("\nResult Convertion :")
						fmt.Println("🆗", num, ":", menit)
					}
				}
			}
		}
	} else {
		fmt.Println("🔘 Waktu yang anda masukkan tidak sesuai format yang sudah ditentukan")
	}
}
