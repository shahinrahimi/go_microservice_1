package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	if err := app.readJSON(w, r, &requestPayload); err != nil {
		app.errorJSON(w, err)
		log.Println("error at SendMail1: ", err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	if err := app.Mailer.SendSMTPMessage(msg); err != nil {
		app.errorJSON(w, err)
		log.Println("error at SendMail2: ", err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "send to " + requestPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
