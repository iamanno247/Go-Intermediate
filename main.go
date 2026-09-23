package main
import ("bufio"; "fmt"; "os"; "strconv"; "strings"; "sync")
func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan(); n, _ := strconv.Atoi(sc.Text())
    sc.Scan(); fields := strings.Fields(sc.Text())
    nums := make([]int, n)
    for i, f := range fields { nums[i], _ = strconv.Atoi(f) }
    // split into 4 chunks, goroutine each, sum total
    var wg sync.WaitGroup
    var total int
    var mu sync.Mutex
		chunk := (n + 3) / 4
		for i := 0; i < 4; i++ {
			start, end := i * chunk, (i+1) * chunk
			if start > n { start = n}
			if end > n { end = n}
			wg.Add(1)
			go func(subSlice []int) {
				defer wg.Done()
				localsum := 0
				for _, value := range subSlice {
					localsum += value
				}
				mu.Lock()
				total += localsum
				mu.Unlock()
			}(nums[start:end])
		}
		wg.Wait()
    fmt.Println(total)  // replace with the real total
}
