package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Get the value of the PORT environment variable, or set a default value of "3000"
	port := os.Getenv("PORT")
	if port == "" {
		os.Setenv("PORT", "3000")
		port = "3000"
	}

	apiBaseURL := os.Getenv("API_BASE_URL")
	apiKey := os.Getenv("API_KEY")

	// Set default value for API base URL if not specified in the .env file
	if apiBaseURL == "" {
		apiBaseURL = "https://api.status.finance"
	}

	// Set up the HTTP handler that proxies requests to the API
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		// Create a new request to the API endpoint
		req, err := http.NewRequest(r.Method, apiBaseURL+"/transactions", r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Do not proxy all headers for privacy reasons
		//req.Header = r.Header
		req.URL.RawQuery = r.URL.RawQuery

		// Set the x-api-key header on the new request
		req.Header.Set("x-api-key", apiKey)

		// Send the new request and get the response
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// Read the response body from the API into a []byte
		body, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the response from the API to the original response writer
		w.WriteHeader(res.StatusCode)
		if _, err := w.Write(body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	log.Println("Listening on http://localhost:" + port + "...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
