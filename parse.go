package main

import (
	"fmt"
	"launchpad.net/xmlpath"
	"math"
	"regexp"
	"strconv"
)

func getPageLinks(root *xmlpath.Node) []string {
	path := xmlpath.MustCompile(`//*[@id="login"]/ul[1]/li[3]/span`)
	value, ok := path.String(root)
	r, _ := regexp.Compile("[0-9]+")
	match := r.FindString(value)
	numBooks, _ := strconv.ParseFloat(match, 64)
	numPages := int(math.Ceil(numBooks / 50))
	pageLinks := make([]string, numPages-1)
	for i := 1; i < numPages; i++ {
		pageLinks[i-1] = fmt.Sprintf("http://www.sena.lt/skelbimai/%s/p/%d", username, i)
	}
	return pageLinks
}

func getEditLinks(root *xmlpath.Node) []string {
	path := xmlpath.MustCompile(`//td[@class="info"]//a[3]/@href`)
	iterator := path.Iter(root)
	editLinks := make([]string, 0, 50)
	for iterator.Next() {
		editLinks = append(editLinks, fmt.Sprintf("http://www.sena.lt%s", iterator.Node().String()))
	}
	return editLinks
}

func parseAdPage(root *xmlpath.Node) {

}
