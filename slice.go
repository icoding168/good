package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)

   printSlice(numbers)
   
   s := []int{7, 2, 8, -9, 4, 0}
   s2 := s[:len(s)/2]
   printSlice(s2)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}