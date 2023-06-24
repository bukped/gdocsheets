package main

import "fmt"
 
func main() {
// Operator aritmatika
a := 10
b := 5
 
fmt.Println("Penjumlahan:", a+b)
fmt.Println("Pengurangan:", a-b)
fmt.Println("Perkalian:", a*b)
fmt.Println("Pembagian:", a/b)
fmt.Println("Sisa bagi:", a%b)
 
// Operator perbandingan
fmt.Println("Lebih besar:", a > b)
fmt.Println("Lebih kecil:", a < b)
fmt.Println("Sama dengan:", a == b)
fmt.Println("Tidak sama dengan:", a != b)
 
// Operator logika
x := true
y := false
 
fmt.Println("AND:", x && y)
fmt.Println("OR:", x || y)
fmt.Println("NOT:", !x)


// Operator Bitwise
// Operator bitwise AND (&)
fmt.Println("a & b:", a&b)   // 00000001 dalam biner
 
// Operator bitwise OR (|)
fmt.Println("a | b:", a|b)   // 00000111 dalam biner
 
// Operator bitwise XOR (^)
fmt.Println("a ^ b:", a^b)   // 00000110 dalam biner
 
// Operator bitwise shift left (<<)
fmt.Println("a << 2:", a<<2) // 00010100 dalam biner
 
// Operator bitwise shift right (>>)
fmt.Println("a >> 2:", a>>2) // 00000001 dalam biner

}
