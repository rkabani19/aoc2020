package main

import (
	"fmt"

	"github.com/rkabani19/aoc2020/utils"
)

const (
	subjectNumber = 7
	divisor       = 20201227
)

func main() {
	input, _ := utils.ReadFile("day25.txt")

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	cardPublicKey := utils.ToInt(input[0])
	doorPublicKey := utils.ToInt(input[1])
	cardLoopSize := getLoopSize(cardPublicKey)
	doorLoopSize := getLoopSize(doorPublicKey)
	doorEncryption := encrypt(doorPublicKey, cardLoopSize)
	cardEncryption := encrypt(cardPublicKey, doorLoopSize)

	if cardEncryption != doorEncryption {
		panic(fmt.Errorf("Error. Encryption codes do not match.\n"))
	}

	return doorEncryption
}

func encrypt(key int, loopSize int) int {
	out := 1
	for i := 0; i < loopSize; i++ {
		out = (out * key) % divisor
	}
	return out
}

func getLoopSize(key int) int {
	i, product := 0, 1
	for key != product {
		product = (product * subjectNumber) % divisor
		i++
	}
	return i
}
