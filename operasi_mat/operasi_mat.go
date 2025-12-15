package main

func main() { // fungsi math biasa program
	//operasimatematika() // memanggil fungsi operasimatematika
	//operasibandingan() // memanggil fungsi operasibandingan
	operasiboolean() // memanggil fungsi operasiboolean
}
func operasimatematika() { // fungsi math operasi matematika
	var (
		a = 10    // deklarasi variabel a dengan nilai 10
		b = 20    // deklarasi variabel b dengan nilai 20
		c = a + b // deklarasi variabel c dengan hasil penjumlahan a dan b
		d = c - a // deklarasi variabel d dengan hasil pengurangan c dan a
		e = d * b // deklarasi variabel e dengan hasil perkalian d dan b
		f = e / a // deklarasi variabel f dengan hasil pembagian e dan a
		g = f % 3 // deklarasi variabel g dengan hasil sisa bagi f dan 3
		h = 15    // deklarasi variabel h dengan nilai 15
	)
	println("Hasil penjumlahan a + b =", c)    // mencetak nilai variabel c
	println("Hasil pengurangan c - a =", d)    // mencetak nilai variabel d
	println("Hasil perkalian d * b =", e)      // mencetak nilai variabel e
	println("Hasil pembagian e / a =", f)      // mencetak nilai variabel f
	println("Hasil sisa bagi f % 3 =", g)      // mencetak nilai variabel g
	h += 5                                     // operasi penambahan nilai h dengan 5 dengan operator penugasan
	println("Nilai h setelah ditambah 5 =", h) // mencetak nilai variabel h setelah ditambah 5
	a--
	println("Nilai a setelah dikurangi 1 =", a) // mencetak nilai variabel a setelah dengan operator unary dikurangi 1
}

func operasibandingan() { // fungsi math operasi matematika perbandingan
	var (
		a = 10 // deklarasi variabel a dengan nilai 10
		b = 20 // deklarasi variabel b dengan nilai 20
	)
	println("Apakah a lebih besar dari b?", a > b)              // mencetak hasil operasi perbandingan lebih besar
	println("Apakah a lebih kecil dari b?", a < b)              // mencetak hasil operasi perbandingan lebih kecil
	println("Apakah a sama dengan b?", a == b)                  // mencetak hasil operasi perbandingan sama dengan
	println("Apakah a tidak sama dengan b?", a != b)            // mencetak hasil operasi perbandingan tidak sama dengan
	println("Apakah a lebih besar atau sama dengan b?", a >= b) // mencetak hasil operasi perbandingan lebih besar atau sama dengan
	println("Apakah a lebih kecil atau sama dengan b?", a <= b) // mencetak hasil operasi perbandingan lebih kecil atau sama dengan
}

func operasiboolean() { // fungsi math operasi matematika boolean
	var (
		nilaiA = 90  // deklarasi variabel nilaiA dengan nilai true
		nilaiB = 100 // deklarasi variabel nilaiB dengan nilai false
	)
	bolean1 := nilaiA > 80                                 // deklarasi variabel boolean1 dengan hasil operasi perbandingan nilaiA lebih besar dari 80
	bolean2 := nilaiB < 90                                 // deklarasi variabel boolean2 dengan hasil operasi perbandingan nilaiB lebih kecil dari 90
	bolean3 := nilaiA == nilaiB                            // deklarasi variabel boolean3 dengan hasil operasi perbandingan nilaiA sama dengan nilaiB
	println("Apakah nilaiA lebih besar dari 80?", bolean1) // mencetak hasil operasi perbandingan nilaiA lebih besar dari 80
	println("Apakah nilaiB lebih kecil dari 90?", bolean2)
	println("Apakah nilaiA sama dengan nilaiB?", bolean3)                               // mencetak hasil operasi perbandingan boolean nilaiA sama dengan nilaiB
	println("apakah nilaiA >= 90 || nilaiB <= 90 ", nilaiA >= 90 || nilaiB <= 90)       // mencetak hasil operasi logika OR
	println("apakah nilaiA >= 90 && nilaiB <= 90 ", nilaiA >= 90 && nilaiB <= 90)       // mencetak hasil operasi logika AND
	println("apakah !(nilaiA >= 90 && nilaiB <= 90) ", !(nilaiA >= 90 && nilaiB <= 90)) // mencetak hasil operasi logika NOT

}
