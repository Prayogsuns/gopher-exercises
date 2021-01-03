package main

import (
	"fmt"
	"github.com/prayogsuns/gopher-exercises/quiz1/traps"
	"log"
	"os"
)

func main() {
	filename := "quiz.csv"
	if os.Args[1] != nil {
        	filename = os.Args[1]
	}
	f, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	r := quiz.ReadQuiz(f)

	quiz.WriteQuiz(r, 0)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

}
