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
    var assignment = 0
    
    for _, existingWord := range *concordancePtr {
        comparison++
        if *word == (*existingWord).Word {
            (*existingWord).Frequency++
            return comparison, assignment
        }
    }
    assignment++
    tmp := &WordFrequency.WordFrequency{Word: *word, Frequency: 1 }
    (*concordancePtr)= append((*concordancePtr), tmp)
    return comparison, assignment
}

func BuildConcordance(words *[]*string) *[]*WordFrequency.WordFrequency{
    var comparison = 0
    var assignment = 0
    var concordance []*WordFrequency.WordFrequency
    for _, word := range *words {
     tmpComparison, tmpAssignment :=  searchWord(word, &concordance)
     comparison+=tmpComparison
     assignment+=tmpAssignment
    }
    fmt.Println("comparison ", comparison, " assignment ", assignment)   
    return &concordance
}
