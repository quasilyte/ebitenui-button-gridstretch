package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

func main() {
	if err := ebiten.RunGame(&myGame{}); err != nil {
		panic(err)
	}
}

type myGame struct {
	initialized bool
	ui          *ebitenui.UI
}

func (g *myGame) Update() error {
	if !g.initialized {
		g.Init()
		g.initialized = true
	}
	g.ui.Update()
	return nil
}

func (g *myGame) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *myGame) Layout(_, _ int) (int, int) {
	return 1920 / 4, 1080 / 4
}

func (g *myGame) Init() {
	buttonImage := &widget.ButtonImage{
		Idle:     NineSliceImage(LoadImage("_data/button-idle.png"), 25, 20),
		Hover:    NineSliceImage(LoadImage("_data/button-hover.png"), 25, 20),
		Pressed:  NineSliceImage(LoadImage("_data/button-pressed.png"), 25, 20),
		Disabled: NineSliceImage(LoadImage("_data/button-disabled.png"), 25, 20),
	}

	normalTextColor := color.NRGBA{R: 0xdf, G: 0xf4, B: 0xff, A: 0xff}
	disabledTextColor := color.NRGBA{R: 0x5a, G: 0x7a, B: 0x91, A: 0xff}
	buttonTextColor := &widget.ButtonTextColor{
		Idle:     normalTextColor,
		Disabled: disabledTextColor,
	}

	newButton := func(label string) *widget.Button {
		opts := []widget.ButtonOpt{
			widget.ButtonOpts.Text(label, bitmapfont.Face, buttonTextColor),
			widget.ButtonOpts.Image(buttonImage),
			widget.ButtonOpts.TextPadding(widget.Insets{
				Left:  30,
				Right: 30,
			}),
		}
		return widget.NewButton(opts...)
	}

	root := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))

	rows := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Spacing(0, 5),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true, true, false}),
		)),
	)
	root.AddChild(rows)

	rows.AddChild(widget.NewText(
		widget.TextOpts.Text("text", bitmapfont.Face, normalTextColor),
	))

	rows.AddChild(newButton("button1"))
	rows.AddChild(newButton("button2"))

	g.ui = &ebitenui.UI{}
	g.ui.Container = root
}
