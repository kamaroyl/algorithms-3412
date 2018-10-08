package HashTable

import "fmt"

func BuildConcordance(words *[]*string) *map[string]int {
    var comparison = 0
    var assignment = 0
    var concordance = make(map[string]int)
    for _, wordPtr := range *words {
        tmp := concordance[*wordPtr]
        comparison ++
        if tmp > 0 {
            tmp++
            concordance[*wordPtr] = tmp
        } else {
        assignment ++
            concordance[*wordPtr] = 1
        }
    }
    fmt.Println("comparison ", comparison, " assignment ", assignment)
    return &concordance
}
