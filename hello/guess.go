package hello

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Game() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	for x := 1; x <= 10; x++ {
		fmt.Println(11-x, "Guesses Left,", "Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Low")
		} else if guess > target {
			fmt.Println("High")
		} else {
			fmt.Println("Right")
			break
		}

	}

}
