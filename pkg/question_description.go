package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// GetQuestionDetailByTitleSlug gets question contents by title slug.
// Return empty string if failed.
func GetQuestionDetailByTitleSlug(titleSlug string, language string) string {
	detailMap := getQuestionData(titleSlug)
	return getQuestionContentFromQuestionMap(detailMap, language)
}

func getQuestionContentFromQuestionMap(detailMap map[string]string, language string) string {
	if len(detailMap) == 0 {
		fmt.Println("Empty detail map.")
		return ""
	}
	if language == "zh" {
		return detailMap["translatedContent"]
	} else {
		return detailMap["content"]
	}
}

func GetQuestionDetailByFrontendId(value string) string {
	// TODO
	return ""
}

func getQuestionData(titleSlug string) map[string]string {
	url := "https://leetcode-cn.com/graphql/"
	method := "POST"

	//binary-subarrays-with-sum
	payload := strings.NewReader(`{"operationName":"questionData","variables":{"titleSlug":"` + titleSlug + `"},"query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    categoryTitle\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    exampleTestcases\n    __typename\n  }\n}\n"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "csrftoken=TcsWO3MAaL0FZ0NBKt4kkGp1N3HokGbVbhhCkgNpEjiXGqLJAeqVgpZ28WJdILo2")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	detailMap := extractQuesitionDataToMap(string(body))
	return detailMap
}

func extractQuesitionDataToMap(result string) map[string]string {
	//fmt.Println(result)
	filedList := []string{
		"questionId",
		"questionFrontendId",
		"categoryTitle",
		"boundTopicId",
		"title",
		"translatedTitle",
		"titleSlug",
		"content",
		"translatedContent",
		"difficulty",
		"likes",
		"dislikes",
		"topicTags",
		"codeSnippets",
	}
	detailMap := make(map[string]string)
	for _, fieldName := range filedList {
		field := gjson.Get(result, "data.question."+fieldName)
		if fieldName == "topicTags" {
			field = gjson.Get(field.String(), "#.name")
		}
		//detailMap = append(detailMap, []string{fieldName, field.String()})
		detailMap[fieldName] = field.String()
	}
	return detailMap
}

func GetPrettyText(html string) string {
	return Html2BashText(html)
}
