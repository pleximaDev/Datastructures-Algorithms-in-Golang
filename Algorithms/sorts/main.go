package main

import  (
    "fmt"
    "sort"
    S "./mySorts"
)

func main() {

    /*Testcases*/
    testArrays := [][]int{
        {9,5,3,5,1,7},
        {8,5,2,25,6,3,7,3,4,6},
        {1},
        {2,1},
        {},
        {1,2,3,4,5}, /* Note the coma */
        } // Composite literals

    fmt.Println("Amount of testcases: ", len(testArrays))

    for i, expectation, succes := 0, []int{}, ""; i < len(testArrays); i++ {
        expectation = testArrays[i]
        sort.Sort(sort.IntSlice(testArrays[i]))
        fmt.Println("Testcase Nº", i, "\nArray Nº", i+1, "before sorting:\n", testArrays[i])
        S.BubbleSort(testArrays[i])
        fmt.Println("Array Nº", i+1, "after sorting:\n", testArrays[i])
        fmt.Println("Expected:", expectation)
        if !areSlicesEqual(expectation, testArrays[i]) {succes = "not"}
        fmt.Println("Testcase has", succes, "been passed")
        fmt.Printf("\n\n")
    }
}


func areSlicesEqual(arr1, arr2 []int) bool {
    if len(arr1) != len(arr2) { return false }
    for i := range arr1 {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
