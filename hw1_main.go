package main
import (
       "hw1/BucketSort"
       "hw1/BubbleSort"
       "hw1/FileHandler"
       "hw1/QuickSort"
       "flag"
       "fmt"
       "time"
       )

const( 
        BUBBLE_SORT = iota
        QUICK_SORT = iota
        BUCKET_SORT = iota
)

func runTime(fp func([]int), arg []int) {
    start := time.Now()
    fp(arg)
    finish := time.Now()
    fmt.Println("Total Time: ", finish.Sub(start))
}

func main() {
    
    var ints []int
    var debug  = flag.Bool("debug", false, "debug print options")
    var count  = flag.Int("count", 1000, "number of items to parse")    
    var file   = flag.String("file", "hw1_resource/duplicate.txt", "location of file of integers")
    var method = flag.Int("method", 0, "Bubble = 0, Quick = 1, Bucket = 2")
    flag.Parse()
    
    ints = FileHandler.OpenIntsFile(*file, *count)
    if *method == BUBBLE_SORT {
        runTime(BubbleSort.BubbleSort, ints)
    }

    if *method == QUICK_SORT {
        runTime(QuickSort.QuickSort, ints)
    }
    
    if *method == BUCKET_SORT {
        runTime(BucketSort.BucketSort, ints)
    }

    if *debug { 
        fmt.Println( "Sorted: ", ints)
        fmt.Printf("len=%d cap=%d\n", len(ints), cap(ints))
    }    
}
