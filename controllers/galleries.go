package controllers

import (
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
