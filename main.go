package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var matrixA = [500][500]int{}
var matrixB = [500][500]int{}
var matrixC = [500][500]int{}
var matrixD = [500][500]int{}
var matrixE = [500][500]int{}
var matrixF = [500][500]int{}

func fillmatA(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < len(matrixA); i++ {
		for j := 0; j < len(matrixA); j++ {
			matrixA[i][j] = rand.Intn(100)
		}
	}
}

func fillmatB(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < len(matrixB); i++ {
		for j := 0; j < len(matrixB); j++ {
			matrixB[i][j] = rand.Intn(100)
		}
	}
}

func fillmatD(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < len(matrixD); i++ {
		for j := 0; j < len(matrixD); j++ {
			matrixD[i][j] = rand.Intn(100)
		}
	}
}

func fillmatE(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < len(matrixE); i++ {
		for j := 0; j < len(matrixE); j++ {
			matrixE[i][j] = rand.Intn(100)
		}
	}
}

func matrixMult1() {
	for i := 0; i < len(matrixA); i++ {
		for j := 0; j < len(matrixB); j++ {
			for k := 0; k < len(matrixB); k++ {
				matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
}

func matrixMult2() {

	for i := 0; i < len(matrixD); i++ {
		for j := 0; j < len(matrixE); j++ {
			for k := 0; k < len(matrixE); k++ {
				matrixF[i][j] += matrixD[i][k] * matrixE[k][j]
			}
		}
	}
}

func concurrentMatrixMult() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < len(matrixA); i++ {
			for j := 0; j < len(matrixB); j++ {
				for k := 0; k < len(matrixB); k++ {
					matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
				}
			}
		}
		defer wg.Done()
	}()

	go func() {
		for i := 0; i < len(matrixD); i++ {
			for j := 0; j < len(matrixE); j++ {
				for k := 0; k < len(matrixE); k++ {
					matrixF[i][j] += matrixD[i][k] * matrixE[k][j]
				}
			}
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func displayMat(b [500][500]int) {
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			print(b[i][j])
		}
		println()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	start := time.Now()
	go fmt.Print("\nGenerating matrices.\n")
	go fillmatA(&wg, 1)
	go fillmatB(&wg, 2)
	go fillmatD(&wg, 3)
	go fillmatE(&wg, 4)
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Time taken to fill matrices %dns\n", duration.Nanoseconds())
	var input = 0
	for input > 3 || input < 1 {
		fmt.Println("Perform Matrix multiplication \n 1. Both \n 2. Sequentially\n 3. Concurrently")
		fmt.Scanln(&input)
	}
	runs := 0
	var totalTime int64 = 0
	var totalTimeseq int64 = 0
	var totalTimeCon int64 = 0
	switch input {
	case 1:
		fallthrough

	case 2:
		println("Sequential Simulation")
		for runs < 10 {
			start := time.Now()
			matrixMult1()
			matrixMult2()
			duration := time.Since(start)
			totalTime += duration.Nanoseconds()
			runs++
			fmt.Printf("Time taken to solve %d sequentially %dns\n", runs, duration.Nanoseconds())
		}
		totalTimeseq = totalTime / (int64(runs))
		fmt.Println()

		if input != 1 {
			break
		}
		fallthrough

	case 3:
		totalTime = 0
		runs = 0
		println("Concurrent Simulation")
		for runs < 10 {
			start := time.Now()
			concurrentMatrixMult()
			duration := time.Since(start)
			totalTime += duration.Nanoseconds()
			runs++
			fmt.Printf("Time taken to solve %d concurrently %dns\n", runs, duration.Nanoseconds())
		}
		totalTimeCon = totalTime / (int64(runs))
	}
	println()
	if input == 1 || input == 2 {
		fmt.Printf("Average time taken to run %d sequential simulations %d ns\n", runs, totalTimeseq)
	}
	if input == 1 || input == 3 {
		fmt.Printf("Average time taken to run %d concurrent simulations %d ns\n", runs, totalTimeCon)
	}
}
