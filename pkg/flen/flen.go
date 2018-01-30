package flen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type FuncLenT struct {
	Name     string
	Size     int
	Filepath string
	fType    int
}

func DoFlen(cPkgs []string) {
	flenInfo := &FuncLenT{}
	for _, v := range cPkgs {
		flenInfo.GenerateFuncLens(v)
		fmt.Printf("%+v\n", flenInfo)
	}

}
func (flen *FuncLenT) GenerateFuncLens(pkg string) error {
	fset := token.NewFileSet()
	pkgs, ferr := parser.ParseDir(fset, pkg, func(f os.FileInfo) bool {
		//if opts.IncludeTests {
		return true
		//}
		//return !strings.HasSuffix(f.Name(), "_test.go")
	}, parser.AllErrors)
	if ferr != nil {
		fmt.Println("generateFuncLens error:", ferr)
		os.Exit(-1)
	}
	for _, v := range pkgs {
		for filepath, astf := range v.Files {
			for _, decl := range astf.Decls {
				ast.Inspect(decl, func(node ast.Node) bool {
					var lb, rb token.Pos
					var rln, lln, diff int
					if x, ok := node.(*ast.FuncDecl); ok {
						flen.fType = implemented
						flen.Name = x.Name.Name
						if x.Body == nil {
							flen.fType = implementedAtRuntime // externally implemented
						} else {
							lb = x.Body.Lbrace
							rb = x.Body.Rbrace
							if !lb.IsValid() || !rb.IsValid() {
								return false
							}
							rln = fset.Position(rb).Line
							lln = fset.Position(lb).Line
							diff = rln - lln - 1
							if diff == -1 {
								diff = 1 // single line func
							}
						}
						flen.Filepath = filepath
						flen.Size = diff
					}
					return false
				})
			}
		}
	}
	fmt.Printf("%+v", flen)
	return nil
}
