package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")

	name(w, r)
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hangman")

}

func Setting(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "setting")

}

func name(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Analyser les données du formulaire
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
			return
		}

		// Récupérer la valeur du champ "name" du formulaire
		name := r.Form.Get("name")

		// Faire quelque chose avec la valeur récupérée, par exemple, l'afficher dans la console
		fmt.Println("Name:", name)

		// Vous pouvez également renvoyer une réponse à l'utilisateur, par exemple, en affichant la valeur sur la page
		w.Write([]byte("Votre nom est : " + name))
	}
}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(filepath.Join("html", tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
