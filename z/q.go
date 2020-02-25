package z

import "github.com/AlecAivazis/survey/v2"

type Answer struct {
	Project string
	Author  string
	Workdir string
}

const (
	default_project = "app"
	default_author  = "user"
)

func NewAnswer() *Answer {
	return &Answer{
		Project: default_project,
		Author:  default_author,
		Workdir: "",
	}
}

func NewQ() []*survey.Question {
	return []*survey.Question{
		{
			Name: "Project",
			Prompt: &survey.Input{
				Message: "工程名: ",
				Default: default_project,
			},
			Validate: survey.Required,
		},
		{
			Name: "Author",
			Prompt: &survey.Input{
				Message: "作者名: ",
				Default: default_author,
			},
			Validate: survey.Required,
		},
	}
}
