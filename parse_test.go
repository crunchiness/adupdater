package main

import (
	"gopkg.in/xmlpath.v2"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, _ := os.Open("parse_test.html")
	adRoot, _ := xmlpath.ParseHTML(file)
	adData, _ := parseAdPage(adRoot)
	expectedValue := "26975"
	if adData["photo"] != expectedValue {
		t.Errorf("Photo id parsed wrong should be %s, got %s", expectedValue, adData["photo"])
	}
}
