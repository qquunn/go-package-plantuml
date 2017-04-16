package main

import (
	uml "git.oschina.net/jscode/go-package-plantuml/testdata/uml"
	sub "git.oschina.net/jscode/go-package-plantuml/testdata/uml/sub"
	sub2 "git.oschina.net/jscode/go-package-plantuml/testdata/uml/sub2"
	"fmt"
)


func main() {
	fmt.Println(uml.SA{})
	fmt.Println(sub.SA{})
	fmt.Println(sub2.Sub2A{})
}

