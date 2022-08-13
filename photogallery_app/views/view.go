package views

import "html/template"

type View struct {
	Template *template.Template
	Layouts  string
}

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/footer.gohtml")

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
		Layouts:  layout,
	}
}
