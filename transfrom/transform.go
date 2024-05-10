package transfrom

import "fmt"

import "raytracing/linal"

func Hello() {
	fmt.Println("Hello from Transform!")
}

type Transform struct {
	Translation linal.Vec3
	Scale       linal.Vec3
	Rotation    linal.Vec3
}
