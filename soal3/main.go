package main

import "fmt"

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b *Bebek) {
	b.hidup = false
}
func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}
func Makan(b *Bebek) {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}

func main() {
	// Contoh penggunaan
	bebek := Bebek{
		energi:       5,
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "Kwek kwek kwek!",
	}

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi) 

	Makan(&bebek)
	fmt.Println("Energi setelah makan:", bebek.energi)

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi) 

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi) 

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi) 

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi) 
	fmt.Println("apakah bebek masih hidup?", bebek.hidup)

	Terbang(&bebek)                                      
	fmt.Println("Energi setelah terbang:", bebek.energi)
	fmt.Println("apakah bebek masih hidup?", bebek.hidup)
}