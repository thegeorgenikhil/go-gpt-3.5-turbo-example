package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/otiai10/openaigo"
)

const (
	preamble1 = `The following is the data from a csv file from a school which contains info about their name,gender,class level,home state, major and extracurricular activities. Give me a detailed analysis of the data which contains the following information:
	1. The number of students in each class level
	2. Most popular major
	3. Which extracurricular activity has the highest female representation
	4. Which state has the highest number of students

	The results need to be highly accurate and detailed. The data is given below:

	`

	preamble = `Make me understand what the given table is about and talk me through the data points in the table.The results need to be highly accurate and comprehensive. Also talk about the following points:
	1. The number of students in each class level
	2. Most popular major
	3. Highest female representation in extracurricular activities
	
	The data is given below:

	`
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENAI_API_KEY environment variable")
		return
	}
	client := openaigo.NewClient(apiKey)
	ctx := context.Background()

	// open and read from csv file in go
	fd, error := os.Open("dummy-sheet.csv")
	defer fd.Close()

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Successfully opened the CSV file")

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	// crete a string from the csv data
	var csvData string
	for index, row := range records {
		if index == 0 {
			continue
		}
		csvData += strings.Join(row, ",") + "\n\n"
	}

	// create the prompt
	prompt := preamble + csvData

	request := openaigo.ChatCompletionRequestBody{
		Temperature: 0,
		Model:       "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 1000,
	}
	response, err := client.Chat(ctx, request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response: ", response.Choices[0].Message.Content)
}
