package postprocess

import "raytracing/camera"

type NoProcessing struct{}

func (n *NoProcessing) Process(img camera.Image) camera.Image {
	return img
}
