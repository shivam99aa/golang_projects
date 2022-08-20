package controllers

import "golang_projects/photogallery_app/views"

type Static struct {
	HomeView    *views.View
	ContactView *views.View
	FAQView     *views.View
}

func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
		FAQView:     views.NewView("bootstrap", "views/static/faq.gohtml"),
	}
}
