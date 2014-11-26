package main

import (
	"fmt"
	"os"
)

var (
	subdir = []string{"bin", "pkg", "src"}

	mainContent = []string{
		"package main",
		"import (",
		"	\"fmt\"",
		")",
		"func main() {",
		"	fmt.Println(\"project is ready!!!\")",
		"}",
	}

	dockerContent = []string{
		"FROM google/golang",
		"WORKDIR /gopath",
		"ADD . /gopath",
		"CMD [\"bash\"]",
	}
)

func main() {

	if checkArgErr() {
		return
	}

	proName := os.Args[1]
	path := getPathSeparator()
	dir, _ := os.Getwd()
	proPath := dir + path + proName

	//create subdir "src pkg bin"
	for _, val := range subdir {
		err := os.MkdirAll(proPath+path+val, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// create Makefile
	mkfile := proPath + path + "Makefile"
	mkContent := []string{
		"export GOPATH :=" + proPath,
		"export PATH := ${PATH}:${GOPATH}" + path + "bin",
		"export GOBIN := ${GOPATH}" + path + "bin",
		"main:",
		"	go run main.go",
		"build:",
		"	go install main.go",
		"images:",
		"	docker build -t " + proName + " .",
		"run:",
		"	docker run -it -v " + proPath + ":/gopath --rm " + proName,
		"docker: images run",
	}
	writeFile(mkfile, mkContent, "\n")

	//create main.go
	mainfile := proPath + path + "main.go"
	writeFile(mainfile, mainContent, "\n")

	//create dockerfile
	dockerFile := proPath + path + "Dockerfile"
	writeFile(dockerFile, dockerContent, "\n")

	fmt.Println("project is ready...")

}

func getPathSeparator() string {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	return path
}

func checkArgErr() bool {
	arg_num := len(os.Args)
	if arg_num == 1 || arg_num > 2 {
		fmt.Println("please add the project name！！！")
		fmt.Println("like then :")
		fmt.Println("    gopro <your project name>")
		return true
	}
	return false
}

func writeFile(fNamePath string, content []string, postfix string) (err error) {

	mf, err := os.Create(fNamePath)
	defer mf.Close()
	if err != nil {
		return
	}

	for _, val := range content {
		mf.WriteString(val + postfix)
	}
	return
}
