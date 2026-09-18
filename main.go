package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "WarpScout Chain",
		Description: "Warp endpoint scanner and configuration generator",
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	service := NewWarpScoutService(app)
	app.RegisterService(application.NewService(service))

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:          "main",
		Title:         "WarpScout Chain v1.0",
		Width:         740,
		Height:        580,
		DisableResize: true,
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
