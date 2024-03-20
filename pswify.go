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
func randomI() int {
	min := 0
	max := 9
	return rand.Intn((max - min) + min)
}
func randomS() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	str := charset[random.Intn(len(charset))]
	return string(str)
}
func randomC() string {
	const charset = "!@#$%^&*()_+{}|?><.,/;"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	str := charset[random.Intn(len(charset))]
	return string(str)
}

func main() {
	var str strings.Builder
	state := NewState()

	// CLI flags
	numCountPtr := flag.Int("num", 0, "number of numbers in the generated password")
	symbolCountPtr := flag.Int("symboli", 12, "number of special characters in the generated password")
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
	}

	if flagCount >= 1 && flags[0] == "gen" {
		state.NumCount = *numCountPtr
		state.CharCount = *charCountPtr
		state.SymbolCount = *symbolCountPtr
		for i := 0; i < state.CharCount; i++ {
			str.WriteString(randomS())
		}
		for i := 0; i < state.NumCount; i++ {
			num := randomI()
			str.WriteString(strconv.Itoa(num))
		}
		for i := 0; i < state.SymbolCount; i++ {
			str.WriteString(randomC())
		}
		fmt.Println("Your generated password is:\n------------------------\n" + str.String())
	}
}
