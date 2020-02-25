package z

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"
)

type MSB = map[string]bool
type MSS = map[string]string
type MSI = map[string]interface{}

type Node struct {
	Name    string  `json:"name,omitempty"`
	Dir     string  `json:"dir,omitempty"`
	IsDir   bool    `json:"isDir"`
	Content string  `json:"content,omitempty"`
	Childs  []*Node `json:"childs,omitempty"`
}

func NewNode(name, dir string, isDir bool) *Node {
	dir = strings.Trim(dir, "/")
	return &Node{
		Name:  name,
		Dir:   dir,
		IsDir: isDir,
	}
}

func (this *Node) SetContent(content string) {
	this.Content = base64.StdEncoding.EncodeToString([]byte(content))
}

func (this *Node) AddNodes(childs ...*Node) {
	this.Childs = append(this.Childs, childs...)
}

func (this *Node) Write(name string) {
	if name == "" {
		name = "./z/t.go"
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("package z\n\n")
	bts, _ := json.MarshalIndent(this, "", "	")
	f.WriteString(fmt.Sprintf("const T = `%s`", bts))
}

func (this *Node) Build(an *Answer) {
	root := "./"
	dir := path.Join(root, an.Project)
	err := os.Mkdir(dir, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", an)
	this.build(this.Childs, an, dir)
}

func (this *Node) build(childs []*Node, an *Answer, dir string) {
	for _, v := range childs {
		if v.IsDir {
			this.build(v.Childs, an, dir)
		} else {
			fpath := path.Join(MakeDirs(dir, v.Dir), v.Name)
			f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}
			bts, err := base64.StdEncoding.DecodeString(v.Content)
			if err != nil {
				panic(err)
			}
			tpl := template.Must(template.New("template.html").Parse(string(bts)))
			buf := bytes.NewBufferString("")
			err = tpl.Execute(buf, an)
			if err != nil {
				panic(err)
			}
			f.WriteString(buf.String())
			f.Close()
		}
	}
}
