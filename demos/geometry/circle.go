package geometry

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["geometry.circle"] = &Circle{}
}

type Circle struct {
	circ1 *graphic.Mesh
	circ2 *graphic.Mesh
	circ3 *graphic.Mesh
}

// Start is called once at the start of the demo.
func (t *Circle) Start(a *app.App) {

	geom1 := geometry.NewCircle(1, 3)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	t.circ1 = graphic.NewMesh(geom1, mat1)
	t.circ1.SetPosition(-1.5, 0, 0)
	a.Scene().Add(t.circ1)

	geom2 := geometry.NewCircle(1, 50)
	mat2 := material.NewStandard(&math32.Color{0, 0, 1})
	t.circ2 = graphic.NewMesh(geom2, mat2)
	t.circ2.SetPosition(0, 0, -0.5)
	a.Scene().Add(t.circ2)

	geom3 := geometry.NewCircle(1, 5)
	mat3 := material.NewStandard(&math32.Color{1, 0, 0})
	t.circ3 = graphic.NewMesh(geom3, mat3)
	t.circ3.SetPosition(1.5, 0, -1.0)
	a.Scene().Add(t.circ3)
}

// Update is called every frame.
func (t *Circle) Update(a *app.App, deltaTime time.Duration) {

	// Rotate at 1 rotation each 5 seconds
	delta := float32(deltaTime.Seconds()) * 2 * math32.Pi / 5
	t.circ1.RotateZ(delta)
	t.circ3.RotateZ(-delta)
}

// Cleanup is called once at the end of the demo.
func (t *Circle) Cleanup(a *app.App) {}