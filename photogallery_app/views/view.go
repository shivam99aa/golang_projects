package views

import "html/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View{
	files = append(files, "views/layouts/footer.gohtml")

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
	}
}