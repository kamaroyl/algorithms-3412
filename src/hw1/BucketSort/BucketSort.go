package BucketSort

import "fmt"


/******************************************************************
 *      Bucket sort runs insertion sort after recombining the     *
 *      Buckets -> optimization from wikipedia article            *
 ******************************************************************/
func BucketSort(ints []int) {
    numberOfBuckets := 1000
    BucketSortFull(ints, numberOfBuckets)
}

func BucketSortFull(ints []int, numberOfBuckets int) {
    var hashKey int

    // Number of swaps in sorting and assignments
    swaps := 0

    // Number of comparisons made
    comparisons := 0

    bucket := make([][]int, numberOfBuckets)
    for i := range bucket {
        bucket[i] = make([]int, 0)
    }
 
    for i := range ints {
        hashKey = ints[i]
        bucket[hashKey] = append(bucket[hashKey], ints[i])
        swaps++   
    }
    k := 0
    for i := range bucket {
        for j := range bucket[i] {
            ints[k] = bucket[i][j]
            swaps++
            k++
        }
    }
    
    //https://www.geeksforgeeks.org/insertion-sort/
    //https://en.wikipedia.org/wiki/Insertion_sort/
    for i := 1; i < len(ints); i++ {
        val := ints[i]
        j := i - 1
        comparisons++
        for j >= 0 && ints[j] > val {
            comparisons++
            ints[j + 1] = ints[j]
            j = j - 1
        }
        ints[j + 1] = val
    }

    fmt.Println("Comparisons: ", comparisons)
    fmt.Println("Swaps: ", swaps)
}
