package importbug

import (
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"testing"
)

func TestImportBug(t *testing.T) {
	buildPkg, err := build.Import("github.com/netbrain/importbug/foo", "", build.ImportComment)
	if err != nil {
		t.Fatal(err)
	}

	fset := token.NewFileSet()
	packages, err := parser.ParseDir(fset, buildPkg.Dir, nil, 0)
	if err != nil {
		t.Fatal(err)
	}

	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	for pName, p := range packages {
		files := make([]*ast.File, 0, len(p.Files))
		for _, f := range p.Files {
			files = append(files, f)
		}

		conf := &types.Config{
			FakeImportC: true,
			Importer:    importer.Default(),
		}
		_, err := conf.Check(pName, fset, files, info)
		if err != nil {
			log.Fatal(err)
		}
	}
}
