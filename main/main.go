package main

import (
	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		count := 500

		// create new window
		w := new(app.Window)
		w.Option(app.Title("Go GUI Test"))
		w.Option(app.Size(unit.Dp(count), unit.Dp(count)))
		//w.Option(app.MinSize(unit.Dp(0), unit.Dp(0)))

		var ops op.Ops

		var startButton widget.Clickable

		theme := material.NewTheme()

		// listen for events in the window.
		for {
			// first grab the event
			evt := w.Event()

			// then detect the type
			switch typ := evt.(type) {
			// re-render
			case app.FrameEvent:
				gtx := app.NewContext(&ops, typ)

				btn := material.Button(theme, &startButton, "Start")

				btn.Layout(gtx)
				typ.Frame(gtx.Ops)

			// exit event
			case app.DestroyEvent:
				count -= 10
				w.Option(app.Size(unit.Dp(count), unit.Dp(count)))
			}
		}
	}()
	app.Main()
}
