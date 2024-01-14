package api

import (
	"fmt"
	"net/http"
	cors "txcrawlerdegen/api/middleware"
	"txcrawlerdegen/api/route"
	"txcrawlerdegen/api/types"
	"txcrawlerdegen/constants"
	"txcrawlerdegen/internal/appconfig"
	"txcrawlerdegen/pkg/logger"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	// We can set to ginReleaseMode here when scaling to demo/prod

	router := gin.New()

	router.HandleMethodNotAllowed = true

	router.Use(cors.CORSMiddleware())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": types.StatusError, "error": &types.ErrorResponse{
			Code:    constants.ERROR_TYPES[constants.METHOD_NOT_ALLOWED].HttpStatus,
			Message: constants.METHOD_NOT_FOUND_MESSAGE,
		}})
	})
	logger.Info(fmt.Sprintf("PORT: %s HOST: %s", appconfig.PORT, appconfig.HOST))

	route.Register(router)

	router.Run(":" + appconfig.PORT)

}
