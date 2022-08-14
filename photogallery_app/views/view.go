package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	TemplateDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layouts  string
}

func NewView(layout string, files ...string) *View {
	files = append(files,
		layoutFiles()...)

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
		Layouts:  layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layouts, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(TemplateDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
