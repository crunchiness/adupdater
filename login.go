package adupdater

import (
	"fmt"
    "net/url"
	"net/http"
    "net/http/cookiejar"
    "io/ioutil"
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

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
