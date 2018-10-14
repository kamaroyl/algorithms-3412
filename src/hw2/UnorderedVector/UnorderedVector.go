package UnorderedVector

import(
      "hw2/WordFrequency"
      "fmt"
      )


/******************************************
 *  O(n) addition or frequency increase
 *      - Returns comparison count, assignment count
 ******************************************/
func searchWord(word *string, concordancePtr *[]*WordFrequency.WordFrequency) (int, int) {
    var comparison = 0
    var length = len(*concordancePtr)
    
    for _, existingWord := range *concordancePtr {
        comparison++
        if *word == (*existingWord).Word {
            (*existingWord).Frequency++
            return comparison, length
        }
    }
    tmp := &WordFrequency.WordFrequency{Word: *word, Frequency: 1 }
    (*concordancePtr)= append((*concordancePtr), tmp)
    return comparison, length
}

func BuildConcordance(words *[]*string) *[]*WordFrequency.WordFrequency{
    var comparison = make(map[int]int) 
    var concordance []*WordFrequency.WordFrequency
    for _, word := range *words {
        tmpComparison, tmpLength :=  searchWord(word, &concordance)
        tmpEntry := comparison[tmpLength]
        if tmpEntry > 0 {
            comparison[tmpLength] = (tmpEntry + tmpComparison)/2
        } else {
            comparison[tmpLength] = tmpComparison
        }
    }
    for key, value := range comparison {
        fmt.Println(key,",",value)
    }
    return &concordance
}
