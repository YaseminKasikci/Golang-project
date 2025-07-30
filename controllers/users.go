package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, " Email: ", r.FormValue("email"))
	fmt.Fprint(w, " Password: ", r.FormValue("password"))
	fmt.Fprint(w, " sex: ", r.FormValue("sex"))
	
	interest := r.Form["interest[]"]
	fmt.Fprint(w, " interest: ")
	for _,i := range interest{
		fmt.Fprint(w,", " ,i)
	}

	fmt.Fprint(w, " avatar: ", r.FormValue("avatar"))

}
