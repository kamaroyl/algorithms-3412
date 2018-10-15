package main
import (
       "hw2/FileHandler"
       "hw2/HashTable"
       "hw2/MyMapImpl"
       "hw2/OrderedVector"
       "hw2/UnorderedVector"
       "hw2/WordFrequency"
       "flag"
       "fmt"
       "time"
       )

const(
    UNORDERED_VECTOR = iota
    ORDERED_VECTOR = iota
    HASH_MAP = iota
)

func runTime(fp func(*[]*string)*[]*WordFrequency.WordFrequency, words *[]*string) *[]*WordFrequency.WordFrequency {
    start := time.Now()
    var val = fp(words)
    finish := time.Now()
    fmt.Println("Total Time: ", finish.Sub(start))
    return val
}

func runTimeHashMap(fp func(*[]*string)*MyMapImpl.MyMap, words *[]*string) *MyMapImpl.MyMap {
    start := time.Now()
    var val = fp(words)
    finish := time.Now()
    fmt.Println("Total Time: ", finish.Sub(start)) 
    return val
}

func print(words *[]*WordFrequency.WordFrequency ) {
    var printWords = [...]string{"a", "a'", "a's", "a-bat-fowling", "a-bed", "a-birding", "a-bleeding", "a-breeding", "a-brewing", "a-broach", "zenelophon", "zenith", "zephyrs", "zo", "zodiac", "zodiacs", "zone", "zounds", "zur", "zwaggered"}
    for _, element := range *words { 
        for _, known := range printWords {
            if (*element).Word == known {
                fmt.Println(*element)
            }
        }   
    }
   
     
}

func main() {
    var words []*string
    var debug  = flag.Bool("debug", false, "debug print options")
    var file   = flag.String("file", "hw2_resource/wordlist.txt", "location of word list of words")
    var method = flag.Int("method", 1, "Unorder Vector = 0, Ordered Vector = 1, Hash Map = 2")
    flag.Parse()
    
    words = FileHandler.OpenStringsFile(*file)
    fmt.Println("All words parsed")
    if *debug {
        fmt.Printf("len: %d cap: %d\n", len(words), cap(words))
    } 
    fmt.Println("Method is ", *method)
    
    if *method == 0 {
        var concordance *[]*WordFrequency.WordFrequency =  runTime(UnorderedVector.BuildConcordance, &words)
        if *debug {
            print(concordance)
        }
    }

    if *method == 1 {
        var concordance *[]*WordFrequency.WordFrequency = runTime(OrderedVector.BuildConcordance, &words)
        if *debug {
            print(concordance)
        } 
    }
    
    if *method == 2 {
        var concordance *MyMapImpl.MyMap = runTimeHashMap(HashTable.BuildConcordance, &words)
        
        if *debug {
            fmt.Println("a: ", (*concordance).At("a"))
            fmt.Println("a': ", (*concordance).At("a'"))
            fmt.Println("a's: ", (*concordance).At("a's")) 
            fmt.Println("a-bat-fowling: ", (*concordance).At("a-bat-fowling"))
            fmt.Println("a-bed: ", (*concordance).At("a-bed"))
            fmt.Println("a-birding: ", (*concordance).At("a-birding"))
            fmt.Println("a-bleeding: ", (*concordance).At("a-bleeding"))
            fmt.Println("a-breeding: ", (*concordance).At("a-breeding"))
            fmt.Println("a-brewing: ", (*concordance).At("a-brewing"))
            fmt.Println("a-broach: ", (*concordance).At("a-broach"))
            fmt.Println("zenelophon: ", (*concordance).At("zenelophon"))
            fmt.Println("zenith: ", (*concordance).At("zenith"))
            fmt.Println("zephyrs: ", (*concordance).At("zephyrs"))
            fmt.Println("zo: ", (*concordance).At("zo"))
            fmt.Println("zodiac: ", (*concordance).At("zodiac"))
            fmt.Println("zodiacs: ", (*concordance).At("zodiacs"))
            fmt.Println("zone: ", (*concordance).At("zone"))
            fmt.Println("zounds", (*concordance).At("zounds"))
            fmt.Println("zur: ", (*concordance).At("zur"))
            fmt.Println("zwaggered: ", (*concordance).At("zwaggered"))
            /*for key, value := range *concordance {
                fmt.Println("{ ", key, " ", value, " }")
            }*/
        } 
    }
}
