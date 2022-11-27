package generate

import (
	"fmt"
	"frames/util"

	"github.com/fogleman/gg"
)

func Frame() {

	width := 1280 / 4 //1920
	height := 720 / 4 //1080

	frameCount := 1

	dc := gg.NewContext(width, height)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.LoadFontFace("andale.ttf", 36)
	dc.SetRGB(1, 1, 1)
	dc.DrawString("testing", 60, 60)

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
	util.OpenData()

}
