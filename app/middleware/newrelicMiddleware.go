package middleware

import (
	"github.com/gin-gonic/gin"
	nr "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrgin/v1"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type newRelicMiddleware struct {
}

func NewNewRelicMiddleware() *newRelicMiddleware {
	return &newRelicMiddleware{}
}

func (middleware *newRelicMiddleware) Use(engine *gin.Engine, logger *logrus.Logger) {
	newRelicEnabled, parsingErr := strconv.ParseBool(os.Getenv("NEW_RELIC_ENABLED"))
	if parsingErr == nil && newRelicEnabled {
		logger.Infof("New relic enabled:  %t", newRelicEnabled)
		config := nr.NewConfig(os.Getenv("NEW_RELIC_APP_NAME"), os.Getenv("NEW_RELIC_LICENSE_KEY"))
		newRelicApp, newRelicErr := nr.NewApplication(config)
		if newRelicErr == nil {
			engine.Use(nrgin.Middleware(newRelicApp))
		}
	}

}
