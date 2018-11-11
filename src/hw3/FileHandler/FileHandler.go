package FileHandler

import(
        "hw3/HeroDistanceMapping"
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
        "strings"
      )

func checkSuccess(e error) {
    if e!= nil {
        log.Fatal("failed to load file")
        panic(e)
    }
}

func ConstructCharacterToComicMatrix(filePath string) (*[]HeroDistanceMapping.HeroDistanceMap, *[][]uint8) {
    fp, err := os.Open(filePath)
    defer fp.Close()
    checkSuccess(err)
    scanner := bufio.NewScanner(fp)
    scanner.Split(bufio.ScanLines)
    scanner.Scan() // first line is *Verticies #verticies #edges
    header := strings.Split(strings.TrimSuffix(scanner.Text(), "\n"), " ")
    
    fmt.Println(header[1])
    totalVertecies, _ := strconv.Atoi(header[1])
    fmt.Printf("Total Vertecies: %d\n", totalVertecies)
    heroCount, _ := strconv.Atoi(header[2])
    fmt.Printf("Total Number of Super Heros: %d\n", heroCount)
    comicCount := totalVertecies - heroCount
    fmt.Printf("Total Number of Comics: %d\n", comicCount)
    //Initialize matrix of uint8 to represent comic book appearances
     //   dim = hero x comic

    matrix := make([][]uint8, comicCount + 1)
    for i := 0; i <= comicCount; i++ {
        matrix[i] = make([]uint8, heroCount + 1)
    }

    nameTable := make([]HeroDistanceMapping.HeroDistanceMap, totalVertecies + 1) // init super hero name lookup

    for i := 0; i < totalVertecies; i++ {
        scanner.Scan()
        nameAndId := strings.SplitN(strings.TrimSuffix(scanner.Text(), "\n"), " ", 2)
        id, _ := strconv.Atoi(nameAndId[0])
        nameTable[id] = HeroDistanceMapping.HeroDistanceMap { Name: nameAndId[1], Distance: -1 }
    }

    scanner.Scan() // Second header

    for i := 0; i <= comicCount; i++ {
        scanner.Scan()
        line := strings.TrimSuffix(scanner.Text(), "\n") 
        edgeList := strings.Split(line, " ")
        index, _ := strconv.Atoi(edgeList[0]) //hero
        for i := 1; i < len(edgeList); i++ {  //for list of comics
            jndex := translateComicIndex(edgeList[i], heroCount)
           // fmt.Printf("%v,%v\n",index, jndex)
            matrix[jndex][index] = 1
        }
    }
    
    return &nameTable, &matrix
}

func translateComicIndex( index string, heroCount int ) (int) {
    jndex, _ := strconv.Atoi(index)
    jndex = jndex - heroCount
    return jndex
}

