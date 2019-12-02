package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fuelSum, fuelForFuelSum, err := calcFuelSum()
	if err != nil {
		panic(nil)
	}

	fmt.Printf("Fuel needed (1): %d\n", fuelSum)
	fmt.Printf("Fuel needed (2): %d\n", fuelForFuelSum)
}

func calcFuelSum() (int, int, error) {
	fuelSum, fuelForFuelSum := 0, 0

	ch, err := ingestInputs()
	if err != nil {
		return 0, 0, nil
	}

	for i := range ch {
		fuelSum += calculateFuel(i)
		fuelForFuelSum += calculateFuelForFuel(i)
	}

	return fuelSum, fuelForFuelSum, nil
}

func ingestInputs() (<-chan int, error) {
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	ch := make(chan int)
	scanner := bufio.NewScanner(file)

	go func() {
		defer close(ch)
		defer file.Close()

		var i int

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Sscanf(line, "%d", &i)
			ch <- i
		}
	}()

	return ch, nil
}

func calculateFuel(mass int) int {
	// fmt.Println(mass)
	return (mass / 3) - 2
}

func calculateFuelForFuel(mass int) int {
	fuel := calculateFuel(mass)
	additionalFuel := calculateFuel(fuel)
	sum := fuel

	for additionalFuel > 0 {
		sum += additionalFuel
		additionalFuel = calculateFuel(additionalFuel)
	}

	return sum
}
