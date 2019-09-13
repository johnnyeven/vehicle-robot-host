package routes

import (
	"github.com/johnnyeven/vehicle-robot-host/global"
	"github.com/johnnyeven/vehicle-robot-host/routes/camera"
)

func InitRouters() {
	server := global.Config.RobotClient
	server.RegisterPushRouter(&camera.Camera{})
}
