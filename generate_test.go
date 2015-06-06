package main

import (
	"testing"
	"gopkg.in/xmlpath.v2"
    "os"
    "strings"
)

func TestGenerate(t *testing.T) {
    file, _ := os.Open("parse_test.html")
    adRoot, _ := xmlpath.ParseHTML(file)
    adData, _ := parseAdPage(adRoot)
    payload, _ := generateRequestPayload(adData)
    expected := "Content-Disposition: form-data; name=\"photo_radio\"\n\n26975"
    if !strings.Contains(payload, expected) {
        t.Errorf("Does not contain the thing")
    }
}
