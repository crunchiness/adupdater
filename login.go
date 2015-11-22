package main

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Jar: cookieJar,
	}

	loginValues := url.Values{}
	loginValues.Set("login[username]", username)
	loginValues.Set("login[password]", password)

	// login
	fmt.Println("Logging in...")
	loginResp, err := client.PostForm("http://www.sena.lt/user/login", loginValues)
	if err != nil {
		log.Fatal(err)
	}
	defer loginResp.Body.Close()

	turl, err := url.Parse("http://www.sena.lt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cookieJar.Cookies(turl))

	// retrieve the first page of ads
	fmt.Println("Retrieving edit links 1")
	resp, err := client.Get("http://www.sena.lt/skelbimai")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	root, err := xmlpath.ParseHTML(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// generate links of all pages of items
	pageLinks := getPageLinks(root)

	// extract editing links from the first page
	editLinks := getEditLinks(root)

	// extract editing links from the rest of the pages
	for i, link := range pageLinks {
		fmt.Printf("Retrieving edit links %d of %d\n", i+2, len(pageLinks)+1)
		pageResp, err := client.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer pageResp.Body.Close()
		pageRoot, err := xmlpath.ParseHTML(pageResp.Body)
		if err != nil {
			log.Fatal(err)
		}
		pageEditLinks := getEditLinks(pageRoot)
		editLinks = append(editLinks, pageEditLinks...)
	}

	// reverse, newest items need to be renewed last
	reverse(editLinks)
	for i, link := range editLinks {
		fmt.Printf("Renewing %d of %d\n", i+1, len(editLinks))
		adResp, err := client.Get(link)
		if err != nil {
			fmt.Println(err)
		}
		defer adResp.Body.Close()
		adRoot, err := xmlpath.ParseHTML(adResp.Body)
		if err != nil {
			log.Fatal(err)
		}
		adData, err := parseAdPage(adRoot)
		if err != nil {
			log.Fatal(err)
		}
		payload, boundary := generateRequestPayload(adData)
		req, err := http.NewRequest("POST", "http://www.sena.lt/naujas_skelbimas", strings.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Origin", "http://www.sena.lt")
		req.Header.Set("Host", "www.sena.lt")
		req.Header.Set("Referer", link)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:31.0) Gecko/20100101 Firefox/31.0")
		req.AddCookie(cookieJar.Cookies(turl)[0])
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		time.Sleep(30 * time.Second)
	}
}
