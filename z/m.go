// 模板制作
package z

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
)

type C struct {
	length   int
	dir      string
	excludes MSB
	mss      MSS
}

func NewC() *C {
	return &C{}
}

func (this *C) SetDir(dir string) *C {
	dir = strings.ReplaceAll(dir, `\\`, "/")
	dir = strings.ReplaceAll(dir, `\`, "/")
	this.length = len(dir)
	this.dir = dir
	return this
}

func (this *C) SetExcludes(excludes MSB) *C {
	this.excludes = excludes
	return this
}

func (this *C) AddExclude(excludes ...string) *C {
	if this.excludes == nil {
		this.excludes = make(MSB)
	}
	for _, v := range excludes {
		this.excludes[v] = true
	}
	return this
}

func (this *C) SetMSS(mss MSS) *C {
	this.mss = mss
	return this
}

func (this *C) AddMSS(key, value string) *C {
	if this.mss == nil {
		this.mss = make(MSS)
	}
	this.mss[key] = value
	return this
}

func (this *C) Create() {
	node := NewNode("gg", "/", true)
	this.walk(this.dir, node)
	node.Write("")
}

func (this *C) walk(dir string, node *Node) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range files {
		// 过滤
		if this.excludes[v.Name()] {
			continue
		}
		if v.IsDir() {
			next := NewNode(v.Name(), dir[this.length:], true)
			this.walk(path.Clean(path.Join(dir, v.Name())), next)
			node.AddNodes(next)
		} else {
			bts, err := ioutil.ReadFile(path.Join(dir, v.Name()))
			if err != nil {
				panic(err)
			}
			content := string(bts)
			for k, v := range this.mss {
				content = strings.ReplaceAll(content, k, v)
			}
			child := NewNode(v.Name(), dir[this.length:], false)
			child.SetContent(content)
			node.AddNodes(child)
		}
	}
}
