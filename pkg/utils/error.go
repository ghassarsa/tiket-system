package utils

import "strings"

func FormatStackTrace(stack []byte) []string {
	// Konversi byte slice ke string, lalu pisahkan per baris
	lines := strings.Split(string(stack), "\n")
	var cleanedLines []string

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue // Lewati baris kosong
		}

		// Hapus bagian "+0x..." dari baris jika ada.
		// Contoh: "main.myFunctionC(): /path/to/file.go:39 +0x29"
		// Akan menjadi: "main.myFunctionC(): /path/to/file.go:39"
		parts := strings.Split(trimmedLine, " +0x")
		cleanedLines = append(cleanedLines, parts[0])
	}
	return cleanedLines
}
