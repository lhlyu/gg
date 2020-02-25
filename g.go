package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/lhlyu/gg/z"
	"github.com/lhlyu/logger/v3/color"
)

const tip = `
----------------------------------------------------------------------------------------
|                                                                                      
|      version: 1.0.0                                                                                   
|
|      go建议1.13+
----------------------------------------------------------------------------------------
`

const success = `
----------------------------------------------------------------------------------------
|                                                                                       
|     SUCCESS !!!
|
|     cd %s
|     go run main.go
|     ----------------  工具推荐 --------------------------------------
|     mysql table to go struct : https://github.com/lhlyu/got/releases
|     json to go struct        : http://lhlyu.gitee.io/json-to-go/ 
|
----------------------------------------------------------------------------------------
`

func main() {
	//createTemplate()
	cli()
}

func createTemplate() {
	// 制作模板
	c := z.NewC()
	c.SetDir(`F:\projects\src\alpha\beta`)
	c.AddExclude(".idea", "deploy_tmp", "beta")
	c.AddMSS("alpha/", "{{.Workdir}}")
	c.AddMSS("beta", "{{.Project}}")
	c.AddMSS("gama", "{{.Author}}")

	c.Create()
}

// Cli
func cli() {
	fmt.Println(color.Cyan(tip))
	qs := z.NewQ()
	an := z.NewAnswer()
	if err := survey.Ask(qs, an); err != nil {
		panic(err)
		return
	}
	node := &z.Node{}
	err := json.Unmarshal([]byte(z.T), node)
	if err != nil {
		panic(err)
	}
	node.Build(an)
	fmt.Println(color.Greenf(success, an.Project))
}
