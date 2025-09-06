package main

import (
	"context"
	"embed"
	"karina/backend/models"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	problemService := &models.ProblemService{}
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "karina",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			problemService.Initialize(ctx)
		},
		Bind: []interface{}{
			app, problemService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
