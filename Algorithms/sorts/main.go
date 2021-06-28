package main

import  (
    "fmt"
    "sort"
    "sorts/mySorts"
)

func main() {
    // mySorts.BubbleSort
    srts := []func([]int) {
        mySorts.BubbleSort,
    }

    for _, fnc := range srts {
        testing(fnc)
    }

}

func areSlicesEqual(arr1, arr2 []int) bool {
    if len(arr1) != len(arr2) { return false }
    for i, _ := range arr1 {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}

func testing(srtFnc func([]int)) {
    /*Testcases*/
    testArrays := [][]int{
        {9,5,3,5,1,7},
        {8,5,2,25,6,3,7,3,4,6},
        {7,4,3,-14,0,8,-1,0,-3,-2,1},
        {1},
        {2,1},
        {},
        {8,5,2,25,6,3,7,3,4,6,4,7,2,2,3,24,4,8,5,9},
        {1,2,3,4,5}, /* Note the coma */
        } // Composite literals

    var failedTestcases []int

    fmt.Println("Amount of testcases: ", len(testArrays), "\n")

    for i, expectation, succes := 0, []int{}, ""; i < len(testArrays); i++ {
        succes = ""
        expectation = testArrays[i]
        sort.Sort(sort.IntSlice(expectation))
        fmt.Println("Testcase Nº", i, "\nArray Nº", i+1, "before sorting:\n", testArrays[i])
        srtFnc(testArrays[i])
        fmt.Println("Array Nº", i+1, "after sorting:\n", testArrays[i])
        fmt.Println("Expected:", expectation)
        if !areSlicesEqual(expectation, testArrays[i]) {
            succes = "not"
            failedTestcases = append(failedTestcases, i)
        }
        fmt.Println("Testcase has", succes, "been passed")
        fmt.Printf("\n\n")
    }

    // if len(failedTestcases) == 0 {
    if failedTestcases == nil {
        fmt.Printf("%sPassed all testcases%s\n", "\033[32m", "\033[0m")
    } else {
        fmt.Printf("%sFailed testcases\n%v%s\n", "\033[31m", failedTestcases, "\033[0m")

    }

}
