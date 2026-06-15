func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/hello", HelloHandler).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}