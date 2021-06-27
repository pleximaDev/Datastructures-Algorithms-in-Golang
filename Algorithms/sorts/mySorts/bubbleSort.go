package mySorts

func  BubbleSort(l []int) /*[]int*/ {
    var n int = len(l)
    var swapped bool = true
    for swapped {
        swapped = false
        for i := 1; i < n - 1; i++ {
            if l[i-1] > l[i] {
                swap(&l[i-1], &l[i])
                swapped = true
                // swapArrIndx(l, j) // swapping l[j] and l[j+1]
                // l[j], l[j+1] = l[j+1], l[j] // swap via tuple assignment
            }
        }
    }
}

func swapArrIndx(l []int, i int) /*int*/ {
	tmp := l[i]
    l[i] = l[i+1]
    l[i+1] = tmp
}

func swap(a, b *int) {
    tmp := *a
    *a = *b
    *b = tmp
}
