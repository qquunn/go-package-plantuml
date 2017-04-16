package codeanalysis

import (
	"testing"
	"github.com/stvp/assert"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)


var gopathDir = os.Getenv("GOPATH")
var testdataPath = gopathDir + "/src/git.oschina.net/jscode/go-package-plantuml/testdata"

func Test_findGoPackageNameInDirPath(t *testing.T) {
	assert.Equal(t, "b", findGoPackageNameInDirPath(testdataPath + "/b"))
	assert.Equal(t, "sub2", findGoPackageNameInDirPath(testdataPath + "/b/sub"))
}

func Test_InterfacesSign(t *testing.T) {

	config := Config{
		CodeDir: testdataPath + "/a",
		GopathDir :gopathDir,
		IgnoreDirs:[]string{},
	}

	result := AnalysisCode(config)

	analysisTool1, _ := result.(*analysisTool)

	assert.Equal(t, 1, len(analysisTool1.interfaceMetas))

	interfaceMeta := analysisTool1.interfaceMetas[0]
	assert.Equal(t, "IA", interfaceMeta.Name)
	assert.Equal(t, testdataPath + "/a/a.go", interfaceMeta.FilePath)
	assert.Equal(t, "git.oschina.net/jscode/go-package-plantuml/testdata/a", interfaceMeta.PackagePath, "error in interfaceMeta")

	fmt.Println(interfaceMeta.MethodSigns)
	assert.Equal(t, 4, len(interfaceMeta.MethodSigns))
	assert.Equal(t, "Add()", interfaceMeta.MethodSigns[0])
	assert.Equal(t, "Add2(int)int", interfaceMeta.MethodSigns[1])
	assert.Equal(t, "Add3(int,int,int)(int,int)", interfaceMeta.MethodSigns[2])
	assert.Equal(t, "Add4(int)int", interfaceMeta.MethodSigns[3])

	assert.Equal(t, 1, len(analysisTool1.structMetas))

	structmeta := analysisTool1.structMetas[0]
	assert.Equal(t, "SA", structmeta.Name)
	assert.Equal(t, testdataPath + "/a/a.go", structmeta.FilePath)
	assert.Equal(t, "git.oschina.net/jscode/go-package-plantuml/testdata/a", structmeta.PackagePath, "error in structmeta")

	fmt.Println(structmeta.MethodSigns)
	assert.Equal(t, 4, len(structmeta.MethodSigns))
	assert.Equal(t, "Add()", structmeta.MethodSigns[0])
	assert.Equal(t, "Add2(int)int", structmeta.MethodSigns[1])
	assert.Equal(t, "Add3(int,int,int)(int,int)", structmeta.MethodSigns[2])
	assert.Equal(t, "Add4(int)int", structmeta.MethodSigns[3])

	structMetas := analysisTool1.findInterfaceImpls(interfaceMeta)
	assert.Equal(t, 1, len(structMetas))
	assert.Equal(t, "SA", structMetas[0].Name)

}

/**
 * 测试Struct的方法声明不使用指针
 * 测试Struct使用路径引用; 路径引用使用别名; 标准库引用;
 * 测试import . 包路径
 */
func Test_complex(t *testing.T) {

	log.SetLevel(log.WarnLevel)

	config := Config{
		CodeDir: testdataPath + "/b",
		GopathDir :gopathDir,
		IgnoreDirs:[]string{},
	}

	result := AnalysisCode(config)

	analysisTool1, _ := result.(*analysisTool)

	assert.Equal(t, 1, len(analysisTool1.interfaceMetas))
	interfaceMeta := analysisTool1.interfaceMetas[0]
	assert.Equal(t, "IA", interfaceMeta.Name)
	assert.Equal(t, "Add(git.oschina.net/jscode/go-package-plantuml/testdata/b/sub.SubSA,sync.Locker,git.oschina.net/jscode/go-package-plantuml/testdata/b.B,git.oschina.net/jscode/go-package-plantuml/testdata/b/suba.SubSa1)", interfaceMeta.MethodSigns[0])

	assert.Equal(t, 4, len(analysisTool1.structMetas))
	structMetas := analysisTool1.findInterfaceImpls(interfaceMeta)
	assert.Equal(t, 1, len(structMetas))
	assert.Equal(t, "SB", structMetas[0].Name)

}


func Test_uml(t *testing.T) {

	config := Config{
		CodeDir: testdataPath + "/uml",
		GopathDir :gopathDir,
		IgnoreDirs:[]string{},
	}

	result := AnalysisCode(config)

	analysisTool1, _ := result.(*analysisTool)

	fmt.Println(analysisTool1.UML())

	assert.Equal(t, 3, len(analysisTool1.interfaceMetas))
	interfaceMeta := analysisTool1.interfaceMetas[0]
	assert.Equal(t, "namespace git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml {\n interface IA  {\n  Add()\n} \n}", interfaceMeta.UML)

	assert.Equal(t, 3, len(analysisTool1.structMetas))
	structMeta1 := analysisTool1.structMetas[0]
	assert.Equal(t, "namespace git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml {\n class SA {\n  a int\n  b sync.Mutex\n  c sub2.Sub2A\n  m map[string]sub2.Sub2A\n} \n}", structMeta1.UML)

	interfaceImpls := analysisTool1.findInterfaceImpls(interfaceMeta)
	assert.Equal(t, 2, len(interfaceImpls))
	assert.Equal(t, "git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml.IA <|- git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml.SA\n", interfaceImpls[0].implInterfaceUML(interfaceMeta))

	assert.Equal(t, 2, len(analysisTool1.dependencyRelations))
	assert.Equal(t, "git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml.SA ---> git.oschina.net\\\\jscode\\\\go_package_plantuml\\\\testdata\\\\uml\\\\sub2.Sub2A : c", analysisTool1.dependencyRelations[0].uml)

}