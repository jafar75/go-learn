package main

import (
	"fmt"
	"sync"
)

func combineChannels(leftChan, rightChan chan int, resultChan chan int) {
    defer close(resultChan)

    val1, ok1 := <-leftChan
    val2, ok2 := <-rightChan
    for ok1 && ok2 {
        if val1 < val2 {
            resultChan <- val1
            val1, ok1 = <-leftChan
        } else {
            resultChan <- val2
            val2, ok2 = <-rightChan
        }
    }
    for ok1 {
        resultChan <- val1
        val1, ok1 = <-leftChan
    }
    for ok2 {
        resultChan <- val2
        val2, ok2 = <-rightChan
    }
}

func sortConcurrently(data []int, outputChan chan int) {
    if len(data) < 2 {
        outputChan <- data[0]
        defer close(outputChan)
        return
    }

    leftChan := make(chan int)
    rightChan := make(chan int)

    go sortConcurrently(data[len(data)/2:], leftChan)
    go sortConcurrently(data[:len(data)/2], rightChan)
    go combineChannels(leftChan, rightChan, outputChan)
}

func main() {
    inputData := []int{5, 4, 3, 2, 1}
    fmt.Println("unsorted array:", inputData)
    outputChan := make(chan int)

    var waitGroup sync.WaitGroup
    waitGroup.Add(1)

    go func() {
        defer waitGroup.Done()
        sortConcurrently(inputData, outputChan)
    }()

    waitGroup.Wait()

    var sortedData []int
    for value := range outputChan {
        sortedData = append(sortedData, value)
    }

    fmt.Println("after sort:", sortedData) // Output: [1 2 3 4 5]
}