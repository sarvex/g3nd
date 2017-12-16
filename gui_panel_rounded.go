package main

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/window"
)

type GuiPanelRounded struct {
	p1 *gui.Panel
	p2 *gui.Panel
	p3 *gui.Panel
}

func init() {
	TestMap["gui.panel_rounded"] = &GuiPanelRounded{}
}

func (t *GuiPanelRounded) Initialize(ctx *Context) {

	// Adds axis helper
	axis := graphic.NewAxisHelper(2)
	ctx.Scene.Add(axis)

	// Panel 1
	t.p1 = gui.NewPanel(100, 50)
	t.p1.SetRoundness(1, 1, 1, 1)
	t.p1.SetPosition(10, 10)
	t.p1.SetMargins(0, 0, 0, 0)
	t.p1.SetBorders(2, 2, 2, 2)
	t.p1.SetBordersColor(math32.NewColor("black"))
	//t.p1.SetPaddings(8, 8, 8, 8)
	t.p1.SetColor(math32.NewColor("white"))
	//	t.p1.SetPaddingsColor(math32.NewColor("green"))
	//t.p1.SetContentSize(100, 100)
	t.p1.SetSize(200, 200)
	ctx.Gui.Add(t.p1)

	// Panel 2
	t.p2 = gui.NewPanel(t.p1.Width(), t.p1.Height())
	t.p2.SetRoundness(0.5, 0.5, 0.5, 0.5)
	t.p2.SetPosition(t.p1.Width()+20.0, 10)
	//t.p2.SetMargins(4, 4, 4, 4)
	t.p2.SetBorders(10, 10, 10, 10)
	t.p2.SetBordersColor(math32.NewColor("red"))
	t.p2.SetPaddings(10, 10, 10, 10)
	t.p2.SetColor(math32.NewColor("white"))
	t.p2.SetPaddingsColor(math32.NewColor("green"))
	t.p2.SetSize(300, 300)
	ctx.Gui.Add(t.p2)

	// Panel 4
	p3 := gui.NewPanel(t.p1.Width(), t.p1.Height())
	p3.SetRoundness(0.5, 0.5, 0.5, 0.5)
	p3.SetPosition(t.p2.Position().X, t.p2.Position().Y+t.p2.Height()+10)
	p3.SetBorders(10, 10, 10, 10)
	p3.SetBordersColor(math32.NewColor("red"))
	p3.SetPaddings(10, 10, 10, 10)
	p3.SetColor(math32.NewColor("white"))
	p3.SetPaddingsColor(math32.NewColor("green"))
	p3.SetSize(400, 200)
	ctx.Gui.Add(p3)

	//	// Panel 3 with several children
	//	t.p3 = gui.NewPanel(400, 200).SetColor(math32.NewColor("gray"))
	//	t.p3.SetPosition(50, 160)
	//	t.p3.SetBorders(6, 6, 6, 6)
	//	t.p3.SetBordersColor(math32.NewColor("black"))
	//	t.p3.SetPaddings(2, 2, 2, 2)
	//	ctx.Gui.Add(t.p3)
	//	p3Event := func(name string, ev interface{}) { log.Debug("Parent:%v", name) }
	//	t.p3.Subscribe(gui.OnCursor, p3Event)
	//	t.p3.Subscribe(gui.OnCursorEnter, p3Event)
	//	t.p3.Subscribe(gui.OnCursorLeave, p3Event)
	//	t.p3.Subscribe(gui.OnMouseDown, p3Event)
	//	t.p3.Subscribe(gui.OnMouseUp, p3Event)
	//
	//	// Child Panel 1
	//	c := gui.NewPanel(50, 50).SetColor(math32.NewColor("blue"))
	//	c.SetPosition(-25, -25)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel 2
	//	c = gui.NewPanel(10, 240).SetColor(math32.NewColor("blue"))
	//	c.SetPosition(100, -10)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//	c2Event := func(name string, ev interface{}) { log.Debug("Child 2:%v", name) }
	//	c.Subscribe(gui.OnCursor, c2Event)
	//	c.Subscribe(gui.OnCursorEnter, c2Event)
	//	c.Subscribe(gui.OnCursorLeave, c2Event)
	//	c.Subscribe(gui.OnMouseDown, c2Event)
	//	c.Subscribe(gui.OnMouseUp, c2Event)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("red"))
	//	c.SetPosition(175, -25)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("red"))
	//	c.SetPosition(375, -25)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("red"))
	//	c.SetPosition(375, 75)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("green"))
	//	c.SetPosition(375, 175)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("green"))
	//	c.SetPosition(175, 175)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("green"))
	//	c.SetPosition(-25, 175)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("green"))
	//	c.SetPosition(-25, 75)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel
	//	c = gui.NewPanel(50, 50).SetColor(math32.NewColor("green"))
	//	c.SetPosition(175, 75)
	//	c.SetBorders(4, 4, 4, 4)
	//	t.p3.Add(c)
	//
	//	// Child Panel from previous child
	//	c1 := gui.NewPanel(20, 20).SetColor(math32.NewColor("white"))
	//	c1.SetPosition(10, 10)
	//	c1.SetBorders(1, 1, 1, 1)
	//	c.Add(c1)
	//
	//	// Image panel 1
	//	im, err := gui.NewImage(ctx.DirData + "/images/tiger1.jpg")
	//	if err != nil {
	//		log.Fatal("%s", err)
	//	}
	//	im.SetPosition(50, 400)
	//	im.SetMargins(0, 0, 0, 0)
	//	im.SetBorders(2, 2, 2, 2)
	//	im.SetBordersColor(math32.NewColor("red"))
	//	im.SetPaddings(6, 6, 6, 6)
	//	im.SetColor(math32.NewColor("white"))
	//	im.SetContentAspectWidth(128)
	//	ctx.Gui.Add(im)
	//
	//	// Image panel 2
	//	im, err = gui.NewImage(ctx.DirData + "/images/tiger1.jpg")
	//	if err != nil {
	//		log.Fatal("%s", err)
	//	}
	//	im.SetPosition(250, 400)
	//	im.SetMargins(0, 0, 0, 0)
	//	im.SetBorders(2, 2, 2, 2)
	//	im.SetBordersColor(math32.NewColor("black"))
	//	im.SetPaddings(0, 0, 40, 40)
	//	im.SetColor(math32.NewColor("white"))
	//	im.SetContentAspectWidth(90)
	//	ctx.Gui.Add(im)
	//
	//	// Subscribe to key events
	//	ctx.Win.Subscribe(window.OnKeyDown, t.onKey)
}

func (t *GuiPanelRounded) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	if kev.Action == window.Release {
		return
	}
	switch kev.Keycode {
	case window.KeyB:
		for _, ichild := range t.p3.Children() {
			pan := ichild.(*gui.Panel)
			pan.SetBounded(!pan.Bounded())
		}
	}

}

func (t *GuiPanelRounded) Render(ctx *Context) {

	//	time := ctx.Win.GetTime()
	//	delta := (math.Sin(time) + 1.0) / 2
	//	maxWidth := 200
	//	maxHeight := 100
	//	t.p1.SetSize(float32(delta)*float32(maxWidth), float32(delta)*float32(maxHeight))
	//	t.p2.SetContentSize(float32(delta)*float32(maxWidth), float32(delta)*float32(maxHeight))
	//	t.p3.SetPosition(50+60*math32.Sin(float32(time)), t.p3.Position().Y)
}