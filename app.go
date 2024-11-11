package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML file from the "public" folder
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, "Error loading the page", http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		num1, _ := strconv.ParseFloat(r.FormValue("num1"), 64)
		exponent, _ := strconv.ParseFloat(r.FormValue("num2"), 64) // Change exp to exponent
		result := fmt.Sprintf("%.0f", pow(num1, exponent))         // Show as integer if it's a whole number
		tmpl.Execute(w, map[string]interface{}{"Result": result})
	} else {
		tmpl.Execute(w, nil)
	}
}

// Function to calculate the exponent using math.Pow
func pow(base, exponent float64) float64 { // Change exp to exponent
	return math.Pow(base, exponent)
}

func main() {
	// Serve static files from the "public" folder
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Route to display the calculator
	http.HandleFunc("/", indexHandler)

	// Start the server on port 6005
	fmt.Println("Server running at http://localhost:6005")
	http.ListenAndServe(":6005", nil)
}
