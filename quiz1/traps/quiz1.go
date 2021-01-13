package quiz

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var results = make([]int, 0, 1000)
var total_quiz_count int

func ReadQuiz(f *os.File) [][]string {
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func WriteQuiz(records [][]string, rf int, stop chan time.Duration, limit int) {
	total_quiz_count = len(records)

	for i, _ := range records {
		record := records[i][rf]
		fmt.Printf("Problem #%d: %s = ", i+1, record)
		var sum int
		_, scan_err := fmt.Scanln(&sum)
		//fmt.Printf("%v, %T\n", sum, sum)
		if scan_err != nil {
			log.Fatal(scan_err)

		}

		real_sum, err_s2i := strconv.Atoi(strings.TrimSpace(records[i][rf+1]))
		if err_s2i != nil {
			log.Fatal(err_s2i)
		}

		updateResult(i, sum, real_sum)
	}
  stop <- time.Duration(0)
}

func updateResult(index, sum, real_sum int) {
	result := 0
	if sum == real_sum {
		result = 1
	}

	results = append(results, result)
}

func PrintResult() {
	total_correct := len(results)
	fmt.Println("\nResult = ", total_correct, "/", total_quiz_count)
}
