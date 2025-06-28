package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run initgen.go <nombre>")
		os.Exit(1)
	}

	name := strings.ToLower(os.Args[1])
	capitalized := capitalize(name)
	moduleBase := filepath.Join("..", "..", "modules", name)
	types := []string{"model", "core", "routes"}

	for _, t := range types {
		dir := filepath.Join(moduleBase, t)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("❌ Error creando carpeta %s: %v\n", dir, err)
			os.Exit(1)
		}

		filename := filepath.Join(dir, fmt.Sprintf("%s_%s.go", name, t))
		packageName := fmt.Sprintf("%s_%s", name, t)

		var content string
		switch t {
		case "model":
			content = fmt.Sprintf(`// Archivo generado automáticamente para el módulo %s (%s)
package %s

type Struct struct {
}
`, name, t, packageName)

		case "routes":
			content = fmt.Sprintf(`// Archivo generado automáticamente para el módulo %s (%s)
package %s

import (
	"net/http"
	"main/base/router"
	"main/modules/%s/core"
)

func Init() {
	var r = router.Router()
	r.GET("/medidores", http.HandlerFunc(%s_core.Controller.Read))
}
`, name, t, packageName, name, name)

		case "core":
			content = fmt.Sprintf(`// Archivo generado automáticamente para el módulo %s (%s)
package %s

import (
	"main/base/base_controller"
	"main/base/base_service"
	user_model "main/modules/%s/model"
)

var Controller = base_controller.NewController(
	*base_service.Init[user_model.Struct](nil),
)
var Service = base_service.Init[user_model.Struct](nil)
`, name, t, packageName, name)
		}

		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			fmt.Printf("❌ Error creando archivo %s: %v\n", filename, err)
			os.Exit(1)
		}

		fmt.Printf("✅ %s creado con package '%s'\n", filename, packageName)
	}

	// Crear module.go
	moduleFile := filepath.Join(moduleBase, "module.go")
	modulePkg := name

	moduleContent := fmt.Sprintf(`// Archivo de entrada para el módulo %s
package %s

import (
	"fmt"
	"main/modules/%s/routes"
)

func Init() {
	fmt.Println("%s Module Initialized")
	routes.Init()
}
`, name, modulePkg, name, capitalized)

	if err := os.WriteFile(moduleFile, []byte(moduleContent), 0644); err != nil {
		fmt.Printf("❌ Error creando module.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ %s creado con función Init()\n", moduleFile)
	fmt.Println("🎉 Módulo completo generado en '../../modules/" + name + "'")
}
