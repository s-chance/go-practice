package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(maxNum)
	fmt.Println("The secret is", secret)
	fmt.Println("Please input your answer")
	var ans int
	for {
		_, err := fmt.Scanf("%d", &ans)
		if err != nil {
			fmt.Println("An invalid input. Please try again,", err)
			continue
		}

		if ans > secret {
			fmt.Println("Your answer is bigger than the secret. Please try again")
		} else if ans < secret {
			fmt.Println("Your answer is smaller than the secret. Please try again")
		} else {
			fmt.Println("Correct!")
			break
		}
	}
}
