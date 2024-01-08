package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SendEmail(w http.ResponseWriter, r *http.Request) {

	envLoad := godotenv.Load(".env")
	if envLoad != nil {
		log.Fatalf("Error loading .env file: %s", envLoad)
	}

	EMAIL_RECIPIENT := os.Getenv("EMAIL")
	PASS_RECIPIENT := os.Getenv("PASSWORD")

	var emailRequest EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Konfigurasi pengiriman email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", EMAIL_RECIPIENT)
	mailer.SetHeader("To", emailRequest.To)
	mailer.SetHeader("Subject", emailRequest.Subject)
	mailer.SetBody("text/plain", emailRequest.Body)

	// Mengirim email
	dialer := gomail.NewDialer("smtp.gmail.com", 587, EMAIL_RECIPIENT, PASS_RECIPIENT)
	err = dialer.DialAndSend(mailer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Email berhasil dikirim")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/send-email", SendEmail).Methods("POST")

	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
