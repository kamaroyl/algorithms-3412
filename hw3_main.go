package main
import (
       "hw3/Collaboration" 
       "hw3/FileHandler"
       "hw3/HeroDistanceMapping"
       "flag"
       "fmt"
       "strings"
       )
/*
func runTime(fp func()) {
    start := time.Now()
    var val = fp()
    finish := time.Now()
    fmt.Println("Total Time: ", finish.Sub(start))
    return val
}*/

func main() {
    var debug  = flag.Bool("debug", false, "debug print options")
    var file   = flag.String("file", "hw3_resource/porgat.txt", "location of hero graph")
    var superHeros = flag.String("superHeros", "all", "which super heros to look up information for")
    flag.Parse()
    var heroLut *[]HeroDistanceMapping.HeroDistanceMap
    var nameToId = make(map[string]int)
    var comicbookXHero *[][]uint8
    spiderMan := 5306
    heroLut, comicbookXHero = FileHandler.ConstructCharacterToComicMatrix(*file)
    
    collabMatrix := Collaboration.MakeCollaboration(comicbookXHero)     
    spiderManVector := Collaboration.ExtractSuperHeroVector(collabMatrix, spiderMan)
    dim, _ := spiderManVector.Dims()
    
    for i := 1; i < dim; i++ {
        nameToId[strings.ToLower((*heroLut)[i].Name)] = i
    }
    fmt.Printf("%v",nameToId)
    fmt.Printf("RAI is %v\n", nameToId["gaia"]) 
    for i := 0; i < dim; i++ {
        if spiderManVector.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 1
            }
        }
    }

    tmp1 := Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, spiderManVector) 

    for i := 0; i < dim; i++ {
        if tmp1.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 2
            }
        }
    }
    
    tmp2 := Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, tmp1) 

    for i := 0; i < dim; i++ {
        if tmp2.At(i, 0) > 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 3
            }
        }
    }

    tmp1 = Collaboration.MultiplySuperHeroMatrixByVector(collabMatrix, tmp2) 

    for i := 0; i < dim; i++ {
        if tmp1.At(i, 0) < 0 {
            if (*heroLut)[i].Distance < 0 {
                (*heroLut)[i].Distance = 4
            }
        }
    }
    if *superHeros == "all" {
        for i := 0; i < dim; i++ {
            fmt.Printf("%v has a spiderman number of %v and is in %v comics with him\n", (*heroLut)[i].Name, (*heroLut)[i].Distance, spiderManVector.At(i, 0))
        }
    } else {
        heroList := strings.Split(*superHeros, ",")
        for hero := range heroList {
            value :=  (*heroLut)[nameToId[heroList[hero]]]
            fmt.Printf("%v has a spiderman number of %v\n", value.Name, value.Distance)
        }
    }
    fmt.Println("Program Complete")
    if *debug {
    } 
}
