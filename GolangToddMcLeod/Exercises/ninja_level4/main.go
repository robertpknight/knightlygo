package main

import (
	"github.com/KnightlyGo/Golang_Todd_McLeod/ninja_level4/exercises"
)

func main() {
	exercises.Exercise1()

	// declare slice to be used in exercises
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	exercises.Exercise2(a)
	exercises.Exercise3(a)
	exercises.Exercise4(a)
	exercises.Exercise5(a)

	exercises.Exercise6()
	exercises.Exercise7()
	agents := exercises.Exercise8()
	agentsNewRecord := exercises.Exercise9(agents)
	exercises.Exercise10(agentsNewRecord)
}
