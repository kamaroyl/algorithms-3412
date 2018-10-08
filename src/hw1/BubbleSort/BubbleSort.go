package BubbleSort

import "fmt"

func BubbleSort(array [] int) {
    var comparison int = 0
    var assignment int = 0
    var sorted bool = true
    //golang's while loop is a for loop
    for true {
        sorted = true
        for i := 1; i < len(array); i++ {
            if array[i - 1] > array[i] {
                comparison++
                array[i - 1], array[i] = array[i], array[i - 1]
                assignment++
                sorted = false
                }
        }
        if sorted { break }
    }
    fmt.Println("Comparison count: ", comparison)
    fmt.Println("Assignment count: ", assignment) 
}
