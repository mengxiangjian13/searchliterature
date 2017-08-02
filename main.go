package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
						fmt.Print(name)
						fmt.Print("\n")
					}
				}
			}
		}
	} else {
		fmt.Print(e)
	}
}
