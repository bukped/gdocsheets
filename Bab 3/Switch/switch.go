package main

import "fmt"
 
func main() {
var nilai int = 90
//  contoh switch single case
switch nilai {
case 90:
fmt.Println("Nilai sangat baik!")
}

// contoh switch multi case
switch nilai {
case 90:
fmt.Println("Nilai sangat baik!")
case 80, 85:
fmt.Println("Nilai baik!")
case 70, 75:
fmt.Println("Nilai cukup baik.")
default:
fmt.Println("Nilai standar.")

}
}
