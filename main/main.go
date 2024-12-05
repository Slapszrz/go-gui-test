package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
	"os"
)

func main() {
	go func() {
		count := 500

		// create new window
		w := new(app.Window)
		w.Option(app.Title("Go GUI Test"))
		w.Option(app.Size(unit.Dp(count), unit.Dp(count)))
		//w.Option(app.MinSize(unit.Dp(0), unit.Dp(0)))

		if err := draw(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

type C = layout.Context
type D = layout.Dimensions

func draw(w *app.Window) error {

	var ops op.Ops

	var startButton widget.Clickable

	theme := material.NewTheme()

	// listen for events in the window.
	for {
		// first grab the event

		// then detect the type
		switch e := w.Event().(type) {
		// re-render
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(theme, &startButton, "Start")

						return btn.Layout(gtx)
					}),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			e.Frame(gtx.Ops)

		// exit event
		case app.DestroyEvent:
			return e.Err
		}
	}
}
