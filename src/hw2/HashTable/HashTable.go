package HashTable

import (
        "fmt"
        "hw2/MyMapImpl"
       )

func BuildConcordance(words *[]*string) *MyMapImpl.MyMap {
    var comparisons = make(map[int]int)
    var concordance = MyMapImpl.NewMyMap()
    for _, wordPtr := range *words {
        tmp := concordance.At(*wordPtr)
        if *tmp > 0 {
            (*tmp)++
            comparison := concordance.Add(*wordPtr, *tmp)
            tmpValue := comparisons[concordance.Len()]
            if tmpValue > 0 {
                comparisons[concordance.Len()] = (tmpValue + comparison)/2
            } else {
                comparisons[concordance.Len()] = comparison
            }
        } else {
            comparison := concordance.Add(*wordPtr, 1)
            tmpValue := comparisons[concordance.Len()]
            if tmpValue > 0 {
                comparisons[concordance.Len()] = (tmpValue + comparison)/2
            } else {
                comparisons[concordance.Len()] = comparison
            }
        }
    }
    for key, value := range comparisons {
        fmt.Println(key, ",", value)
    }
    return concordance
}
