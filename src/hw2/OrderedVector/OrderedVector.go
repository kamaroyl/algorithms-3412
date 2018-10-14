package OrderedVector

import(
       "hw2/WordFrequency"
       "fmt"
       "math"
      )

/*
   O(lg(n)) if the entry already exists, O(n*lg(n)) if entry doesn't due
   to the need to copy the array to the later peice of the slice
 */
func searchWord(word *string, s *[]*WordFrequency.WordFrequency) (int, int){
    comparison := 0
    length := len(*s)
    if len(*s) == 0 {
        comparison++
        tmp := &WordFrequency.WordFrequency{ Word: *word, Frequency: 1 }
        *s = append(*s, tmp)
        return comparison, length     
    }

    low  := 0
    high :=  len(*s) - 1
    middle := 0
    for low <= high {
        comparison++
        middle = int(math.Floor(float64((low + high)/2)))
        if (*((*s)[middle])).Word < *word {
            low = middle + 1
        } else if (*((*s)[middle])).Word > *word {
            high = middle - 1
        } else {
            (*((*s)[middle])).Frequency++
            return comparison, length
        }
    }
    tmp := &WordFrequency.WordFrequency{ Word: *word, Frequency: 1}
    if (*((*s)[middle])).Word < *word {
        for (middle < len(*s) && (*((*s)[middle])).Word < *word) {
            middle++
        }
        insertIntoSliceAtPosition(s, tmp, middle)
    } else {
        for (middle > 0 && (*((*s)[middle - 1])).Word > *word) {
            middle--
        }
        insertIntoSliceAtPosition(s, tmp, middle)
        }
    

    return comparison, length
}

//From https://github.com/golang/go/wiki/SliceTricks
func insertIntoSliceAtPosition(s *[]*WordFrequency.WordFrequency, add *WordFrequency.WordFrequency, index int){
    if index < 0 {
        *s = append([]*WordFrequency.WordFrequency{add}, (*s)...)
        return
    }
    
    if index > len(*s) {
        *s = append(*s, add)
        return
    }

    *s = append(*s, &WordFrequency.WordFrequency{Word: "", Frequency: 0})
    copy((*s)[index + 1:], (*s)[index:])
    (*s)[index] = add
}

func BuildConcordance(words *[]*string) *[]*WordFrequency.WordFrequency {
    var comparison = make(map[int]int)
    var concordance []*WordFrequency.WordFrequency
    for _, word := range *words {
        tmpComparison, tmpLength := searchWord(word, &concordance)
        tmpEntry := comparison[tmpLength]
        if tmpEntry > 0 {
            comparison[tmpLength] = (tmpEntry + tmpComparison)/2
        } else {
            comparison[tmpLength] = tmpComparison
        }
    }
    for key, value := range comparison {
        fmt.Println(key, ",", value)
    }
    return &concordance
}

