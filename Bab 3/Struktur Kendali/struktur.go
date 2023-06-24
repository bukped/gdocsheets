package main

import "fmt"
 
func main() {
// Contoh penggunaan if-else
angka := 10
 
if angka > 5 {
fmt.Println("Angka lebih besar dari 5")
} else {
fmt.Println("Angka tidak lebih besar dari 5")	}
 
// Contoh penggunaan for loop
for i := 1; i <= 3; i++ {
fmt.Println("Perulangan ke", i)
}
 
// Contoh penggunaan switch
hari := "Senin"
 
switch hari {
case "Senin":
fmt.Println("Hari ini adalah Senin")
case "Selasa":
fmt.Println("Hari ini adalah Selasa")
default:
fmt.Println("Hari ini bukan Senin atau Selasa")
}
}
