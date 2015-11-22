package main

import (
	"gopkg.in/xmlpath.v2"
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("parse_test.html")
	if err != nil {
		log.Fatal("asdf", err)
	}
	adRoot, err := xmlpath.ParseHTML(file)
	if err != nil {
		log.Fatal(err)
	}
	adData, err := parseAdPage(adRoot)
	if err != nil {
		log.Fatal(err)
	}
	expectedValue := "26975"
	if adData["photo"] != expectedValue {
		t.Errorf("Photo id parsed wrong should be %s, got %s", expectedValue, adData["photo"])
	}
}
