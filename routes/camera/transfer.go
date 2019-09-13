package camera

import (
	"bytes"
	"github.com/henrylee2cn/teleport"
	"github.com/johnnyeven/vehicle-robot-host/client"
	"github.com/johnnyeven/vehicle-robot-host/global"
	"github.com/johnnyeven/vehicle-robot-host/modules"
	"github.com/sirupsen/logrus"
	"image/jpeg"
)

func (r *Camera) Transfer(req *client.CameraRequest) *tp.Status {
	buf := bytes.NewReader(req.Frame)
	img, err := jpeg.Decode(buf)
	if err != nil {
		logrus.Errorf("jpeg.Decode err: %v", err)
		return nil
	}
	mat, err := modules.ConvertImageToMat(img)
	if err != nil {
		logrus.Errorf("modules.ConvertImageToMat err: %v", err)
		return nil
	}
	global.Config.Window.IMShow(mat)
	global.Config.Window.WaitKey(1)

	return nil
}
