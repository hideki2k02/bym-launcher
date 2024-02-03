package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	onStartup := func(ctx context.Context) {
		// patcher()
		app.startup(ctx)
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "Backyard Monster Refitted",
		Width:         344,
		Height:        560,
		DisableResize: true,
		Frameless:     true,
		//CSSDragProperty: "widows",
		//CSSDragValue:    "drag",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 255},
		OnStartup:        onStartup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
