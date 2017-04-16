package main

import (
	"git.oschina.net/jscode/go-package-plantuml/codeanalysis"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
	"path"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir string `long:"codedir" description:"要扫描的代码目录" required:"true"`
		GopathDir string `long:"gopath" description:"GOPATH目录" required:"true"`
		OutputFile string `long:"outputfile" description:"解析结果保存到该文件中" required:"true"`
		IgnoreDirs []string `long:"ignoredir" description:"需要排除的目录,不需要扫描和解析"`
	}

	if len(os.Args) == 1 {
		fmt.Println("使用例子\n" +
			os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --gopath /appdev/gopath --outputfile  /tmp/result")
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	if opts.CodeDir == "" {
		panic("代码目录不能为空")
		os.Exit(1)
	}

	if opts.GopathDir == "" {
		panic("GOPATH目录不能为空")
		os.Exit(1)
	}

	if ! strings.HasPrefix(opts.CodeDir, opts.GopathDir) {
		panic(fmt.Sprintf("代码目录%s,必须是GOPATH目录%s的子目录", opts.CodeDir, opts.GopathDir))
		os.Exit(1)
	}

	for _, dir := range opts.IgnoreDirs {
		if ! strings.HasPrefix(dir, opts.CodeDir){
			panic(fmt.Sprintf("需要排除的目录%s,必须是代码目录%s的子目录", dir, opts.CodeDir))
			os.Exit(1)
		}
	}

	config := codeanalysis.Config{
		CodeDir: opts.CodeDir,
		GopathDir :opts.GopathDir,
		VendorDir : path.Join(opts.CodeDir, "vendor"),
		IgnoreDirs:opts.IgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile("/tmp/uml.txt")

}

