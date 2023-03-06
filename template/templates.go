package template

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func mkdir(dir string) {
	if err := os.Chdir(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0750)
	}
}

// This function setup a directory and file which specifies by user
func setup(dir string, template_name string, mainFileName string) {
	// Checking about the folder
	mkdir(dir)
	// Loading template

	file, err := os.Open(template_name)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()

	// Read the contents of the file into a byte slice
	data := make([]byte, 300)
	count, err := file.Read(data)
	if err != nil {
		println(err.Error())
		return
	}
	var template_data = string(data[:count])

	// creating main file in that language
	var fileName = "./" + dir + "/" + mainFileName
	f, err := os.Create(fileName)
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()
	f.WriteString(template_data)
	return
}

// it is for python setup

func Python(dir string) {
	setup(dir, "templates/py.txt", "main.py")
}

// This function for go template setup
// It is perfect condition. it is working
func Go(dir string) {
	// creating go mod file
	var d = dir
	setup(d, "templates/go.txt", "hello_world.go")
	var name string
	fmt.Print("Enter module name: ")
	fmt.Scanln(&name)
	f, err := os.Create(dir + "/go.mod")
	if err != nil {
		println(err.Error())
		return

	}
	defer f.Close()
	version := strings.TrimPrefix(runtime.Version(), "go")
	tokens := strings.Split(version, ".")
	version = ""
	var sep string
	for i := 0; i < len(tokens)-1; i++ {
		version += sep + tokens[i]
		sep = "."
	}

	var text = fmt.Sprintf("module %v \n\ngo %v\n", name, version)
	f.WriteString(text)
	return
}

func Cxx(dir string) {

	setup(dir, "templates/cpp.txt", "main.cpp")
}
func Html(dir string) {
	setup(dir, "templates/html.txt", "main.html")
}
func Rust(dir string) { return }
func Java(dir string) {
	setup(dir, "templates/java.txt", "main.go")
}
func Node(dir string) {
	setup(dir, "templates/node.txt", "main.js")
}
func JavaScript(dir string) {
	setup(dir, "templates/js.txt", "index.js")
}
