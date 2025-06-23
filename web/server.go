package web

import (
	"ascii-art-web-dockerize/helpers"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func RenderError(w http.ResponseWriter, code int, title, message string) {
	w.WriteHeader(code)
	data := struct {
		Code    int
		Title   string
		Message string
	}{
		Code:    code,
		Title:   title,
		Message: message,
	}
	templates.ExecuteTemplate(w, "error.html", data)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	var text, banner string

	switch r.Method {
	case http.MethodGet:
		text = r.URL.Query().Get("text")
		banner = r.URL.Query().Get("banner")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			RenderError(w, http.StatusBadRequest, "Bad Request", "Failed to parse form")
			return
		}
		text = r.FormValue("text")
		banner = r.FormValue("banner")
	default:
		RenderError(w, http.StatusMethodNotAllowed, "Not Allowed", "Method not allowed")
		return
	}

	if text == "" || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
		RenderError(w, http.StatusBadRequest, "Bad Request", "Invalid input or banner")
		return
	}

	if !helpers.AsciiValid(text) {
		RenderError(w, http.StatusBadRequest, "Bad Request", "Invalid input! The character isn't ASCII")
		return
	}

	art, err := helpers.DrawAscii(text, banner)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Server Error", err.Error())
		return
	}

	data := struct{ Art string }{Art: art}
	if err := templates.ExecuteTemplate(w, "result.html", data); err != nil {
		RenderError(w, http.StatusInternalServerError, "Server Error", "Failed to render result.")
	}
}
