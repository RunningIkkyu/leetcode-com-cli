package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/tidwall/gjson"
)

const (
	QUESTION_DATA_FIELD_TITLE_SLUG           = "titleSlug"
	QUESTION_DATA_FIELD_TITLE_CN             = "translatedTitle"
	QUESTION_DATA_FIELD_TITLE                = "title"
	QUESTION_DATA_FIELD_QUESTION_ID          = "questionId"
	QUESTION_DATA_FIELD_FRONTEND_QUESTION_ID = "frontendQuestionId"
	QUESTION_DATA_FIELD_DIFFICULTY           = "difficulty"
	QUESTION_DATA_FIELD_SOLUTION_NUM         = "solutionNum"
	QUESTION_DATA_FIELD_STATUS               = "status"
	QUESTION_DATA_FIELD_ISPAIDONLY           = "isPaidOnly"
	QUESTION_DATA_FIELD_TOPICTAGS            = "topicTags"
	QUESTION_DATA_FIELD_ACRATE               = "acRate"
)

func GetTodayQuestionInfo() map[string]string {
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
		return nil
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15")
	//req.Header.Add("x-csrftoken", "zMjGSSRcSEj4CRUD2W4Swh2262p0Mjtm8AUcKKoWMZjMci9MOY9SRqu1av5npCRT")
	req.Header.Add("Content-Type", "application/json")

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
	result := string(body)
	questionMap := extractQuesitionInfoFromBody(result)
	return questionMap
}

func ShowTodayQuestion() {
	m := GetTodayQuestionInfo()
	PrintMap(m)
}

func PrintMap(questionMap map[string]string) {
	t := table.NewWriter()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	title := "Daily Question"
	t.AppendHeader(table.Row{title, title}, rowConfigAutoMerge)
	for i, j := range questionMap {
		t.AppendRow([]interface{}{i, j})
	}
	fmt.Println(t.Render())
}

func StandardedLanguage(language string) string {
	switch language {
	case "zh":
		return "zh"
	case "en":
		return "en"
	default:
		return "en"
	}
}

// PrettyPrintQuestionByTitleSlug pretty print question by given titleSlug with the specific language.
func PrettyPrintQuestionByTitleSlug(titleSlug string, showLanguage string) {
	questionData := getQuestionData(titleSlug)
	if len(questionData) == 0 {
		fmt.Printf("Question %s no exists", titleSlug)
		return
	}

	PrintQuestionTitle(questionData, showLanguage)
	content := getQuestionContentFromQuestionMap(questionData, showLanguage)
	PrettyPrintContent(content)
}

func PrettyPrintContent(content string) {
	prettyDetail := GetPrettyText(content)
	fmt.Println(prettyDetail)
}

// PrintQuestionTitle pretty pring title and subtitle by given languane.
func PrintQuestionTitle(questionMap map[string]string, language string) {
	var titleStr, difficultyLeft string
	switch language {
	case "zh":
		titleStr = questionMap[QUESTION_DATA_FIELD_TITLE_CN]
		//acRateLeft = "通过率："
		//solutionNumLeft = "题解数："
		difficultyLeft = "难度："
	case "en":
		titleStr = questionMap[QUESTION_DATA_FIELD_TITLE]
		//acRateLeft = "AC Rate: "
		//solutionNumLeft = "Solution Num: "
		difficultyLeft = "Difficulty: "
	}
	title := BoldText(questionMap["questionId"] + "." + titleStr)
	subtitleList := []string{
		//acRateLeft + questionMap[QUESTION_DATA_FIELD_ACRATE],
		//solutionNumLeft + questionMap[QUESTION_DATA_FIELD_SOLUTION_NUM],
		difficultyLeft + questionMap[QUESTION_DATA_FIELD_DIFFICULTY],
	}
	subtitle := ItalicText(strings.Join(subtitleList, "\t"))
	fmt.Println(title)
	fmt.Println(subtitle)
	fmt.Println()
}

// extractQuesitionInfoFromBody extracts info of today question into a map.
func extractQuesitionInfoFromBody(result string) map[string]string {
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
	questionMap := make(map[string]string)
	for _, fieldName := range filedList {
		field := gjson.Get(result, "data.todayRecord.0.question."+fieldName)
		if fieldName == "topicTags" {
			field = gjson.Get(field.String(), "#.name")
		}
		//questionMap = append(questionMap, []string{fieldName, field.String()})
		questionMap[fieldName] = field.String()
	}
	return questionMap
}
