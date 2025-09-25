package main

import (
	"context"
	"embed"
	"karina/backend/helpers"
	"karina/backend/languages"
	"karina/backend/models"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := NewApp()
	modelService := &models.Service{}
	fileService := &helpers.Service{}
	languageService := &languages.Service{}

	os.MkdirAll("data/problems", 0755)

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
			modelService.Initialize(ctx)
			fileService.Initialize(ctx)
			languageService.Initialize(ctx)
		},
		Bind: []interface{}{
			app, modelService, fileService, languageService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
