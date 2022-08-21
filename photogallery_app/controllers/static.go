package controllers

import "golang_projects/photogallery_app/views"

type Static struct {
	HomeView    *views.View
	ContactView *views.View
	FAQView     *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "static/home"),
		ContactView: views.NewView("bootstrap", "static/contact"),
		FAQView:     views.NewView("bootstrap", "static/faq"),
	}
}
