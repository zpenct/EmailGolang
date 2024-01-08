package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func SendEmail(w http.ResponseWriter, r *http.Request) {

	EMAIL_RECIPIENT := os.Getenv("EMAIL")
	PASS_RECIPIENT := os.Getenv("PASSWORD")

	var emailRequest EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %s", err.Error()), http.StatusBadRequest)
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

	response := ResponseMessage{
		Message: "Email berhasil dikirim",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/send-email", SendEmail).Methods("POST")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is required")
	}

	server := &http.Server{
		Addr: port,
		Handler: r,
	}

	log.Println("Server starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
