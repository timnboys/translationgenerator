package endpoints

import (
	database "github.com/TicketsBot/database/translations"
	"github.com/gin-gonic/gin"
	"strconv"
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

	ctx.Header("Content-Disposition", "attachment; filename=translations.json")
	ctx.JSON(200, translations)
}
