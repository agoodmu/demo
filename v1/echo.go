package v1

import (
	"embed"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Echo(tpl_fs *embed.FS, tpl_name string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		tmpl, err := template.ParseFS(tpl_fs, tpl_name)
		if err != nil {
			slog.Error("failed to parse template for echo page", "Error", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		tmpl.Execute(c.Writer, c.Request.Header)
	}

	return gin.HandlerFunc(fn)
}
