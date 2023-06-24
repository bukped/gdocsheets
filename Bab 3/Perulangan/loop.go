package main

import "fmt"
 
func main() {
// Contoh penggunaan for loop
for i := 1; i <= 5; i++ {
fmt.Println(i)
}
// Contoh penggunaan while loop
j := 1
for j <= 3 {
fmt.Println(j)
j++
}
// Contoh penggunaan infinite loop
k := 1
for {
fmt.Println(k)
k++
if k > 5 {
break
}
}
}
