package main

import (
"fmt"
"math/rand"
"sync"
"time"
)

func mergeSortConWG(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr)/2
	var lArr,rArr []int
	var wgTemp sync.WaitGroup
	wgTemp.Add(1)
	go func() {
		lArr = mergeSortConWG(arr[:mid])
		wgTemp.Done()
	}()
	wgTemp.Add(1)
	go func() {
		rArr = mergeSortConWG(arr[mid:])
		wgTemp.Done()
	}()
	wgTemp.Wait()
	sortedArr := mergeTheTwoArrayConWG(lArr, rArr)

	return sortedArr
}

func mergeTheTwoArrayConWG(lArr, rArr []int) []int {
	totalLen := len(lArr) + len(rArr)
	newArr := make([]int, totalLen)
	i := 0
	l := 0
	r := 0
	for(l<len(lArr) && r <len(rArr)) {
		if lArr[l] < rArr[r] {
			newArr[i] = lArr[l]
			i++
			l++
		} else if lArr[l] > rArr[r] {
			newArr[i] = rArr[r]
			i++
			r++
		} else {
			newArr[i] = lArr[l]
			i++
			l++
			newArr[i] = rArr[r]
			i++
			r++
		}
	}

	for (l<len(lArr)) {
		newArr[i] = lArr[l]
		l++
		i++
	}
	for (r<len(rArr)) {
		newArr[i] = rArr[r]
		r++
		i++
	}
	return newArr
}

func main() {
	//arr := []int{2,0,34,12,13,1,4,98,67,45}
	//arr := []int{2,0,34,12,13}
	arr := GetRandomArrayWG()
	var wg sync.WaitGroup
	var sorterArr []int
	wg.Add(1)
	t1 := time.Now()
	go func() {
		sorterArr = mergeSortConWG(arr)
		//fmt.Println(sorterArr)
		wg.Done()
	}()
	wg.Wait()
	t2 := time.Since(t1)

	fmt.Println("Time Taken :", t2.Milliseconds())
	//fmt.Println("Sorted", sorterArr)

}

// Generate an array of 10 million random integers.
func GetRandomArrayWG() []int{
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	arr := make([]int, 10000000)
	for i:=0;i<10000000;i++ {
		arr[i] = r.Intn(1000000000)
	}
	return arr
}
