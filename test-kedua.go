package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string

	fmt.Print("Masukkan email : ")
	fmt.Scanf("%s", &email)

	fmt.Println("\nResults Cheking Email :")

	// check @ pada email
	if strings.Contains(email, "@") {
		fmt.Println("✅ email sudah menggunakan @")
		// Check titik setelah @
		var checkDot = strings.Split(email, "@")
		if strings.Contains(checkDot[1], ".") {
			fmt.Println("✅ email sudah menggunakan titik setelah @")
			// Check panjang email sebelum @
			var checkLength = len(strings.Split(email, "@")[0])
			if checkLength <= 20 {
				fmt.Println("✅ panjang email tidak lebih dari 20 karakter")
				// Check domain yang diperbolehkan
				var checkDomain = strings.Split(email, "@")[1]
				if strings.Contains(checkDomain, ".id") || strings.Contains(checkDomain, ".co.id") {
					fmt.Println("✅ email sudah menggunakan domain .id atau .co.id")
					fmt.Println("\n 🆗 Email sudah memenuhi persyaratan")
				} else {
					fmt.Println("❌ email tidak menggunakan domain .id atau .co.id")
					fmt.Println("\n ❌ Email belum memenuhi persyaratan")
				}
			} else {
				fmt.Println("❌ panjang email tidak lebih dari 20 karakter")
				fmt.Println("\n ❌ Email belum memenuhi persyaratan")
			}
		} else {
			fmt.Println("❌ email sudah menggunakan titik setelah @")
			fmt.Println("\n ❌ Email belum memenuhi persyaratan")
		}
	} else {
		fmt.Println("❌ email tidak menggunakan @")
		fmt.Println("\n ❌ Email belum memenuhi persyaratan")
	}
}
