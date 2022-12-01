package randline

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Input struct {
	InputFilePath  string
	OutputFilePath string
	Times          *int
}

func Run(input Input) {
	log.Info("Starting randLine")
	log.Info("InputFile=", input.InputFilePath)
	log.Info("OutputFile=", input.OutputFilePath)
	log.Info("Times=", input.Times)

	input_file, err := os.Open(input.InputFilePath)
	defer input_file.Close()
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(input_file)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// write result

	output_file, err := os.Create(input.OutputFilePath)

	if err != nil {
		log.Panic(err)
	}

	defer output_file.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *input.Times; i++ {
		randomIndex := rand.Intn(len(text))
		_, err2 := output_file.WriteString(fmt.Sprintf("%v\n", text[randomIndex]))

		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Println("OK")
}
