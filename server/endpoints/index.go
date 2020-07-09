package endpoints

import (
	"github.com/TicketsBot/database/translations"
	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.tmpl", gin.H{
		"messages": database.Messages,
	})
}
