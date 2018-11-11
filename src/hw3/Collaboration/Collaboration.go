package Collaboration

import( 
      "fmt"
      "github.com/james-bowman/sparse"
      )

func compressComicXHerosRow( comicXHeros *[]uint8 ) (*[]int) {
    var compressed []int
    for i := range *comicXHeros {
        if (*comicXHeros)[i] == 1 {
            compressed = append(compressed, i)
        }
    }
    return &compressed
}

func fillCollaboration(heros *[]int, collaboration *sparse.DOK) {
    
    for i := range *heros {
        x := (*heros)[i]
        for j := range *heros {
            y := (*heros)[j]
            //fmt.Printf("%v, %v: %v\n",x,y, (*collaboration).At(x,y))
            tmp := (*collaboration).At(x, y) + 1
            //fmt.Printf("%v, %v: %v\n",x,y,tmp)
            (*collaboration).Set(x, y, tmp)
        }
    }

}

func MakeCollaboration(comicXHeros *[][]uint8)(*sparse.CSR) {
    dim := len((*comicXHeros)[0])
    collaborationMatrix := sparse.NewDOK(dim, dim)

    for i := range *comicXHeros {
        fillCollaboration(compressComicXHerosRow((&(*comicXHeros)[i])), collaborationMatrix)
    }
    return (collaborationMatrix.ToCSR())
}

func ExtractSuperHeroVector(collabMatrix *sparse.CSR, superHero int) (*sparse.CSR){
    dim, _ := collabMatrix.Dims()
    superHeroVector := sparse.NewDOK(dim, 1)
    for i := 0; i < dim; i++ {
        superHeroVector.Set(i, 0, collabMatrix.At(superHero, i))
        /*if superHeroVector.At(superHero, 0) > 0 {
             fmt.Printf("%v: %v\n", i, superHeroVector.At(superHero,0)) 
        }*/
    }
    fmt.Printf("REGULAR: %v: %v\n", 5306, superHeroVector.At(5306,0))
    csrSH := superHeroVector.ToCSR()
    fmt.Printf("COMPRESSED: %v: %v\n", 5306, csrSH.At(5306,0)) 
    return csrSH
}

func MultiplySuperHeroMatrixByVector(collabMatrix *sparse.CSR, superVector *sparse.CSR) (*sparse.CSR) {
    var result sparse.CSR
    result.Mul(collabMatrix, superVector)
    return &result
}
