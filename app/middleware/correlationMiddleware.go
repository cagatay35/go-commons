package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-commons/app/util"
)

type correlationIdMiddleware struct {
}

func NewCorrelationIdMiddleware() *correlationIdMiddleware {
	return &correlationIdMiddleware{}
}

func (middleware *correlationIdMiddleware) SetCorrelationId(c *gin.Context) {

	correlationId := c.GetHeader(util.CORRELATION_ID_STRING)
	if correlationId == util.EMPTY_STRING {
		correlationId = uuid.New().String()
	}
	c.Set(util.CORRELATION_ID_STRING, correlationId)
	c.Writer.Header().Set(util.CORRELATION_ID_STRING, correlationId)

}
