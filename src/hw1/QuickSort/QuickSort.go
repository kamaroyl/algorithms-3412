package QuickSort

import (
        "fmt"
       )
/********************************************************************
 *             https://en.wikipedia.org/wiki/Quicksort              *
 *       https://www.geeksforgeeks.org/iterative-quick-sort/        *
 ********************************************************************/

func QuickSort(array []int) {
    low, high := 0, (len(array) -1)
    QuickSortFull(array, low, high)
}


func QuickSortFull(array []int, low int, high int) {
    comparison := 0
    swap := 0
    var stack [] int
    
    top := -1
    top++
    stack = append(stack, low)
    top++
    stack = append(stack, high)    

    //for == while in golang
    for top >= 0 {
        high = stack[top]
        stack = append(stack[:top], stack[top+1:]...)
        top--
        low  = stack[top]
        stack = append(stack[:top], stack[top+1:]...)
        top--
        pivot := array[high]
        p := low
        for j := low; j < high; j++ {
            if array[j] < pivot {
                comparison++
                //swap 
                array[p], array[j] = array[j], array[p]
                swap++
                p++ 
            }
        }
         //swap
        array[p], array[high] = array[high], array[p]
        swap++
        
        if p - 1 > low {
            stack = append(stack, low)
            top++
            stack = append(stack, (p - 1))
            top++
        }
        
        if p + 1 < high {
            stack = append(stack, (p + 1))
            top++
            stack = append(stack, high)
            top++
        }       
    }
    fmt.Println("Swaps: ", swap)
    fmt.Println("Comparisons: ", comparison)
}
