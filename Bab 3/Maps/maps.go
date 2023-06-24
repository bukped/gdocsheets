package main

import "fmt"
 
type Person struct {
Name   string
Age    int
Gender string
}



func main() {
// Membuat objek dari struct Person
p := Person{
Name:   "John Doe",
Age:    25,
Gender: "Male",
}
 
// Menampilkan nilai-nilai dari objek p
fmt.Println("Name:", p.Name)
fmt.Println("Age:", p.Age)
fmt.Println("Gender:", p.Gender)
}
