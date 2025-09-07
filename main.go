package main

import (
	"context"
	"embed"
	"karina/backend/helpers"
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
	problemService := &models.ProblemService{}
	participantService := &models.ParticipantService{}
	fileService := &helpers.FileService{}

	os.MkdirAll("data/problems", 0755)
	os.MkdirAll("data/participants", 0755)

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
			participantService.Initialize(ctx)
			fileService.Initialize(ctx)
		},
		Bind: []interface{}{
			app, problemService, participantService, fileService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
