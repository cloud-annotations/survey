package main

import (
	"github.com/cloud-annotations/survey"
	"github.com/cloud-annotations/survey/tests/util"
)

var answer = false

var goodTable = []TestUtil.TestTableEntry{
	{
		"Enter 'yes'", &survey.Confirm{
			Message: "yes:",
		}, &answer, nil,
	},
	{
		"Enter 'no'", &survey.Confirm{
			Message: "yes:",
		}, &answer, nil,
	},
	{
		"default", &survey.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
	{
		"not recognized (enter random letter)", &survey.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
	{
		"no help - type '?'", &survey.Confirm{
			Message: "yes:",
			Default: true,
		}, &answer, nil,
	},
}

func main() {
	TestUtil.RunTable(goodTable)
}
