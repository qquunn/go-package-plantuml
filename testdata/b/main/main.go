package main

import (
	"git.oschina.net/jscode/go-package-plantuml/testdata/b"
	"git.oschina.net/jscode/go-package-plantuml/testdata/b/sub"
)

func m1(a1 b.IA) {
	a1.Add(sub2.SubSA{})
}

func main() {
	m1(&b.SB{})
}

