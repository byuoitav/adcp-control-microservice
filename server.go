package main

import (
	"net/http"

	"github.com/byuoitav/adcp-control-microservice/handlers"
	"github.com/byuoitav/common"
	"github.com/byuoitav/common/log"
	"github.com/byuoitav/hateoas"
	"github.com/labstack/echo"
)

func main() {
	port := ":8012"
	router := common.NewRouter()
	// Use the `secure` routing group to require authentication
	//secure := router.Group("", echo.WrapMiddleware(authmiddleware.Authenticate))

	router.GET("/", echo.WrapHandler(http.HandlerFunc(hateoas.RootResponse)))

	router.GET("/:address/volume/set/:level", handlers.SetVolume)
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.PowerStandby)
	router.GET("/:address/volume/mute", handlers.Mute)
	router.GET("/:address/volume/unmute", handlers.UnMute)
	router.GET("/:address/display/blank", handlers.DisplayBlank)
	router.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	router.GET("/:address/input/:port", handlers.SetInputPort)

	//status endpoints
	router.GET("/:address/volume/level", handlers.VolumeLevel)
	router.GET("/:address/volume/mute/status", handlers.MuteStatus)
	router.GET("/:address/power/status", handlers.PowerStatus)
	router.GET("/:address/display/status", handlers.BlankedStatus)
	router.GET("/:address/input/current", handlers.InputStatus)

	//------------------
	//Pooled endpoints
	//------------------
	router.GET("/pooled/:address/volume/set/:level", handlers.SetVolumePooled)
	router.GET("/pooled/:address/power/on", handlers.PowerOnPooled)
	router.GET("/pooled/:address/power/standby", handlers.PowerStandbyPooled)
	router.GET("/pooled/:address/volume/mute", handlers.MutePooled)
	router.GET("/pooled/:address/volume/unmute", handlers.UnMutePooled)
	router.GET("/pooled/:address/display/blank", handlers.DisplayBlankPooled)
	router.GET("/pooled/:address/display/unblank", handlers.DisplayUnBlankPooled)
	router.GET("/pooled/:address/input/:port", handlers.SetInputPortPooled)

	//status endpoints
	router.GET("/pooled/:address/volume/level", handlers.VolumeLevelPooled)
	router.GET("/pooled/:address/volume/mute/status", handlers.MuteStatusPooled)
	router.GET("/pooled/:address/power/status", handlers.PowerStatusPooled)
	router.GET("/pooled/:address/display/status", handlers.BlankedStatusPooled)
	router.GET("/pooled/:address/input/current", handlers.InputStatusPooled)

	router.PUT("/log-level/:level", log.SetLogLevel)
	router.GET("/log-level", log.GetLogLevel)

	server := http.Server{
		Addr:           port,
		MaxHeaderBytes: 1024 * 10,
	}

	router.StartServer(&server)
}
