package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read(filename string) (chan int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	c := make(chan int)
	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			var v int
			fmt.Sscanf(scanner.Text(), "%d", &v)
			c <- v
		}

		close(c)
	}()

	return c, nil
}

func main() {
	sum := 0
	fcounts := make(map[int]int)

	fcounts[0] = 1

	for i := 1; ; i++ {
		c, err := read("input.txt")
		if err != nil {
			log.Fatal(err)
		}

		for v := range c {
			sum += v
			fcounts[sum] += 1

			if fcounts[sum] == 2 {
				fmt.Printf("fcounts[%d] == %d\n", sum, fcounts[sum])
				return
			}
		}

		fmt.Printf("End of loop %d: %d\n", i, sum)
	}
}
