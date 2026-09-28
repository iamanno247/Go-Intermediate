package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sync"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Buffer(make([]byte, 1024*1024), 1024*1024)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        sc.Scan()
        nums[i], _ = strconv.Atoi(sc.Text())
    }

    var mu sync.Mutex
    var wg sync.WaitGroup
    total := 0

    // TODO: split `nums` into 4 chunks and launch a goroutine per chunk.
    // Each goroutine should sum its chunk and add the result into `total`,
    // guarded by `mu`. Use `wg` (Add/Done/Wait) to wait for all goroutines
    // to finish before printing.
    chunks := (n + 3) / 4
		for i := 0; i < 4; i++ {
			start, end := i * chunks, (i + 1) * chunks
			if start > n {
				start = n
			}
			if end > n {
				end = n
			}
			numchunk := nums[start:end]
			wg.Add(1)
			go func(s []int) {
				defer wg.Done()
				localtotal := 0
				for _, x := range s {
				localtotal += x
				}
				mu.Lock()
				total += localtotal
				mu.Unlock()
			}(numchunk)
		}
		wg.Wait()

    fmt.Println(total)
}
