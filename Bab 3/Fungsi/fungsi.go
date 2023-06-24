package main

import "fmt"

// Definisikan fungsi sederhana
func sayHello() {
fmt.Println("Halo, dunia!")
}
 
// Definisikan fungsi dengan parameter
func greet(name string) {
fmt.Println("Halo,", name)
}



 // Definisikan fungsi dengan return value
func addNumbers(a, b int) int {
return a + b
}

func main() {
// Panggil fungsi sayHello
sayHello()
 
// Panggil fungsi greet dengan argument
greet("John")
 
// Panggil fungsi addNumbers dengan argument dan cetak hasilnya
result := addNumbers(5, 3)
fmt.Println("Hasil penjumlahan:", result)
}
