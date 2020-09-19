package endpoints

import (
	"encoding/json"
	database "github.com/TicketsBot/database/translations"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Translation struct {
	Id         database.MessageId `json:"id"`
	SimpleName string             `json:"simple_name"`
	Value      string             `json:"value"`
}

func JsonHandler(ctx *gin.Context) {
	var translations []Translation

	for simpleName, id := range database.Messages {
		translations = append(translations, Translation{
			Id:         id,
			SimpleName: simpleName,
			Value:      ctx.PostForm(strconv.Itoa(int(id))),
		})
	}

	jsonRaw, err := json.MarshalIndent(translations, "", "    ")
	if err != nil {
		panic(err)
	}

	json := string(jsonRaw)
	json = strings.Replace(json, "\\u003c", "<", -1)
	json = strings.Replace(json, "\\u003e", ">", -1)
	json = strings.Replace(json, "\\u0026", "&", -1)

	ctx.Header("Content-Disposition", "attachment; filename=translations.json")
	ctx.Data(200, "application/json", []byte(json))
}
