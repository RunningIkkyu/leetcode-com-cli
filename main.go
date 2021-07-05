package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/tidwall/gjson"
)

func main() {
	ShowTodayQuestion()
}

func ShowTodayQuestion() {

	url := "https://leetcode-cn.com/graphql/"
	method := "GET"

	payload := strings.NewReader(`{
    "query": "\n    query questionOfToday {\n  todayRecord {\n    date\n    userStatus\n    question {\n      questionId\n      frontendQuestionId: questionFrontendId\n      difficulty\n      title\n      titleCn: translatedTitle\n      titleSlug\n      paidOnly: isPaidOnly\n      freqBar\n      isFavor\n      acRate\n      status\n      solutionNum\n      hasVideoSolution\n      topicTags {\n        name\n        nameTranslated: translatedName\n        id\n      }\n      extra {\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n    lastSubmission {\n      id\n    }\n  }\n}\n    ",
    "variables": {}
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
    req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15")
	//req.Header.Add("Cookie", "LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiOTk2NzQ2IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYXV0aGVudGljYXRpb24uYXV0aF9iYWNrZW5kcy5QaG9uZUF1dGhlbnRpY2F0aW9uQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjEzNTBhNjYwMTFlZGM5NDA5MGEwMGVlMzc0ZGY5ODA5OWM2N2ViMzZmMGM3NDUyZDg3MTQ3MjRiZGM3OTRiNWQiLCJpZCI6OTk2NzQ2LCJlbWFpbCI6IiIsInVzZXJuYW1lIjoiaWtreXUtMiIsInVzZXJfc2x1ZyI6Imlra3l1LTIiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC9kZWZhdWx0X2F2YXRhci5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTYyNTQ5ODY2Mi4yMDQ3MTksImV4cGlyZWRfdGltZV8iOjE2MjgwMTcyMDAsImxhdGVzdF90aW1lc3RhbXBfIjoxNjI1NTAwNTQwfQ.bF5vzyd4GnRxkXam-zI4U6h-OgJSM9WjyUBaYrEaKs4; csrftoken=zMjGSSRcSEj4CRUD2W4Swh2262p0Mjtm8AUcKKoWMZjMci9MOY9SRqu1av5npCRT; __asc=07f7112417a77450ceb882b616c; __auc=07f7112417a77450ceb882b616c; a2873925c34ecbd2_gr_cs1=ikkyu-2; a2873925c34ecbd2_gr_last_sent_cs1=ikkyu-2; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=3d64d059-530c-4c4a-a403-7301ba994303; a2873925c34ecbd2_gr_session_id=3d64d059-530c-4c4a-a403-7301ba994303; a2873925c34ecbd2_gr_session_id_3d64d059-530c-4c4a-a403-7301ba994303=true; gr_user_id=94a94f08-455d-41fc-a490-93a061c3767c; NEW_PROBLEMLIST_PAGE=1; _ga=GA1.2.907764100.1625498655; _gid=GA1.2.1539144914.1625498655; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1625500334;  Hm_lvt_fa218a3ff7179639febdb15e372f411c=1625498652; csrftoken=zMjGSSRcSEj4CRUD2W4Swh2262p0Mjtm8AUcKKoWMZjMci9MOY9SRqu1av5npCRT")
	//req.Header.Add("x-csrftoken", "zMjGSSRcSEj4CRUD2W4Swh2262p0Mjtm8AUcKKoWMZjMci9MOY9SRqu1av5npCRT")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(body)
	printMap(result)
}

// printMap prints a [][]string in table format
func printMap(result string) {
	questionMap := extractQuesitionInfoFromBody(result)
	t := table.NewWriter()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
    title := "Daily Question"
	t.AppendHeader(table.Row{title, title}, rowConfigAutoMerge)
	for _, j := range questionMap {
		t.AppendRow([]interface{}{j[0], j[1]})
	}
	fmt.Println(t.Render())

}

// extractQuesitionInfoFromBody extracts info of today question into a slice.
// Example: 
//
//   []string{
//       {"item1", "value1"}
//       {"item2", "value2"}
//       {"item3", "value3"}
//       ...
//   }
func extractQuesitionInfoFromBody(result string) [][]string {
	//fmt.Println(result)
	filedList := []string{
		"questionId",
		"frontendQuestionId",
		"difficulty",
		"title",
		"titleCn",
		"titleSlug",
		"acRate",
		"solutionNum",
		"topicTags",
	}
	questionMap := [][]string{}
	for _, fieldName := range filedList {
		field := gjson.Get(result, "data.todayRecord.0.question."+fieldName)
		if fieldName == "topicTags" {
			field = gjson.Get(field.String(), "#.name")
		}
        questionMap = append(questionMap, []string{fieldName, field.String()})
	}
	return questionMap
}
