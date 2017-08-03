package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// input literature name example:Burger, H., Kuželički, J., & Marinek, Č. R. T. (2005). Transition from sitting to standing after trans-femoral amputation. Prosthetics and Orthotics International, 29(2), 139-151. doi: 10.1080/03093640500199612

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	path := strings.Replace(text, "\n", "", -1)

	f, e := ioutil.ReadFile(path)
	if e == nil {
		content := string(f)
		literatureList := strings.Split(content, "\n")
		literatureNames := []string{}
		count := len(literatureList)
		for i := 0; i < count; i++ {
			literature := literatureList[i]
			if len(literature) > 0 {
				nameSplit := strings.SplitAfter(literature, "). ")
				if len(nameSplit) > 1 {
					nameString := nameSplit[1]
					nameSplit := strings.Split(nameString, ".")
					if len(nameSplit) > 0 {
						name := nameSplit[0]
						literatureNames = append(literatureNames, name)
					}
				}
			}
		}

		if len(literatureNames) > 0 {
			fmt.Print(literatureNames)
		}

		pwd := "/Users/mengxiangjian/Desktop/Go/src/github.com/mengxiangjian13/searchliterature/"
		literatureFiles := []string{}
		walkFn := func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				literatureFiles = append(literatureFiles, path)
			}
			return nil
		}
		filepath.Walk(pwd, walkFn)
		for i := 0; i < len(literatureNames); i++ {
			name := literatureNames[i]
			localFile := pathForLocalLiterature(literatureFiles, name)
			if len(localFile) > 0 {
				fmt.Print(localFile + "\n")
			}
		}
	} else {
		fmt.Print(e)
	}
}

func pathForLocalLiterature(literatureFiles []string, literatureName string) string {
	length := 0
	result := ""
	for i := 0; i < len(literatureFiles); i++ {
		literaturePath := literatureFiles[i]
		_, file := filepath.Split(literaturePath)
		name := trimFileExtension(file)
		if strings.HasPrefix(literatureName, name) && len(name) > length {
			length = len(name)
			result = literaturePath
		}
	}
	return result
}

func trimFileExtension(file string) string {
	if strings.HasPrefix(file, ".") {
		// 如果是.开头的文件返回文件名
		return file
	}
	ex := filepath.Ext(file)
	name := file[0 : len(file)-len(ex)]
	return name
}
