func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bienvenido a AgmLanding"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Ruta no encontrada"))
}