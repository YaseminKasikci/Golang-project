package controllers

import (
	"fmt"
	"github/yaseminkasikci/lenslocked/context"
	"github/yaseminkasikci/lenslocked/models"
	"net/http"
)

type Galleries struct {
	Templates struct {
		New Template
	}
	GalleryService *models.GalleryService
}

func (g Galleries) New(w http.ResponseWriter, r *http.Request) {
	//struc sert à passer les données au template
	var data struct {
		Title string
	}
	// on reccupere la valuer du champs title du formulaore de la requete http
	data.Title = r.FormValue("title")
	// data données injecter dans la page Templace exécuté
	//cette ligne “remplit” le HTML du template New avec les données du formulaire, puis l’envoie au navigateur.
	g.Templates.New.Execute(w, r, data)
}

func (g Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var data struct {
		UserID int
		Title  string
	}
	data.UserID = context.User(r.Context()).ID
	data.Title = r.FormValue("title")

	gallery, err := g.GalleryService.Create(data.Title, data.UserID)
	if err != nil {
		g.Templates.New.Execute(w, r, data, err)
		return
	}
	editPath := fmt.Sprintf("/galleries/%d/edit", gallery.ID)
	http.Redirect(w, r, editPath, http.StatusFound )
}
