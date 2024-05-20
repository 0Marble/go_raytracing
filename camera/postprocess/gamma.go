package postprocess

import (
	"raytracing/camera"
)

type GammaCorrection struct {
	next  camera.PostProcess
	gamma float32
}

func InitGammaCorrection(gamma float32, nextStep camera.PostProcess) GammaCorrection {
	return GammaCorrection{nextStep, gamma}
}

func (g *GammaCorrection) Process(img camera.Image) camera.Image {
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			idx := y*img.Width + x
			img.Pixels[idx] = img.Pixels[idx].PowNum(g.gamma)
		}
	}

	return g.next.Process(img)
}
