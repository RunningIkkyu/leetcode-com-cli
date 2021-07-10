package pkg_test

import (
	"fmt"
	"testing"

	"github.com/RunningIkkyu/leetcode-com-cli/pkg"
)

func TestShowTodayQuestion(t *testing.T){
    m := pkg.GetTodayQuestionInfo()
    pkg.PrintMap(m)
    titleSlug := m["titleSlug"]
    detail := pkg.GetQuestionDetail(titleSlug)
    prettyDetail := pkg.GetPrettyText(detail)
    fmt.Println(prettyDetail)
}
