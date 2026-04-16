package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"one-api/common"
	"one-api/middleware"
	"one-api/model"
	"one-api/router"
)

func main() {
	common.SetupLogger()
	common.SysLog("New API starting...")

	// Initialize database
	err := model.InitDB()
	if err != nil {
		common.FatalLog("failed to initialize database: " + err.Error())
	}
	defer func() {
		err := model.CloseDB()
		if err != nil {
			common.SysError("failed to close database: " + err.Error())
		}
	}()

	// Initialize Redis if configured
	err = common.InitRedisClient()
	if err != nil {
		common.SysError("failed to initialize Redis: " + err.Error())
	}

	// Initialize options from database
	model.InitOptionMap()
	common.SysLog("options initialized")

	// Initialize token encoder
	common.InitTokenEncoders()

	// Setup Gin
	// Default to release mode unless explicitly set to debug for local development
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.RequestId())
	middleware.SetUpLogger(server)

	// Setup routes
	router.SetRouter(server)

	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}

	common.SysLog(fmt.Sprintf("server started on http://localhost:%s", port))

	if err := server.Run(":" + port); err != nil {
		common.FatalLog("failed to start server: " + err.Error())
	}
}
