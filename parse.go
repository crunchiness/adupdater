package main

import (
	"errors"
	"fmt"
	"gopkg.in/xmlpath.v2"
	"math"
	"regexp"
	"strconv"
)

func getPageLinks(root *xmlpath.Node) []string {
	path := xmlpath.MustCompile(`//*[@id="login"]/ul[1]/li[3]/span`)
	value, _ := path.String(root)
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

func parseAdPage(root *xmlpath.Node) (map[string]string, error) {
	const errorStub = "Failed to parse %s"

	adData := make(map[string]string)

	itemTypeX := `//form[@action="/naujas_skelbimas"]//input[(@name="used_item_create[type]") and (@checked)]/@value`
	path := xmlpath.MustCompile(itemTypeX)
	value, ok := path.String(root)
	if ok {
		adData["itemType"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "itemType"))
	}

	itemActionX := `//form[@action="/naujas_skelbimas"]//input[(@name="used_item_create[item_action]") and (@checked)]/@value`
	path = xmlpath.MustCompile(itemActionX)
	value, ok = path.String(root)
	if ok {
		adData["itemAction"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "itemAction"))
	}

	authorX := `//form[@action="/naujas_skelbimas"]//select[@name="used_item_create[author]"]/option[@selected]/@value`
	path = xmlpath.MustCompile(authorX)
	value, ok = path.String(root)
	if ok {
		adData["author"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "author"))
	}

	authorJournalX := `//form[@action="/naujas_skelbimas"]//select[@name="used_item_create[author_journal]"]/option[@selected]/@value`
	path = xmlpath.MustCompile(authorJournalX)
	value, ok = path.String(root)
	if ok {
		adData["authorJournal"] = value
	} else {
		adData["authorJournal"] = "0"
	}

	createItemX := `//form[@action="/naujas_skelbimas"]//select[@name="used_item_create[item]"]/option[@selected]/@value`
	path = xmlpath.MustCompile(createItemX)
	value, ok = path.String(root)
	if ok {
		adData["createItem"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "createItem"))
	}

	createJournalX := `//form[@action="/naujas_skelbimas"]//select[@name="used_item_create[journal]"]/option[@selected]/@value`
	path = xmlpath.MustCompile(createJournalX)
	value, ok = path.String(root)
	if ok {
		adData["createJournal"] = value
	} else {
		adData["createJournal"] = "0"
	}

	photoX := `//form[@action="/naujas_skelbimas"]//input[(@name="photo_radio") and (@checked)]/@value`
	path = xmlpath.MustCompile(photoX)
	value, ok = path.String(root)
	if ok {
		adData["photo"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "photo"))
	}

	itemConditionX := `//form[@action="/naujas_skelbimas"]//select[@name="used_item_create[item_condition_id]"]/option[@selected]/@value`
	path = xmlpath.MustCompile(itemConditionX)
	value, ok = path.String(root)
	if ok {
		adData["itemCondition"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "itemCondition"))
	}

	priceX := `//form[@action="/naujas_skelbimas"]//input[@name="used_item_create[price]"]/@value`
	path = xmlpath.MustCompile(priceX)
	value, ok = path.String(root)
	if ok {
		adData["price"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "price"))
	}

	commentX := `//form[@action="/naujas_skelbimas"]//textarea[@name="used_item_create[used_comment]"]`
	path = xmlpath.MustCompile(commentX)
	value, ok = path.String(root)
	if ok {
		adData["comment"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "comment"))
	}

	itemIdX := `//form[@action="/naujas_skelbimas"]//input[@name="used_item_create[id]"]/@value`
	path = xmlpath.MustCompile(itemIdX)
	value, ok = path.String(root)
	if ok {
		adData["itemId"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "itemId"))
	}

	userIdX := `//form[@action="/naujas_skelbimas"]//input[@name="used_item_create[user_id]"]/@value`
	path = xmlpath.MustCompile(userIdX)
	value, ok = path.String(root)
	if ok {
		adData["userId"] = value
	} else {
		return nil, errors.New(fmt.Sprintf(errorStub, "userId"))
	}

	return adData, nil
}
