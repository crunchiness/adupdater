package main

import (
	"gopkg.in/xmlpath.v2"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	file, err := os.Open("parse_test.html")
	if err != nil {
		log.Fatal(err)
	}
	adRoot, err := xmlpath.ParseHTML(file)
	if err != nil {
		log.Fatal(err)
	}
	adData, err := parseAdPage(adRoot)
	if err != nil {
		log.Fatal(err)
	}
	payload, err := generateRequestPayload(adData)
	if err != nil {
		log.Fatal(err)
	}
	expected := "Content-Disposition: form-data; name=\"photo_radio\"\n\n26975"
	if !strings.Contains(payload, expected) {
		t.Errorf("Does not contain the thing")
	}
}
