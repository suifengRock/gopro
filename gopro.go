package main

import (
	"fmt"
	"os"
)

func main() {

	arg_num := len(os.Args)
	if arg_num == 1 || arg_num > 2 {
		fmt.Println("please add the project name！！！")
		fmt.Println("like then :")
		fmt.Println("    gopro <your project name>")
		return
	}

	proName := os.Args[1]
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	dir, _ := os.Getwd()
	proPath := dir + path + proName
	subdir := []string{"bin", "pkg", "src"}
	for _, val := range subdir {
		err := os.MkdirAll(proPath+path+val, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	mkfile := proPath + path + "Makefile"
	mf, err := os.Create(mkfile)
	if err != nil {
		panic(err)
	}
	mf.WriteString("export GOPATH :=" + proPath + "\n")
	mf.WriteString("export PATH := ${PATH}:${GOPATH}\\bin\n")
	mf.WriteString("export GOBIN := ${GOPATH}\\bin\n")
	mf.WriteString("main:\n")
	mf.WriteString("	go run main.go\n")

	defer mf.Close()

	mainfile := proPath + path + "main.go"
	mainf, err := os.Create(mainfile)
	if err != nil {
		panic(err)
	}
	mainf.WriteString("package main\n")
	mainf.WriteString("import (\n")
	mainf.WriteString("	\"fmt\"\n")
	mainf.WriteString(")\n")
	mainf.WriteString("func main() {\n")
	mainf.WriteString("	fmt.Println(\"project is ready!!!\")\n")
	mainf.WriteString("}\n")

	defer mainf.Close()

	fmt.Println("project is ready...")

}
