package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	port  = "80"
	token string
)

func init() {
	godotenv.Load()

	token = os.Getenv("SLACK_TOKEN")
	if "" == token {
		panic("SLACK_TOKEN is not set!")
	}

	if "" != os.Getenv("PORT") {
		port = os.Getenv("PORT")
	}
}

func main() {
	http.HandleFunc("/circle-ci", ciBuildHandler)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

func ciBuildHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if token != r.FormValue("token") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)
	circleResponse, err := buildBranch(text)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("```%s```", circleResponse),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResp))
}

func buildBranch(text string) (string, error) {
	// do circle ci api call

	// return ephemeral response
	return string("got it!"), nil
}
