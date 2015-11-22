package main

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
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

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	loginValues := url.Values{}
	loginValues.Set("login[username]", username)
	loginValues.Set("login[password]", password)

	// login
	fmt.Println("Logging in...")
	loginResp, loginErr := client.PostForm("http://www.sena.lt/user/login", loginValues)
	if loginErr != nil {
		fmt.Println(loginErr)
	}
	defer loginResp.Body.Close()

	turl, _ := url.Parse("http://www.sena.lt")
	fmt.Println(cookieJar.Cookies(turl))

	// retrieve the first page of ads
	fmt.Println("Retrieving edit links 1")
	resp, mainErr := client.Get("http://www.sena.lt/skelbimai")
	if mainErr != nil {
		fmt.Println(mainErr)
	}
	defer resp.Body.Close()

	root, _ := xmlpath.ParseHTML(resp.Body)

	// generate links of all pages of items
	pageLinks := getPageLinks(root)

	// extract editing links from the first page
	editLinks := getEditLinks(root)

	// extract editing links from the rest of the pages
	for i, link := range pageLinks {
		fmt.Printf("Retrieving edit links %d of %d\n", i+2, len(pageLinks)+1)
		pageResp, err := client.Get(link)
		if err != nil {
			fmt.Println(err)
		}
		defer pageResp.Body.Close()
		pageRoot, _ := xmlpath.ParseHTML(pageResp.Body)
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
		adRoot, _ := xmlpath.ParseHTML(adResp.Body)
		adData, _ := parseAdPage(adRoot)
		payload, boundary := generateRequestPayload(adData)
		req, _ := http.NewRequest("POST", "http://www.sena.lt/naujas_skelbimas", strings.NewReader(payload))
		req.Header.Set("Origin", "http://www.sena.lt")
		req.Header.Set("Host", "www.sena.lt")
		req.Header.Set("Referer", link)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; rv:31.0) Gecko/20100101 Firefox/31.0")
		req.AddCookie(cookieJar.Cookies(turl)[0])
		res, _ := client.Do(req)
		defer res.Body.Close()
		time.Sleep(30 * time.Second)
		//      contents, _ := ioutil.ReadAll(res.Body)
	}
}
