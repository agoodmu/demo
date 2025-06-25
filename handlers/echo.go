package handlers

import (
	"embed"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type EchoServer struct{}

func (e *EchoServer) template() *template.Template {
	return nil
}

type echoData struct {
	Headers    http.Header
	ClientInfo map[string]any
}

func Echo(tpl_fs *embed.FS, tpl_name string) gin.HandlerFunc {
	data := echoData{Headers: make(http.Header), ClientInfo: make(map[string]any)}
	fn := func(c *gin.Context) {
		tmpl, err := template.ParseFS(tpl_fs, tpl_name)
		if err != nil {
			slog.Error("failed to parse template for echo page", "Error", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		data.Headers = c.Request.Header
		data.ClientInfo["Host"] = c.Request.Host
		data.ClientInfo["RemoteAddr"] = c.Request.RemoteAddr
		if proxies := c.Request.Header.Get("X-Forwarded-For"); proxies != "" {
			data.ClientInfo["Proxies"] = proxies
		}
		tmpl.Execute(c.Writer, data)
	}

	return gin.HandlerFunc(fn)
}
