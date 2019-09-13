package global

import (
	"github.com/johnnyeven/libtools/servicex"
	"github.com/johnnyeven/vehicle-robot-host/client"
	"gocv.io/x/gocv"
)

func init() {
	servicex.SetServiceName("vehicle-robot-host")
	servicex.ConfP(&Config)
}

var Config = struct {
	RobotClient *client.RobotClient
	Window      *gocv.Window `conf:"-"`
}{
	RobotClient: &client.RobotClient{
		RemoteAddr: "robot.profzone.net:9090",
		NodeKey:    "456",
	},
	Window: gocv.NewWindow("Display"),
}
