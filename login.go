package main

import (
	"fmt"
	"launchpad.net/xmlpath"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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
	loginResp, _ := client.PostForm("http://www.sena.lt/user/login", loginValues)
	defer loginResp.Body.Close()

	// retrieve the first page of ads
	resp, _ := client.Get("http://www.sena.lt/skelbimai")
	defer resp.Body.Close()

	root, _ := xmlpath.ParseHTML(resp.Body)

	// generate links of all pages of items
	pageLinks := getPageLinks(root)

	// extract editing links from the first page
	editLinks := getEditLinks(root)

	// extract editing links from the rest of the pages
	for _, link := range pageLinks {
		pageResp, _ := client.Get(link)
		defer pageResp.Body.Close()

		pageRoot, _ := xmlpath.ParseHTML(pageResp.Body)
		pageEditLinks := getEditLinks(pageRoot)
		editLinks = append(editLinks, pageEditLinks...)
	}

	// reverse, newest items need to be renewed last
	reverse(editLinks)

	for _, link := range editLinks {
		fmt.Println(link)
	}
}
