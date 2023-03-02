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
	preambleForData = `
	You are given a set of csv data, keep this in mind and answer the subsequent questions from the user.
	The data is as follows: 
	`
	prompt1 = `You are a data analyst that anaylyzes complex data sets and provides comprehensive reports.`
	
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
	prompt2 := preambleForData + csvData

	request := openaigo.ChatCompletionRequestBody{
		Temperature: 0,
		Model:       "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "system", Content: prompt1},
			{Role: "system", Content: prompt2},
			{Role: "user", Content: "What is the following data set about?"},
			{Role: "user", Content: "What is most popular major/majors?"},
			{Role: "user",Content: "Where are most of the students from?"},
			{Role: "user", Content: "Which extra-curricular activity/activities is most popular? Which among them has the highest female participation?"},
			// add in as many questions as you want
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
