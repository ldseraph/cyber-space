// Package ui handles the PocketBase Admin frontend embedding.
package ui

import (
	"embed"

	"github.com/labstack/echo/v5"
)

//go:embed dist/spa/*
var f embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS = echo.MustSubFS(f, "dist/spa/")
