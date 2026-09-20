package main
import ("bufio"; "errors"; "fmt"; "os"; "strconv")

func validateAge(s string) (int, error) {
    // implement
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("parse: %v", err)
		} else if n < 0 {
			return 0, errors.New("negative")
		} else {
    return n, nil
	  }
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    age, err := validateAge(sc.Text())
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
    } else {
        fmt.Printf("age: %d\n", age)
    }
}
