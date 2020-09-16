package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/TicketsBot/database"
	db_translations "github.com/TicketsBot/database/translations"
	"github.com/TicketsBot/translationgenerator/server/endpoints"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
)

var (
	connString = flag.String("c", "", "postgres connection string")
	language   = flag.String("l", "", "language")
)

func main() {
	flag.Parse()

	pool, err := pgxpool.Connect(context.Background(), *connString)
	if err != nil {
		panic(err)
	}

	client := database.NewDatabase(pool)
	translations, err := client.Translations.GetAll()
	if err != nil {
		panic(err)
	}

	if *language == "" { // all
		for language, messages := range translations {
			writeFile(language, messages)
		}
	} else {
		// filter by language
		lang := db_translations.Language(*language)
		if _, ok := db_translations.Flags[lang]; !ok {
			panic(fmt.Errorf("invalid language: %s", lang))
		}

		writeFile(lang, translations[lang])
	}
}

func writeFile(language db_translations.Language, messages map[db_translations.MessageId]string) {
	fileName := fmt.Sprintf("%s.json", language)

	f, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	var formatted []endpoints.Translation
	for id, message := range messages {
		formatted = append(formatted, endpoints.Translation{
			Id:         id,
			SimpleName: *db_translations.GetSimpleName(id),
			Value:      message,
		})
	}

	jsonRaw, err := json.MarshalIndent(formatted, "", "    ")
	if err != nil {
		panic(err)
	}

	json := string(jsonRaw)
	json = strings.Replace(json, "\\u003c", "<", -1)
	json = strings.Replace(json, "\\u003e", ">", -1)
	json = strings.Replace(json, "\\u0026", "&", -1)

	if _, err := f.WriteString(json); err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %s\n", fileName)
}
