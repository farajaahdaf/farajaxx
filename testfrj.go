package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	fmt.Println("Masukkan data (format: ID|Password|KeyA2f|Email|PasswordMail atau ID|Password|KeyA2f|Email|PasswordMail|), ketik 'cukup' untuk selesai:")

	// Membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	var outputBuffer strings.Builder
	var isFirstEntry = true

	for {
		scanner.Scan()
		input := scanner.Text()

		if input == "cukup" {
			break
		}

		// Menghapus tanda "|" di akhir baris jika ada
		input = strings.TrimRight(input, "|")

		data := strings.Split(input, "|")

		if len(data) != 5 {
			fmt.Println("Format data tidak valid. Harap masukkan data dalam format yang benar.")
			continue
		}

		if !isFirstEntry {
			outputBuffer.WriteString("\n") // Tambahkan baris kosong di antara setiap hasil
		} else {
			isFirstEntry = false
		}

		fmt.Printf("ID : %s\n", data[0])
		fmt.Printf("Password : %s\n", data[1])
		fmt.Printf("KeyA2f : %s\n", data[2])
		fmt.Printf("Email : %s\n", data[3])
		fmt.Printf("PasswordMail : %s\n", data[4])

		// Menambahkan hasil output ke buffer
		outputBuffer.WriteString(fmt.Sprintf("ID : %s\n", data[0]))
		outputBuffer.WriteString(fmt.Sprintf("Password : %s\n", data[1]))
		outputBuffer.WriteString(fmt.Sprintf("KeyA2f : %s\n", data[2]))
		outputBuffer.WriteString(fmt.Sprintf("Email : %s\n", data[3]))
		outputBuffer.WriteString(fmt.Sprintf("PasswordMail : %s\n", data[4]))
	}

	// Menyalin hasil output ke clipboard
	outputText := outputBuffer.String()
	clipboard.WriteAll(outputText)

	fmt.Println("Hasil telah disalin ke clipboard.")
}
