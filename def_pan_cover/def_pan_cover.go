package main

import "fmt"

// --- Authentication ---
func login() {
	fmt.Println("User berhasil login")
}

func logout() {
	fmt.Println("User berhasil logout")
}

// --- Core Logic ---
func calculate(value int) int {
	if value == 0 { // simulasi error pembagian dengan nol
		panic("Tidak memenuhi syarat") // memicu panic
	}
	return 10 / value // pembagian sederhana
}

// --- Application Runner ---
func runApp(value int) (err error) {

	// Panic handler (boundary)
	defer func() { // menangkap panic yang terjadi di dalam runApp closure
		if r := recover(); r != nil { // recover menangkap panic
			err = fmt.Errorf("aplikasi gagal: %v", r) // mengubah panic menjadi error
		}
	}() // end defer

	login()        // memanggil login
	defer logout() // memastikan logout dipanggil saat runApp selesai

	fmt.Println("Aplikasi berjalan dengan nilai:", value) // menampilkan nilai

	result := calculate(value) // memanggil calculates

	// jika tidak ada panic, lanjutkan eksekusi
	fmt.Println("Hasil perhitungan:", result) // menampilkan hasil
	fmt.Println("Aplikasi selesai dijalankan")

	return nil // tidak ada error
}

func main() {
	if err := runApp(10); err != nil { // menjalankan aplikasi dengan nilai 10
		fmt.Println("ERROR:", err) // menampilkan error jika ada
	}
}
