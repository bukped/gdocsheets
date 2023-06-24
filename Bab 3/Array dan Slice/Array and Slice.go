package main

import "fmt"
 
func main() {
// Contoh penggunaan array
var fruits [3]string
fruits[0] = "Apple"
fruits[1] = "Banana"
fruits[2] = "Orange"
 
fmt.Println(fruits) // Output: [Apple Banana Orange]
 
// Contoh penggunaan slices
numbers := []int{1, 2, 3, 4, 5}

              fmt.Println(numbers)         // Output: [1 2 3 4 5]
fmt.Println(numbers[1:3])    // Output: [2 3]
fmt.Println(numbers[:4])     // Output: [1 2 3 4]
fmt.Println(numbers[3:])     // Output: [4 5]
fmt.Println(numbers[:])      // Output: [1 2 3 4 5]
  
             // Mengubah elemen dalam slices
numbers[0] = 10
fmt.Println(numbers)         // Output: [10 2 3 4 5]


   
// Menggunakan fungsi append untuk menambahkan elemen ke slices
numbers = append(numbers, 6)
fmt.Println(numbers)         // Output: [10 2 3 4 5 6]
}

