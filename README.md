# Playing with gpt-3.5 turbo API using Go

Trying out the new gpt-3.5 turbo api by getting the model to answers a set of question based on a data from .csv file.

`dummy-sheet.csv` contains a list of data regarding a fictional college along with its student details.

## CSV File Credits

The data used in this project is from [this Example Spreadsheet](https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit#gid=0)

## Usage

You need to have a [OpenAI API Key](https://beta.openai.com/account/api-keys) to run this project.

Change the `.env.example` file to `.env` and add your API key.

### Install dependencies

```bash
go mod tidy
```

### Run

```bash
go run main.go
```
