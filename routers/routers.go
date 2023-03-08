package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sagar4797/microservice/controller"
)

func Setup(cf *controller.ControllerFactory) (engine *gin.Engine) {

	switch os.Getenv("ENV") {
	case "production", "staging":
		// Disable default logger(stdout/stderr)
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	default:
		engine = gin.Default()
	}

	// engine.Use(middleware.Logger())
	// engine.Use(middleware.Cors())

	return
}
