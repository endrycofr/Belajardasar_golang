package main

import (
	"container/ring"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(6) // membuat ring dengan 6 elemen kosong
	for i := 0; i < data.Len(); i++ { // menambahkan iterasi  elemen ke dalam ring
		data.Value = "Data " + strconv.Itoa(i+1) // mengisi elemen dengan string "Data " diikuti nomor urut
		data = data.Next()                       // pindah ke elemen berikutnya
	}
	data.Do(func(value any) { // mencetak setiap elemen dalam ring dengan
		// menggunakan method Do yang memanggil fungsi untuk setiap elemen dalam ring
		println(value.(string))
	})
}
