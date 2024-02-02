package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadi-1024/straysafe/internals/models"
)

type jsonRequestPayload struct {
	To string `json:"to"`
	From string `json:"from"`
	Content string `json:"content"`
}

func SendMailPostHandler(c chan models.MailData) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request)  {
		var reqJson jsonRequestPayload
		err := json.NewDecoder(r.Body).Decode(&reqJson)
		if err != nil {
			log.Println(err)
		}
		
		mailPayload := models.MailData{
			To: []string{reqJson.To},
			From: reqJson.From,
			Content: reqJson.Content,
		}

		c <- mailPayload

		w.Write([]byte("OK"))
	}
}