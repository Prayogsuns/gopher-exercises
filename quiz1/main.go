package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/prayogsuns/gopher-exercises/quiz1/traps"
	"log"
	"os"
	"time"
)

var flagSetName string = os.Args[0]
var stop = make(chan time.Duration)

func usage() (string, int) {
	fs := flag.NewFlagSet(flagSetName, flag.ExitOnError)
	fcsv := fs.String("csv", "quiz.csv", "a csv file in the format of 'question.answer'")
	flimit := fs.Int("limit", 30, "the time limit of quiz in seconds")

	//ferror := fs.Parse(os.Args[1:])
	fs.Parse(os.Args[1:])
	//fmt.Println("XYZ")
	//fmt.Println(fmt.Print(ferror))
	//fmt.Println("AAA")
	return *fcsv, *flimit
}

func main() {
	//if len(os.Args) == 2 && os.Args[1] == "-h" {
	csv, limit := usage()
	//fmt.Println(csv, limit)
	if (csv == "-limit") || (csv == "--limit") {
		fmt.Println("flag csv can't be set to -limit|--limit")
		os.Exit(1)
	}
	//fmt.Println("BBB")

	filename := csv
	f, err := os.OpenFile(filename, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	r := quiz.ReadQuiz(f)

	startQuiz()

	startTime := time.Now()
	go watchTime(stop, startTime, limit)
	//quiz.WriteQuiz(r, 0, stop, limit)
	go quiz.WriteQuiz(r, 0, limit)

	if time_left := <-stop; time_left == time.Duration(0) {
		quiz.PrintResult()
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

}

func startQuiz() {
	fmt.Printf("Start Quiz, [y|Y]: ")
	var start []byte = make([]byte, 1)
	_, start_err := fmt.Scanln(&start)
	if start_err != nil {
		log.Fatal(start_err)
	}

	bad_start := false
	if len(start) != 1 {
		bad_start = true
	} else {
		if !bytes.Equal(bytes.ToLower(start), []byte("y")) {
			bad_start = true
		}
	}

	if bad_start {
		fmt.Println("Wrong Input! Expected y|Y. Exiting...")
		os.Exit(1)
	}

}

func watchTime(stop chan time.Duration, start time.Time, limit int) {
	for {
		t := time.Now()
		elapsed := t.Sub(start)
		//fmt.Println(elapsed, (time.Duration(limit) * time.Second))
		if elapsed >= (time.Duration(limit) * time.Second) {
			stop <- time.Duration(0)
		} // else {
		//  stop <- elapsed
		//}
	}
}
