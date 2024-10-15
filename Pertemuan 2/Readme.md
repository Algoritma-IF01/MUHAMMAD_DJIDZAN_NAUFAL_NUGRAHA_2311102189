# <h1 align="center"> Laporan Praktikum_2 - Modul 2 - Review Struktur Kontrol </h1>

## <p align="center"> Muhammad Djidzan Naufal Nugraha </p>

## <p align="center"> 2311102189 </p>

# Guided

### Guided 1

**Source Code :**
```go
package main

import (
	"fmt"
)

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
```
### Ouput Code
<img width="960" alt="Guided1" src="https://github.com/user-attachments/assets/0634fc0d-2273-4f97-9c4f-ae34ce0e4973">

### Guided 2

**Source Code :**
```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Masukkan nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scanln(&c)

	hipotenusa = (c * c) == (a*a + b*b)

	fmt.Println("Apakah sisi c adalah hipotenusa segitiga a,b,c: ", hipotenusa)
}
```

### Ouput Code
<img width="960" alt="Guided2" src="https://github.com/user-attachments/assets/0ddf1ed8-63b0-4ad4-b2b8-3db089c33428">

# Latihan

### Latihan 1
**Source Code :**
```go
package main

import (
	"fmt"
)

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```

### Ouput Code
<img width="960" alt="Latihan1" src="https://github.com/user-attachments/assets/58a58e68-cb2d-44d3-836d-e937feca1c69">

### Latihan 2
**Source Code :**
```go
package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukan Tahun: ")
	fmt.Scanln(&tahun)

	kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Printf("Kabisat: %t", kabisat)
}
```

### Ouput Code
<img width="960" alt="Latihan2" src="https://github.com/user-attachments/assets/526010d2-6034-4e9f-8286-16eeee0b23fd">

### Latihan 3
**Source Code :**
```go
package main

import "fmt"

func main() {
	var r int
	fmt.Print("Jari-Jari: ")
	fmt.Scanln(&r)

	pi := 3.1415926535
	volume := (4.0 / 3.0) * pi * float64(r*r*r)
	luas := 4 * pi * float64(r*r)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}
```

### Ouput Code
<img width="960" alt="Latihan3" src="https://github.com/user-attachments/assets/092b7d1c-cc2b-4d55-995d-49990fb59a04">

### Latihan 4
**Source Code :**
```go
package main

import "fmt"

func main() {
	var (
		celcius, fahrenheit, kelvin, reamur float64
	)
	fmt.Print("Masukan Temperatur Celsius: ")
	fmt.Scanln(&celcius)

	fahrenheit = (9.0/5.0)*celcius + 32
	reamur = (4.0 / 5.0) * celcius
	kelvin = celcius + 273.15

	fmt.Printf("Derajat Reamur %.2f \n", reamur)
	fmt.Printf("Derajat Fahrenheit %.2f \n", fahrenheit)
	fmt.Printf("Derajat Kelvin %.2f \n", kelvin)
}
```

### Ouput Code
<img width="960" alt="Latihan4" src="https://github.com/user-attachments/assets/a722115a-713b-4630-85f5-d47e59153706">

### Latihan 5
**Source Code :**
```go
package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var num int

	fmt.Print("Masukkan 5 angka dipisahkan spasi: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&num)
		fmt.Printf("%c", num)
	}
	fmt.Println()

	var input string
	fmt.Print("Masukkan 3 karakter: ")
	fmt.Scan(&input)

	if len(input) <= 3 {
		fmt.Print("Hasil konversi: ")
		for i := 0; i < len(input); i++ {
			fmt.Printf("%c", input[i]+1)
		}
		fmt.Println()
	}
}
```

### Ouput Code
<img width="960" alt="Latihan5" src="https://github.com/user-attachments/assets/1bb161eb-e939-447c-9b7b-3941d21f1041">
