package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-commons/app/util"
	"strings"
)

type loggingMiddleware struct {
	Logger *logrus.Logger
}

func NewLoggingMiddleware(logger *logrus.Logger) *loggingMiddleware {
	return &loggingMiddleware{
		Logger: logger,
	}
}

func (middleware *loggingMiddleware) Log(c *gin.Context) {

	if !strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		correlationId, _ := c.Get(util.CORRELATION_ID_STRING)

		appName := c.GetHeader(util.APP_NAME)

		fields := make(map[string]interface{}, 4)

		fields[util.CORRELATION_ID_STRING] = correlationId
		fields[util.METHOD] = c.Request.Method
		fields[util.PATH] = c.Request.URL.Path
		fields[util.APP_NAME] = appName

		log := middleware.Logger.WithFields(fields)
		log.Infof("Started")
		c.Next()

		fields[util.STATUS_CODE] = c.Writer.Status()
		fields[util.BODY] = blw.body.String()
		log = log.WithFields(fields)
		log.Infof("Finished")
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
