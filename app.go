package main

import (
	"log"
	"net/http"
)

func main() {
	// Ruta para servir archivos estáticos (imágenes, CSS, JavaScript)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rutas para servir los archivos HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/menu.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/menu.html")
	})

	http.HandleFunc("/service.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/service.html")
	})

	http.HandleFunc("/contact.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/contact.html")
	})

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
