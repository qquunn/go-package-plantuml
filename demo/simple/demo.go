package main

import (
	"git.oschina.net/jscode/go-package-plantuml/codeanalysis"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"io/ioutil"
)

func main() {

	log.SetLevel(log.InfoLevel)

	//config := codeanalysis.Config{
	//	CodeDir: "/appdev/go-demo/src/git.oschina.net/jscode/go-package-plantuml/testdata/a",
	//	GopathDir : "/appdev/go-demo",
	//}
	//
	//result := codeanalysis.AnalysisCode(config)
	//
	//result.OutputToFile("/tmp/uml.txt")

	config := codeanalysis.Config{
		CodeDir: "/appdev/go-demo/src/git.oschina.net/jscode/go-package-plantuml/testdata/uml",
		GopathDir : "/appdev/go-demo",
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile("/tmp/uml.txt")

	bytes, _ := ioutil.ReadFile("/tmp/uml.txt")

	fmt.Println(string(bytes))

	// java -jar /app/plantuml.jar  /tmp/uml.txt -tsvg && open2 /tmp/uml.svg

}
