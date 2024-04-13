package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func Test_request_by_page(t *testing.T) {
	file, err := os.OpenFile("../app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
		t.Error(err)
	}
	log.SetOutput(file)
	i := 0

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/Mtaste/API/getRecipeByPage/%d", i+1))
	if err != nil {
		log.Fatalf("Failed to get info to %d page\n\tERROR: %s", i+1, err)
		t.Error(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to get info to %d page\n\tERROR: %s", i+1, err)
			t.Error(err)
		}
		strbody := string(body)
		var doc interface{}
		jsoniter.UnmarshalFromString(strbody, &doc)
		fmt.Println(doc)
	}

}
func TestMain(m *testing.M) {
	go RunServer()
	code := m.Run()
	os.Exit(code)
}
