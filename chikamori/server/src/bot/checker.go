package bot

import (
	"regexp"

	"model"
)

type (
	// Checker
	// messageを受け，条件を満たすか判定するためのinterface
	Checker interface {
		Check(model.Message) bool
	}

	// ThroughChecker
	// 常にtrueを返すchecker
	ThroughChecker struct {
	}

	// RegExpChecker
	// 与えられた正規表現を満たす場合trueを返すChecker
	RegexpChecker struct {
		Pattern string
	}
)

func (c *ThroughChecker) Check(m model.Message) bool {
	return true
}

func (c *RegexpChecker) Check(m model.Message) bool {
	r := regexp.MustCompile(c.Pattern)
	return r.MatchString(m.Body)
}
