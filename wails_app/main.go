package main

import (
	"embed"
	"log"
	repositories "wails_app/repository"
	"wails_app/services"

	_ "github.com/lib/pq"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"wails_app/database"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	// Create an instance of the app structure
	app := NewApp(todoService)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "wails_app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
