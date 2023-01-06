package frontend

import (
	"embed"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:dist
	dist embed.FS
	//go:embed dist/index.html
	indexHTML embed.FS
	//nolint:typecheck
	distDirFS = echo.MustSubFS(dist, "dist")
	//nolint:typecheck
	distIndexHtml = echo.MustSubFS(indexHTML, "dist")
)

//nolint:typecheck
func RegisterHandlers(e *echo.Echo) {
	e.FileFS("/", "index.html", distIndexHtml)
	e.StaticFS("/", distDirFS)
}
