package main

import (
	"ascii-art-web-dockerize/web"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Создаём свой ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/icon/favicon.ico")
	})

	// Регистрируем маршруты на этом mux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			web.IndexHandler(w, r)
		case "/ascii-art":
			web.AsciiArtHandler(w, r)
		default:
			web.RenderError(w, http.StatusNotFound, "Page Not Found", "The page you are looking for it doesn't exist.")
		}
	})

	// Настраиваем обработку статических файлов
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		// отрезаем префикс
		rel := strings.TrimPrefix(r.URL.Path, "/static/")
		fullPath := filepath.Join("static", rel)

		// если запрошен каталог или файл не найден – отдаём свой шаблон ошибки
		info, err := os.Stat(fullPath)
		if err != nil || info.IsDir() {
			web.RenderError(w, http.StatusNotFound, "Page Not Found", "The page you are looking for it doesn't exist.")
			return
		}

		// иначе – отдаем файл
		http.ServeFile(w, r, fullPath)
	})

	log.Println("Starting at this server: http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
