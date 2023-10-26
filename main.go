package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL as the first argument.")
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide the roll seed as the second argument.")
		return
	}

	url := os.Args[1]
	seed, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		panic(err)
	}

	pattern := `https:\/\/[a-zA-Z0-9-]+\.slack\.com\/archives\/[a-zA-Z0-9]+\/p(\d+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(url)

	if len(match) < 1 {
		fmt.Println("Invalid / incorrect message URL provided.")
		return
	}

	unixMsStr := match[1]

	unixMs, err := strconv.ParseInt(unixMsStr, 10, 64)
	if err != nil {
		panic(err)
	}

	unixNs := unixMs * 1000000
	timestampToSeed := unixNs % (2<<31 - 1)
	offset := seed - timestampToSeed

	fmt.Println("Expected seed: ", seed, ". Actual seed: ", timestampToSeed, ". Offset: ", offset)
}
