package main

import (
	"net/http"

	"github.com/byuoitav/adcp-control-microservice/handlers"
	"github.com/byuoitav/authmiddleware"
	"github.com/byuoitav/hateoas"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := ":8012"
	router := echo.New()
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.CORS())

	// Use the `secure` routing group to require authentication
	secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))
	router.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))

	secure.GET("/:address/volume/set/:level", handlers.SetVolume)
	secure.GET("/:address/power/on", handlers.PowerOn)
	secure.GET("/:address/power/standby", handlers.PowerStandby)
	secure.GET("/:address/volume/mute", handlers.Mute)
	secure.GET("/:address/volume/unmute", handlers.UnMute)
	secure.GET("/:address/display/blank", handlers.DisplayBlank)
	secure.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	secure.GET("/:address/input/:port", handlers.SetInputPort)

	//status endpoints
	secure.GET("/:address/volume/level", handlers.VolumeLevel)
	secure.GET("/:address/volume/mute/status", handlers.MuteStatus)
	secure.GET("/:address/power/status", handlers.PowerStatus)
	secure.GET("/:address/display/status", handlers.BlankedStatus)
	secure.GET("/:address/input/current", handlers.InputStatus)

	//------------------
	//Pooled endpoints
	//------------------
	secure.GET("/pooled/:address/volume/set/:level", handlers.SetVolumePooled)
	secure.GET("/pooled/:address/power/on", handlers.PowerOnPooled)
	secure.GET("/pooled/:address/power/standby", handlers.PowerStandbyPooled)
	secure.GET("/pooled/:address/volume/mute", handlers.MutePooled)
	secure.GET("/pooled/:address/volume/unmute", handlers.UnMutePooled)
	secure.GET("/pooled/:address/display/blank", handlers.DisplayBlankPooled)
	secure.GET("/pooled/:address/display/unblank", handlers.DisplayUnBlankPooled)
	secure.GET("/pooled/:address/input/:port", handlers.SetInputPortPooled)

	//status endpoints
	secure.GET("/pooled/:address/volume/level", handlers.VolumeLevelPooled)
	secure.GET("/pooled/:address/volume/mute/status", handlers.MuteStatusPooled)
	secure.GET("/pooled/:address/power/status", handlers.PowerStatusPooled)
	secure.GET("/pooled/:address/display/status", handlers.BlankedStatusPooled)
	secure.GET("/pooled/:address/input/current", handlers.InputStatusPooled)
	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
