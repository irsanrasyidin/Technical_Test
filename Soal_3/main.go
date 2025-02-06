package main

import (
	"Soal_3/config"
	"Soal_3/internal/db"
	"Soal_3/internal/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Membaca body SOAP request
	var request models.SOAPRequest
	err := xml.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse SOAP request", http.StatusBadRequest)
		return
	}

	// Menampilkan userID yang diterima dari request SOAP
	userIDStr := request.Body.GetUserDetails.UserID
	fmt.Println("Received userID:", userIDStr)

	// Konversi userID dari string ke int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID format", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	dbInstance := config.InitDB()

	// Ambil user dari database berdasarkan ID
	user, err := db.GetUserByID(dbInstance, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirimkan response dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Setup HTTP route
	http.HandleFunc("/", getUserHandler)

	// Jalankan server
	fmt.Println("Server berjalan di http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
