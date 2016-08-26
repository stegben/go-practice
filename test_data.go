package main

import "fmt"

func main() {
    // const pi float64 = 3.1415
    // fmt.Println("hi")
    // fmt.Printf("%.2f is pi \n", pi)

    // for i := 0; i <= 10; i ++ {
    //  fmt.Println(i)
    // }

    var arrayTest[3] float64

    arrayTest[0] = 1.01
    arrayTest[1] = 2.3
    arrayTest[2] = 3

    fmt.Println(arrayTest)

    arrayAnother := [4]float64 {1, 2, 3, 4}

    fmt.Println(arrayAnother)

    sliceNum := []uint {1, 2, 3, 4}
    fmt.Println(sliceNum[2:4])
}