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
	Lbrace   int
	Rbrace   int
	fType    string
}

func GenerateFuncLens(pkg string) ([]FuncLenT, error) {

	fset := token.NewFileSet()
	pkgs, ferr := parser.ParseDir(fset, pkg, func(f os.FileInfo) bool {
		//if opts.IncludeTests {
		return true
		//}
		//return !strings.HasSuffix(f.Name(), "_test.go")
	}, parser.AllErrors)
	if ferr != nil {
		panic(ferr)
	}
	flens := make([]FuncLenT, 0)
	for _, v := range pkgs {
		for filepath, astf := range v.Files {
			for _, decl := range astf.Decls {
				ast.Inspect(decl, func(node ast.Node) bool {
					var (
						funcname string
						diff     int
						lb, rb   token.Pos
						rln, lln int
						ftype    string
					)

					if x, ok := node.(*ast.FuncDecl); ok {
						ftype = "implemented"
						funcname = x.Name.Name
						if x.Body == nil {
							ftype = "implementedAtRuntime" // externally implemented
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
						flens = append(flens, FuncLenT{
							Name:     funcname,
							Size:     diff,
							Filepath: filepath,
							Lbrace:   lln,
							Rbrace:   rln,
							fType:    ftype,
						})
					}
					return false

				})
			}
		}
	}
	fmt.Printf("%+v", flens)
	return flens, nil
}
