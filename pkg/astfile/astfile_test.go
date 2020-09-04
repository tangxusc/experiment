package astfile

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
)

func TestAstFile(t *testing.T) {
	source, err2 := os.Open("../hystrix/hystrix.go")
	if err2 != nil {
		panic(err2.Error())
	}
	defer func() {
		err := source.Close()
		if err != nil {
			panic(err.Error())
		}
	}()
	all, err2 := ioutil.ReadAll(source)
	if err2 != nil {
		panic(err2.Error())
	}

	set := token.NewFileSet()
	file, err := parser.ParseFile(set, source.Name(), all, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}
	ast.Inspect(file, func(node ast.Node) bool {
		fmt.Println(node)
		switch node.(type) {
		//case *ast.Comment:
		//	cm := node.(*ast.Comment)
		//	fmt.Println("==========comment==============")
		//	fmt.Println(cm.Slash, cm.Text, cm.End(), cm.Pos())
		//	fmt.Println("==========comment==============")
		//	return false
		//case *ast.CommentGroup:
		//	cmg := node.(*ast.CommentGroup)
		//	fmt.Println("==============commentGroup===========")
		//	fmt.Println(cmg.Pos(), cmg.End(), cmg.Text())
		//	for i, comment := range cmg.List {
		//		fmt.Printf("    i:%v,slash:%v,text:%v,end:%v,pos:%v \n", i, comment.Slash, comment.Text, comment.End(), comment.Pos())
		//	}
		//	fmt.Println("==============commentGroup===========")
		//	return false
		case *ast.GenDecl:
			gen := node.(*ast.GenDecl)

			fmt.Println("gen.doc.text", gen.Doc.Text())
			for i, spec := range gen.Specs {
				fmt.Printf("gen.spec[%v] begin", i)
				switch spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec.(*ast.ImportSpec)
					fmt.Print(importSpec)
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					//typeSpec.Name.Name
					//structType := typeSpec.Type.(*ast.StructType)
					//structType.Fields.List
					fmt.Print(typeSpec)
				case *ast.ValueSpec:
					valueSpec := spec.(*ast.ValueSpec)
					fmt.Print(valueSpec)
				}
				fmt.Printf("gen.spec[%v] end \n", i)

			}

		case *ast.StructType:

			structType := node.(*ast.StructType)
			fmt.Println(structType)
		}
		return true
	})
	syscall.Exec()
	//ast.Print(set, file)
	//ast.Walk()
}
