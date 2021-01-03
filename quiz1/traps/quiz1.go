package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadQuiz(f *os.File) (*csv.Reader) {
	r := csv.NewReader(f)
	return r
}

func WriteQuiz(r *csv.Reader, rf int) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[rf])
	}

}
