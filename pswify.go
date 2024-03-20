package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type State struct {
	CharCount   int
	NumCount    int
	SymbolCount int
}

func NewState() State {
	return State{CharCount: 12, NumCount: 0, SymbolCount: 0}
}

func randomChar(charset string) string {
	return string(charset[rand.Intn(len(charset))])
}

func main() {
	var str strings.Builder
	state := NewState()

	// Seed initialization
	rand.Seed(time.Now().UnixNano())

	// CLI flags
	numCountPtr := flag.Int("num", 0, "number of numbers in the generated password")
	symbolCountPtr := flag.Int("symbol", 0, "number of special characters in the generated password")
	charCountPtr := flag.Int("char", 12, "number of characters in the generated password")
	flag.Parse()

	if *charCountPtr < 1 {
		fmt.Println("Invalid value for char flag. Defaulting to 12.")
		*charCountPtr = 12
	}

	flags := flag.Args()
	flagCount := len(flags)
	if flagCount == 0 {
		fmt.Println("Missing flags")
		return
	}

	if flagCount >= 1 && flags[0] == "gen" {
		state.NumCount = *numCountPtr
		state.CharCount = *charCountPtr
		state.SymbolCount = *symbolCountPtr

		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		const symbolCharset = "!@#$%^&*()_+{}|?><.,/;"

		for i := 0; i < state.CharCount; i++ {
			str.WriteString(randomChar(charset))
		}
		for i := 0; i < state.NumCount; i++ {
			num := rand.Intn(10) // Random number between 0-9
			str.WriteString(strconv.Itoa(num))
		}
		for i := 0; i < state.SymbolCount; i++ {
			str.WriteString(randomChar(symbolCharset))
		}
		fmt.Println("Your generated password is:\n------------------------\n" + str.String())
	}
}
