package main2

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main2() {

	fmt.Println("Welcome to deepu world")

	rand.Seed(int64(time.Now().Second()))

	guessNum := rand.Intn(100)

	fmt.Println(guessNum)

	fmt.Print("Guess the number : ")

	for index := 0; index <= 10; index++ {

		reader := bufio.NewReader(os.Stdin)
		inputString, _ := reader.ReadString('\n')

		inputString = strings.TrimSpace(inputString)

		enteredValue, error := strconv.Atoi(inputString)

		if error != nil {
			fmt.Println("Incorrect input. Please enter the value between 0 to 100")
			index = index - 1
		}

		if guessNum == enteredValue {
			fmt.Println("Guess value is correct")
			break
		} else if guessNum > enteredValue {
			fmt.Println("Guess value is low")
		} else {
			fmt.Println("Guess value is high")
		}

		if index == 10 {
			fmt.Println("Ooops!!!!!!.Your lost the game. The guess value was : ", guessNum)

		} else {
			fmt.Println("You guess was wrong. ", 9-index, "try remaining")
		}

	}

	fmt.Println("Your Guess game is done!!!!!!!!!!")

}
