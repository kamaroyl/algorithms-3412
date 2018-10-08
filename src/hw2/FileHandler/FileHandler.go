package FileHandler

import(
        "hw2/WordFrequency" 
        "bufio"
        "fmt"
        "log"
        "os"
        "strings"
        "unicode"
      )

func checkSuccess(e error) {
    if e!= nil {
        log.Fatal("failed to load file")
        panic(e)
    }
}

func startsWithLetter(s *string) bool {
    if len(*s) > 0 {
        r := rune((*s)[0])
        return unicode.IsLetter(r)
    } 
    return false
}

func OpenStringsFile(filePath string) []*string {
    var words []*string
    fp, err := os.Open(filePath)
    defer fp.Close()
    checkSuccess(err)
    scanner := bufio.NewScanner(fp)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        word := strings.TrimSuffix(scanner.Text(), "\n")
        if startsWithLetter(&word) {
            word = strings.ToLower(word)
            words = append(words, &word)
        }
    }

   
    if err := scanner.Err(); err != nil {
        fmt.Println("Error")
    }
    
    return words
}

func WriteUnorderedVector(filePath string, unorderedVector *[]*WordFrequency.WordFrequency) {

}
