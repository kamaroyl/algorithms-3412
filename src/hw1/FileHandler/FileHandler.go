package FileHandler

import (
        "bufio"
        "log"
        "os" 
        "strconv"
        "strings"
       )

func checkSuccess(e error) {
    if e != nil {
        log.Fatal("Failed to load")
        panic(e)
    }
}

// File is of format
// 1. description
// 2. size N
// 3.->N. value
func OpenIntsFile(filePath string, length int) []int {
    var ints []int
    var lengthTmp int
    var value int
    var intString string
    fp, err := os.Open(filePath)
    defer fp.Close()
    checkSuccess(err)
    scanner := bufio.NewScanner(fp)
    scanner.Split(bufio.ScanLines)
    success := scanner.Scan() 
    if success {
        if scanner.Err() == nil {
            log.Println("Reached EOF")
        }else {
            log.Fatal(scanner.Err())
            panic(scanner.Err())   
        }
    }
    log.Println(" Description: ", scanner.Text())
    success = scanner.Scan()
    intString = strings.TrimSuffix(scanner.Text(), "\n")
    lengthTmp, err = strconv.Atoi(intString)
    if lengthTmp < length { panic("Too long a lenght expected, cleaning up") }
    for i := 0; i < length; i++ {
        success = scanner.Scan()
        intString = strings.TrimSuffix(scanner.Text(), "\n")
        value, err = strconv.Atoi(intString)
        ints = append(ints, value)
    }
    return ints
}

