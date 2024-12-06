package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"log"
	"os"
	"time"
)

var progress float32
var progressIncrementer chan float32

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

	progressIncrementer = make(chan float32)
	go func() {
		for {
			time.Sleep(time.Second / 25)
			progressIncrementer <- 0.004

		}
	}()

	app.Main()
}

type C = layout.Context
type D = layout.Dimensions

func draw(w *app.Window) error {

	var ops op.Ops

	var startButton widget.Clickable
	var isBoiling bool

	theme := material.NewTheme()

	go func() {
		for p := range progressIncrementer {
			if isBoiling && progress < 1 {
				progress += p
				w.Invalidate()
			}
		}
	}()

	// listen for events in the window.
	for {
		// first grab the event

		// then detect the type
		switch e := w.Event().(type) {
		// re-render
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			if startButton.Clicked(gtx) {
				fmt.Println("Button was Clicked!")
				isBoiling = !isBoiling
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						bar := material.ProgressBar(theme, progress)
						return bar.Layout(gtx)
					}),

				layout.Rigid(

					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Right:  unit.Dp(25),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(25),
						}

						return margins.Layout(gtx,
							func(gtx C) D {
								var text string

								if isBoiling {
									text = "Stop"
								} else {
									text = "Start"
								}

								btn := material.Button(theme, &startButton, text)

								return btn.Layout(gtx)
							})

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
