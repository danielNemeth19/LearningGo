package main

import "fmt"

const NUMBER_OF_PARTITION = 4

func swap(arr []int, i1 int, i2 int) {
	if arr[i1] > arr[i2] {
		temp := arr[i1]
		arr[i1] = arr[i2]
		arr[i2] = temp
	}
}

func sortOne(arr []int, last int) {
	for i := 0; i < last-1; i++ {
		nextIndex := i + 1
		swap(arr[:], i, nextIndex)
	}
}

func sortAll(arr []int, length int) {
    fmt.Printf("Sorting here: %v\n", arr)
	for j := length; j > 1; j-- {
		sortOne(arr[:], j)
	}
}

func sortByChannel(arr []int, last int, c chan []int) {
	sortAll(arr, last)
	c <- arr
}

func mergePartition(p1 []int, p2 []int, p3 []int, p4 []int) []int {
	length := len(p1) + len(p1) + len(p3) + len(p4)
	arrnew := make([]int, length)

	index := 0
	for i := 0; i < len(p1); i++ {
		arrnew[index] = p1[i]
		index++
	}

	for i := 0; i < len(p2); i++ {
		arrnew[index] = p2[i]
		index++
	}

	for i := 0; i < len(p3); i++ {
		arrnew[index] = p3[i]
		index++
	}

	for i := 0; i < len(p4); i++ {
		arrnew[index] = p4[i]
		index++
	}

	return arrnew
}

func main() {

	var length int

	fmt.Print("Please enter length of you array: ")
	fmt.Scan(&length)

	arrInt := make([]int, length)
	arrSorted := make([]int, length)
	fmt.Print("Please input a series of ", length, " integers. Each integer can be separated by space:")

	for i := 0; i < length; i++ {
		fmt.Scan(&arrInt[i])
	}

	if length <= 4 {
		sortAll(arrInt[:], length)
	} else {
		c := make(chan []int, NUMBER_OF_PARTITION)
		partitionSize := length / NUMBER_OF_PARTITION
		m := length % NUMBER_OF_PARTITION
		if m > 0 {
			partitionSize++
		}

		var last int
		index := 0

		if m == 0 {
			last = partitionSize
		} else {
			last = partitionSize - 1
		}
		partition1 := make([]int, last)
		for j := 0; j < last; j++ {
			partition1[j] = arrInt[index]
			index++
		}
		fmt.Println("partition1: ", partition1)

		if (m == 0) || (m == 3) {
			last = partitionSize
		} else {
			last = partitionSize - 1
		}
		partition2 := make([]int, last)
		for j := 0; j < last; j++ {
			partition2[j] = arrInt[index]
			index++
		}
		fmt.Println("partition2: ", partition2)

		if (m == 0) || (m >= 2) {
			last = partitionSize
		} else {
			last = partitionSize - 1
		}
		partition3 := make([]int, last)
		for j := 0; j < last; j++ {
			partition3[j] = arrInt[index]
			index++
		}
		fmt.Println("partition3: ", partition3)

		if (m == 0) || (m >= 1) {
			last = partitionSize
		} else {
			last = partitionSize - 1
		}
		partition4 := make([]int, last)
		for j := 0; j < last; j++ {
			partition4[j] = arrInt[index]
			index++
		}
		fmt.Println("partition4: ", partition4)

		go sortByChannel(partition1, len(partition1), c)
		go sortByChannel(partition2, len(partition2), c)
		go sortByChannel(partition3, len(partition3), c)
		go sortByChannel(partition4, len(partition4), c)

		sortedPartition1 := <-c
		sortedPartition2 := <-c
		sortedPartition3 := <-c
		sortedPartition4 := <-c

		fmt.Println("Sorted partition1: ", sortedPartition1)
		fmt.Println("Sorted partition2: ", sortedPartition2)
		fmt.Println("Sorted partition3: ", sortedPartition3)
		fmt.Println("Sorted partition4: ", sortedPartition4)

		arrSorted = mergePartition(sortedPartition1, sortedPartition2, sortedPartition3, sortedPartition4)
		sortAll(arrSorted[:], length)
	}

	fmt.Println("Sorted Array: ", arrSorted)
}
