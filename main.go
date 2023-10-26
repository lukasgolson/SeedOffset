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
		fmt.Println("Please provide the roll rollSeed as the second argument.")
		return
	}

	url := os.Args[1]
	rollSeed, err := strconv.ParseInt(os.Args[2], 10, 64)
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

	TSUnixMs, err := strconv.ParseInt(unixMsStr, 10, 64)
	if err != nil {
		panic(err)
	}

	TSUnixNs := TSUnixMs * 1000000
	expectedSeed := TSUnixNs % (2<<31 - 1)
	offset := rollSeed - expectedSeed

	fmt.Println("Expected rollSeed: ", expectedSeed, ". Actual rollSeed: ", rollSeed, ". Offset: ", offset)
}
