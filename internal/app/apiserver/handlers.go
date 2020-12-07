package apiserver

import (
	"io"

	"github.com/gin-gonic/gin"
)

func TestIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		io.WriteString(ctx.Writer, "Testing is ok")
	}
}
