package main

import "fmt"

func someWrong() error {
    return fmt.Errorf("sss")
}

func main() {
    s0 := make([]int32, 10)
    s1 := append(s0, 12)
    fmt.Printf("len %d cap %d\n%v\n", len(s1), cap(s1), s1)
    s1 = append(s1, 12)
    fmt.Printf("len %d cap %d\n%v\n", len(s1), cap(s1), s1)
    ps0 := &s0
    pps0 := &ps0
    fmt.Printf("ps0 %v\npps0 %v\n", *ps0, **pps0)
    //err := someWrong()
    //fmt.Printf("%s", err)
    
    // array & slice
    arr := [...]string{"hello", "world"}
    fmt.Printf("arr.len = %d, arr.cap = %d\n", len(arr), cap(arr))
    
    slice := new([2]int32)[:]
    slice[0], slice[1] = 123, 456
    fmt.Printf("slice = %v\n", slice)
}