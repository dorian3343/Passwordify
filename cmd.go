package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type State struct {
	CharCount int
	NumCount  int
	Password  string
}

func NewState() State {
	return State{CharCount: 12, NumCount: 0, Password: ""}
}

func randomS() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	str := charset[random.Intn(len(charset))]
	return string(str)
}
func main() {
	var str strings.Builder
	state := NewState()

	// CLI flags
	charCountPtr := flag.Int("char", 12, "number of characters in the generated password")
	flag.Parse()

	// Check if the "char" flag is provided
	if *charCountPtr < 1 {
		fmt.Println("Invalid value for char flag. Defaulting to 12.")
		*charCountPtr = 12 // Defaulting to 12 if provided value is less than 1
	}

	flags := flag.Args()
	flagCount := len(flags)
	if flagCount == 0 {
		fmt.Println("Missing flags")
	}

	if flagCount >= 1 && flags[0] == "gen" {
		state.CharCount = *charCountPtr
		fmt.Println(state.CharCount)
		for i := 0; i < state.CharCount; i++ {
			str.WriteString(randomS())
		}
		fmt.Println("Your generated password is:\n------------------------\n" + str.String())
	}
}
