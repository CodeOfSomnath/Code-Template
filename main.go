// Now all links are hardcorded. So that is not ready for production
package main

import (
	"flag"
	"fmt"
	"github.com/somnath/tpkg/template"
)

// This function is for parsing arguments
// This function implementation complete
func args() []string {
	var values = make([]string, 0)
	var (
		LanguageType = flag.String("new", "python", "Which language to be use for making template.\n\t1.python\n\t2.go\n\t3.java\n\t4.rust\n\t5.cpp\n\t6.node\n")
		dir          = flag.String("dir", ".", "Where to create project.")
	)
	flag.Parse()
	values = append(values, *LanguageType)
	values = append(values, *dir)
	return values
}

// This checks type of template need
// Also call the template function
func TemplateCall(dir, lang string) {
	if lang == "python" {
		template.Python(dir)
	} else if lang == "go" {
		template.Go(dir)
	} else if lang == "rust" {
		template.Rust(dir)
	} else if lang == "cpp" {
		template.Cxx(dir)
	} else if lang == "html" {
		template.Html(dir)
	} else if lang == "node" {
		template.Node(dir)
	} else if lang == "java" {
		template.Java(dir)
	} else if lang == "javascript" {
		template.JavaScript(dir)
	} else {
		fmt.Println("error: No template found.")
	}
}

func main() {
	var values = args()
	var lang = values[0]
	var dir = values[1]
	TemplateCall(dir, lang)
}
