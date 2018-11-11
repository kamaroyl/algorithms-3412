package main
import (
       "hw3/Collaboration" 
       "hw3/FileHandler"
       "hw3/HeroDistanceMapping"
       "flag"
       "fmt"
       "strconv"
       "strings"
       "time"
       )


func main() {
    startTotal := time.Now()
    var file   = flag.String("file", "hw3_resource/porgat.txt", "location of hero graph")
    var superHeros = flag.String("superHeros", "all", "which super heros to look up information for")
    flag.Parse()
    var heroLut *[]HeroDistanceMapping.HeroDistanceMap
    var nameToId = make(map[string]int)
    var comicbookXHero *[][]uint8
    spiderMan := 5306

    start := time.Now()
    heroLut, comicbookXHero = FileHandler.ConstructCharacterToComicMatrix(*file)
    finish := time.Now()

    fmt.Printf("Took %v to construct character to comic matrix\n", finish.Sub(start))
   
    start = time.Now()
    collabMatrix := Collaboration.MakeCollaboration(comicbookXHero)
    finish = time.Now()
    fmt.Printf("Took %v to construct the collaboration matrix\n", finish.Sub(start))

    start = time.Now()
    spiderManVector := Collaboration.ExtractSuperHeroVector(collabMatrix, spiderMan)
    finish = time.Now()
    fmt.Printf("Took %v to construct character to comic vector\n", finish.Sub(start))

    dim, _ := spiderManVector.Dims()
    
    for i := 1; i < dim; i++ {
        nameToId[strings.ToLower((*heroLut)[i].Name)] = i
    }

    start = time.Now()

    for i := 0; i < dim; i++ {
        if spiderManVector.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 1
            }
        }
    }
    finish = time.Now()
    fmt.Printf("Took %v to perform initial look up table assignments\n", finish.Sub(start))

    start = time.Now()
    tmp1 := Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, spiderManVector) 

    for i := 0; i < dim; i++ {
        if tmp1.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 2
            }
        }
    }
    finish = time.Now()
    fmt.Printf("Took %v to perform matrix multiplictation and table assignment\n", finish.Sub(start))

    start = time.Now()
    tmp2 := Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, tmp1) 

    for i := 0; i < dim; i++ {
        if tmp2.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 3
            }
        }
    }
    finish = time.Now()
    fmt.Printf("Took %v to perform second round of matrix multiplictation and table assignment\n", finish.Sub(start))


    start = time.Now()
    tmp1 = Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, tmp2) 

    for i := 0; i < dim; i++ {
        if tmp1.At(i, 0) < 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 4
            }
        }
    }
    finish = time.Now()
    fmt.Printf("Took %v to perform third round of matrix multiplictation and table assignment\n", finish.Sub(start))
  
    if *superHeros == "all" {
        for i := 0; i < dim; i++ {
            fmt.Printf("%v has a spiderman number of %v and is in %v comics with him\n", (*heroLut)[i].Name, (*heroLut)[i].Distance, spiderManVector.At(i, 0))
        }
    } else {
        heroList := strings.Split(*superHeros, ",")
        for hero := range heroList {
            id, _ := strconv.Atoi(heroList[hero])
            value :=  (*heroLut)[id]
            fmt.Printf("%v has a spiderman number of %v\n", value.Name, value.Distance)
        }
    }
    fmt.Println("Program Complete")
    finishTotal := time.Now()
    fmt.Printf("Total Run Time: %v\n", finishTotal.Sub(startTotal))

}
