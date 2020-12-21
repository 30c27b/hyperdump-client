package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/30c27b/hyperdump-client/internal/auth"
)

// Ctx represents a dump server
type Ctx struct {
	server     string
	credential string
}

// PushRequest represents the body of a /push request
type PushRequest struct {
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	Data      string    `json:"data"`
}

// Push creates and uploads a new dump
func Push(input io.ReadWriter, output io.ReadWriter, custemKey string) {
	server, token := auth.Request()

	fmt.Printf("Server: %s\nToken: %s\n", server, token)

	n, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("Error: cannot read input file")
	}

	d := base64.StdEncoding.EncodeToString(n)

	var pr = PushRequest{"test", "txt", "hello", time.Now(), d}

	jr, err := json.Marshal(pr)
	if err != nil {
		log.Fatal("Error: cannot create json object")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodPut, server+"/push", bytes.NewBuffer(jr))

	if err != nil {
		log.Fatal("Error: could not create the request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorisation", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error: could not create the dump:", err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}

// Pull requests and downloads a given dump
func Pull(input io.ReadWriter, output io.ReadWriter) {
	// TODO
}
