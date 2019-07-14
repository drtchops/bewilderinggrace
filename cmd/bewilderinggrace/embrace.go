package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/drtchops/bewilderinggrace/src/bewilderinggrace"
)

func main() {
	fmt.Println("This will overwrite one of your saves with a randomized one. Embrace the grace!")

	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a save slot (default 1): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var saveSlot int64
	text = strings.TrimSpace(text)
	if text != "" {
		saveSlot, err = strconv.ParseInt(strings.TrimSpace(text), 10, 32)
		if err != nil {
			panic(err)
		}
	} else {
		saveSlot = 1
	}

	fmt.Print("Enter a seed number (default random): ")
	text, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var seed int64
	text = strings.TrimSpace(text)
	if text != "" {
		seed, err = strconv.ParseInt(strings.TrimSpace(text), 10, 32)
		if err != nil {
			panic(err)
		}
	} else {
		rand.Seed(time.Now().Unix())
		seed = rand.Int63()
	}

	fmt.Printf("Randomizing slot %d with seed %d\n", saveSlot, seed)

	err = bewilderinggrace.Randomize(int(saveSlot), seed)
	if err != nil {
		panic(err)
	}

	fmt.Println("gl;hf")

	fmt.Println("Press enter to quit.")
	reader.ReadString('\n')
}
