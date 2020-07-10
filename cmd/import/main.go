package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/TicketsBot/database"
	"github.com/TicketsBot/translationgenerator/server/endpoints"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
)

var (
	connString = flag.String("c", "", "postgres connection string")
	file       = flag.String("f", "", "translation json file path")
	language   = flag.String("l", "en", "language")
)

func main() {
	flag.Parse()

	pool, err := pgxpool.Connect(context.Background(), *connString)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	var translations []endpoints.Translation
	if err := json.Unmarshal(data, &translations); err != nil {
		panic(err)
	}

	database := database.NewDatabase(pool)
	database.CreateTables(pool)

	columns := []string{"language", "message_id", "content"}

	var rows [][]interface{}
	for _, translation := range translations {
		rows = append(rows, []interface{}{
			*language,
			translation.Id,
			translation.Value,
		})
	}

	if _, err := database.Translations.CopyFrom(context.Background(), pgx.Identifier{"translations"}, columns, pgx.CopyFromRows(rows)); err != nil {
		panic(err)
	}
}
