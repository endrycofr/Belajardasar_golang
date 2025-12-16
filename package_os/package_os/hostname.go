package packageos

import "os"

func GetHostname() string {
	//mengambil hostname dari sistem operasi
	hostname, err := os.Hostname() //memanggil fungsi os.Hostname untuk mendapatkan hostname
	if err != nil {
		return "Unknown Hostname" //jika terjadi error, kembalikan string default
	}
	return hostname //mengembalikan hostname yang didapatkan
}
