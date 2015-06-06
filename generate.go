package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func generateRequestPayload(adData map[string]string) (payload, boundary string) {
	const payloadStub = `
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[type]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[item_action]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[ajax_author]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[author]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[author_journal]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[author_last_name]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[author_name]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[item]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[journal]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[title]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[item_category_id]"

1
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[published]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[num]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[language_id]"

1
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[summary]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[comment]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[tags]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[photo]"; filename=""
Content-Type: application/octet-stream


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[item_condition_id]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[price]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[currency]"

3
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[used_comment]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[valid_for]"

2

------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[agree]"

on
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[id]"

%s
------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[webkey]"


------WebKitFormBoundary<rndStr>
Content-Disposition: form-data; name="used_item_create[user_id]"

%s
------WebKitFormBoundary<rndStr>--`
	rndStr := randomString(16)
	payload = fmt.Sprintf(strings.Replace(payloadStub, "<rndStr>", rndStr, -1), adData["itemType"], adData["itemAction"], adData["author"], adData["authorJournal"], adData["createItem"], adData["createJournal"], adData["itemCondition"], adData["price"], adData["comment"], adData["itemId"], adData["userId"])
    boundary = "------WebKitFormBoundary" + rndStr
    return 
}
