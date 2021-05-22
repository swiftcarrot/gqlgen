package codegen

import (
	"embed"

	"github.com/99designs/gqlgen/codegen/templates"
)

//go:embed *.gotpl
var fs embed.FS

func GenerateCode(data *Data) error {
	return templates.Render(templates.Options{
		Embed:           &fs,
		PackageName:     data.Config.Exec.Package,
		Filename:        data.Config.Exec.Filename,
		Data:            data,
		RegionTags:      true,
		GeneratedHeader: true,
		Packages:        data.Config.Packages,
	})
}
