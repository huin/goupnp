package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"github.com/tmc/fix"
)

func main() {
	fix.Register(fix.Fix{
		Name: "foo",
		Date: "2018-10-19",
		F:    fooFix,
		Desc: "Fix foos.",
	})
	fix.Main()
}

type pkgFix struct {
	pkgPath    string
	newPkgPath string
	name       string
}

// Fixes for a symbol in a package.
type symFix struct {
	newPkg string
}

// Maps from v1 package path to the information to migrate to v2.
var pkgFixes = map[string]*pkgFix{
	"github.com/huin/goupnp/dcps/goupnp": {
		name: "goupnp",
		// Maps from symbol in package to fixes for it.
		symFixes: map[string]*symFix{
			"RootDevice": {
				newPkg: "github.com/huin/goupnp/v2/metadata",
			},
		},
	},
	"github.com/huin/goupnp/dcps/internetgateway1": {
		newPkgPath: "github.com/huin/goupnp/v2/dcps/internetgateway1",
		name:       "internetgateway1",
	},
}

type fileCtx struct {
	// Map from name that a package is imported as to the pkgFix.
	pkgFixes map[string]*pkgFix

	// Map from package import path to the identifier to import it as (or empty string).
	newImports map[string]string
}

func fooFix(v *ast.File) bool {
	fc := &fileCtx{
		pkgFixes: make(map[string]*pkgFix),
	}
	// Discover how old goupnp packages are imported and with what name they are
	// referred to, and update the import paths.
	for _, imp := range v.Imports {
		if imp.Path.Kind != token.STRING {
			continue
		}
		impPath, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if pf, ok := pkgFixes[impPath]; ok {
			name := pf.name
			if imp.Name != nil {
				name = imp.Name.Name
			}
			fc.pkgFixes[name] = pf

			// Fix import path.
			if pf.newPkgPath != "" {
				imp.Path.Value = strconv.Quote(pf.newPkgPath)
			}
		}
	}
	if len(fc.pkgFixes) == 0 {
		// No goupnp packages imported.
		return false
	}

	fix.Walk(v, func(n interface{}) {
		switch n := n.(type) {
		case *ast.CallExpr:
			fmt.Println("CallExpr", n)
			fixCallExpr(n.Fun, fc)
		case *ast.SelectorExpr:
			fmt.Println("SelectorExpr", n)
		}
	})

	// TODO: update use of found goupnp packages.

	return true
}

func fixCallExpr(n ast.Node, fc *fileCtx) {
	if sel, ok := n.(*ast.SelectorExpr); ok {
		fmt.Println("  ", sel.X, sel.Sel)
	}
}
