package main

import "fmt"
 
func main() {
// Variabel umur
umur := 18
// variabel nilai
nilai := 85
// variabel tinggi
tinggi := 165
 
// If Statemet
if umur >= 18 {
fmt.Println("Anda sudah cukup umur")
}
// if else statement
if umur < 18{
	fmt.Println("Anda masih muda")
}else {
	fmt.Println("Anda sudah cukup umur")
	}

// if-else-if statement
if nilai >= 90 {
	fmt.Println("Anda mendapatkan nilai A")
	} else if nilai >= 80 {
	fmt.Println("Anda mendapatkan nilai B")
	} else if nilai >= 70 {
	fmt.Println("Anda mendapatkan nilai C")
	} else {
	fmt.Println("Anda mendapatkan nilai D")
	}
	
// nested if
if umur >= 18 {
	fmt.Println("Anda sudah cukup umur")
	if tinggi >= 170 {
		fmt.Println("Anda juga cukup tinggi")
	} else {
	fmt.Println("Anda kurang tinggi")
	}
	}	
}
