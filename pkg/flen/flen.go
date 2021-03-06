package flen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type funcLenT struct {
	name     string
	size     int
	filepath string
	fType    int
}
type FlenPathT struct {
	Filepath string
	FlenS    []funcLenT
	Count    map[string]int
	All      int
}

func DoFlen(cPkgs []string) {
	flenInfo := &FlenPathT{}
	flenInfo.Count = make(map[string]int)
	for _, v := range cPkgs {
		flenInfo.GenerateFuncLens(v)
	}
	fmt.Printf("%+v\n", flenInfo.Count)
	fmt.Printf("%d\n", flenInfo.All)

}
func (flen *FlenPathT) getPkgAllSize() (count int) {
	for _, v := range flen.FlenS {
		count = count + v.size
	}
	return count
}
func (flen *FlenPathT) GenerateFuncLens(pkg string) error {
	fLen := &funcLenT{}
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
	flen.Filepath = pkg
	if len(pkgs) == 0 {
		fLen.name = ""
		fLen.size = 0
		fLen.fType = implemented
		flen.FlenS = append(flen.FlenS, *fLen)
		return nil
	}
	for _, v := range pkgs {
		for filepath, astf := range v.Files {
			for _, decl := range astf.Decls {
				ast.Inspect(decl, func(node ast.Node) bool {
					var lb, rb token.Pos
					var rln, lln, diff int
					if x, ok := node.(*ast.FuncDecl); ok {
						fLen.fType = implemented
						fLen.name = x.Name.Name
						if x.Body == nil {
							fLen.fType = implementedAtRuntime // externally implemented
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
						flen.Count[pkg] = diff
						fLen.filepath = filepath
						fLen.size = diff
						flen.FlenS = append(flen.FlenS, *fLen)
					}
					return false
				})
			}
		}
	}
	flen.Count[pkg] = flen.getPkgAllSize()
	return nil
}
