package main

func main() {
	var (
		a  int     = 10
		ab float32 = float32(a) // konversi dari int ke float32
		b  int32   = 200
		bb int8    = int8(b) // konversi dari int32 ke int
	)
	println("Nilai a (int) =", a)
	println("Nilai ab (float32) =", ab)
	println("Nilai b (int32) =", b)
	println("Nilai bb (int8) =", bb)

}
