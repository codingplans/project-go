package main

func main() {
	for counter, n := 0, 2; n >= 0; n-- {
		defer func() {
			print(counter)
			counter++
		}()
	}
}
