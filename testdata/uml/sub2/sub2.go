package sub2

import sub "git.oschina.net/jscode/go-package-plantuml/testdata/uml/sub"

type Sub2I interface {
	Add(d sub.SA)
}

type Sub2A struct {
	a AliasA
}

type AliasA string

