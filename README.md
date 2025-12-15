
# Belajar Dasar Golang 🚀

Repository ini berisi kumpulan contoh dan latihan **dasar bahasa pemrograman Go (Golang)** yang saya pelajari secara bertahap, mulai dari konsep fundamental hingga penerapan sederhana.

Repo ini cocok untuk:
- Pemula yang baru belajar Go
- Referensi ulang konsep dasar Golang
- Latihan memahami pointer, struct, function, dan method

---

## 📌 Materi yang Dipelajari

### ✅ Dasar Golang
- Struktur project Go
- Package dan import
- Fungsi `main`
- Variabel dan tipe data
- Control flow (`if`, `for`, `switch`)

### ✅ Struct
- Deklarasi struct
- Inisialisasi struct
- Akses field struct

```go
type Country struct {
	city     string
	region   string
	province string
	zip      int
}
````

---

### ✅ Function

* Function biasa
* Parameter & return value
* Function dengan pointer

```go
func changeCity(country *Country) {
	country.city = "Bandung"
}
```

---

### ✅ Method

* Method receiver
* Pointer receiver vs value receiver
* Perbedaan function dan method

```go
func (c *Country) ChangeProvince() {
	c.province = "Jawa Barat"
}
```

---

### ✅ Pointer

* Pointer dasar
* Operator `&` dan `*`
* Pass by value vs pass by reference
* Pointer ke struct

---

### ✅ Interface (Dasar)

* Konsep interface
* Method sebagai kontrak
* Implementasi interface sederhana

---

## 🧠 Tujuan Repository

* Memahami **konsep fundamental Golang**
* Membiasakan diri dengan **idiomatic Go**
* Menjadi dasar sebelum masuk ke:

  * Backend development
  * Microservices
  * DevOps tooling (CLI, automation)

---

## ▶️ Cara Menjalankan Code

Pastikan Go sudah terinstall:

```bash
go version
```

Clone repository:

```bash
git clone https://github.com/endrycofr/Belajardasar_golang.git
cd Belajardasar_golang
```

Jalankan file Go:

```bash
go run main.go
```

---

## 📂 Struktur Folder (Contoh)

```
Belajardasar_golang/
├── main.go
├── README.md
└── ...
```

---

## 📝 Catatan

Repository ini bersifat **learning repository**, sehingga:

* Kode dibuat sederhana
* Fokus pada pemahaman konsep
* Disertai komentar penjelasan

---

## 👤 Author

**Endryco Rahmat**
📌 GitHub: [https://github.com/endrycofr](https://github.com/endrycofr)

---

## ⭐ Penutup

Jika repository ini bermanfaat:

* Jangan lupa ⭐ star
* Silakan fork untuk latihan sendiri

Happy Coding with Go! 🐹🔥

```



Tinggal bilang aja 👍
```
