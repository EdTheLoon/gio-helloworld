package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func main() {
	// Start a goroutine to create the window elements
	go func() {
		// Define a new Window.
		window := new(app.Window)

		// Calls the run function which creates and draws the GUI elements
		err := run(window)
		if err != nil {
			log.Fatal(err) // Close program with a fatal error
		}
		os.Exit(0) // Successful code completion, exit app
	}()

	// This should always be run last in most cases as it blocks further code execution
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// The graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define a large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio!")

			// Change the colour of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the text alignment of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
