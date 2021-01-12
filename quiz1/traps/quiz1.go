package quiz

import (
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadQuiz(f *os.File) *csv.Reader {
	r := csv.NewReader(f)
	return r
}

func WriteQuiz(r *csv.Reader, rf int, stop chan time.Duration, limit int) {
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	results := make ([]int, len(records))

	for i, _ := range records {
    time_left := <-stop; 
		//fmt.Println(time_left, time.Duration(0))
		if time_left > time.Duration(0) {
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

			updateResult(results, i, sum, real_sum)
		}
	}
	printResult(results, records)
}

func updateResult(results []int, index, sum, real_sum int) {
	result := 0
	if sum == real_sum {
		result = 1
	}

	results[index] = result
}

func printUniqueValue(arr []int, numv int) int {
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict[numv]
}

func printResult(results []int, records [][]string) {
	total_quiz_count := len(records)
	total_correct := printUniqueValue(results, 1)
	fmt.Println("\nResult = ", total_correct, "/", total_quiz_count)
}

