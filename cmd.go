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
	CapCount  int
	Password  string
}

func NewState() State {
	return State{CharCount: 12, NumCount: 0, CapCount: 0, Password: ""}
}

func randomS() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	str := charset[random.Intn(len(charset))]
	return string(str)
}

func main() {
	flag.Parse()
	flags := flag.Args()
	flagCount := len(flags)
	if flagCount == 0 {
		fmt.Println("Empty")
	}

	if flagCount >= 1 && flags[0] == "gen" {
		var str strings.Builder
		state := NewState()
		for i := 0; i < state.CharCount; i++ {
			str.WriteString(randomS())
		}

		fmt.Println("Your generated password is\n------------------------\n " + str.String())
	}
}
