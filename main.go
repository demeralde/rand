package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	argsLen := len(args)

	min, minErr := strconv.Atoi(os.Args[1])
	max, maxErr := strconv.Atoi(os.Args[2])

	if argsLen < 3 {
		log.Fatal("You must provide a `min` and `max` number, e.g `rand 1 3`")
	} else if argsLen > 3 {
		log.Fatal("You must only provide 2 arguments (the min and max number)")
	}

	if minErr != nil {
		log.Fatal("The `min` arg must be an integer")
	}

	if maxErr != nil {
		log.Fatal("The `max` arg must be an integer")
	}

	if min == max {
		log.Fatal("`min` and `max` must be different numbers")
	}

	if min > max {
		log.Fatal("`max` must be greater than `min`")
	}

	number := rand.Intn((max+1)-min) + min
	fmt.Println(number)
}
