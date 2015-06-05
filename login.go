package main

import (
//    "os"
    "net/url"
	"net/http"
    "net/http/cookiejar"
    "fmt"
    "launchpad.net/xmlpath"
)

func main() {

    cookieJar, _ := cookiejar.New(nil)
    client := &http.Client{
        Jar: cookieJar,
    }

    loginValues := url.Values{}
    loginValues.Set("login[username]", username)
    loginValues.Set("login[password]", password)

    loginResp, _ := client.PostForm("http://www.sena.lt/user/login", loginValues)
    defer loginResp.Body.Close()


	resp, _ := client.Get("http://www.sena.lt/skelbimai")
    defer resp.Body.Close()
    root, _ := xmlpath.ParseHTML(resp.Body)
//    file, _ := os.Open("test.html")
//    defer file.Close()
    pageLinks := getPageLinks(root)
	fmt.Println(pageLinks)
    editLinks := getEditLinks(root)
    fmt.Println(len(editLinks))
}
