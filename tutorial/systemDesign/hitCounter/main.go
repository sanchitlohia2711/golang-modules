package main

import "fmt"

var (
	counter   [5]int
	timestamp [5]int
)

func recordHit(timestamp_epoch int) {

	i := timestamp_epoch % 5

	if timestamp[i] != timestamp_epoch {
		counter[i] = 1
		timestamp[i] = timestamp_epoch
	} else {
		counter[i]++
	}

	fmt.Printf("After hit at time: %d\n", timestamp_epoch)
	fmt.Println(counter)
	fmt.Println(timestamp)
	fmt.Println()

}

func getNumberOfHit(timestamp_epoch int, minute int) (int, error) {

	if minute > 5 {
		return 0, fmt.Errorf("incorrect value of past minute passed")
	}
	hits := 0

	for k := timestamp_epoch; k > timestamp_epoch-minute; k-- {
		i := k % 5
		if timestamp[i] > timestamp_epoch-5 {
			hits = hits + counter[i]
		}
	}

	return hits, nil

}
func main() {
	recordHit(1000)

	recordHit(1000)

	hits, _ := getNumberOfHit(1000, 1)
	fmt.Printf("Current timestamp: %d,Last minue: %d,  Hits:%d\n\n", 1000, 1, hits)

	recordHit(1001)

	hits, _ = getNumberOfHit(1001, 2)
	fmt.Printf("Current timestamp: %d, Last minue: %d, Hits:%d\n\n", 1001, 2, hits)

	recordHit(1005)

	hits, _ = getNumberOfHit(1005, 5)
	fmt.Printf("Current timestamp: %d, Last minue: %d,  Hits:%d\n", 1005, 5, hits)

	hits, _ = getNumberOfHit(1005, 2)
	fmt.Printf("Current timestamp: %d, Last minue: %d, Hits:%d\n", 1005, 2, hits)

	recordHit(1007)
	hits, _ = getNumberOfHit(1007, 5)
	fmt.Printf("Current timestamp: %d, Last minue: %d, Hits:%d\n", 1007, 5, hits)

}
