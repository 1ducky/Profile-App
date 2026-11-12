package localstorage

import (
	"os"
)

func StorageWorkerMKDIR(job ...string) {
	Jobs := make(chan string)
	for _, j := range job {
		Jobs <- j
	}
	close(Jobs)

	for j := range Jobs {
		err := os.MkdirAll(j, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
