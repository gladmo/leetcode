package leetcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func (th SaveOption) golang() (err error) {
	baseDir := th.saveDir
	solutionDir := path.Join(baseDir, th.language, "solution")
	err = os.MkdirAll(solutionDir, 0755)
	if err != nil {
		return
	}

	code, ok := parseGoCode(th.codeSnippet)
	if !ok {
		code = `func Export()  {

}`
	}

	solution := fmt.Sprintf(`package solution

%s

/**************************************/
/******** 以下为 Leetcode 源码部分 *******/
/**************************************/

%s
`, code, th.codeSnippet)

	err = ioutil.WriteFile(
		path.Join(solutionDir, fmt.Sprintf("%s.go", th.titleSlug)),
		[]byte(solution),
		0755)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(
		path.Join(solutionDir, fmt.Sprintf("%s.go.clean", th.titleSlug)),
		[]byte(solution), 0755)
	if err != nil {
		return
	}

	mainCode := fmt.Sprintf(`package main

import (
	"github.com/gladmo/leetcode/%s"
)

func main() {
	/*
     
	%s

    */

	solution.Export()
}
`, solutionDir, th.sampleTestCase)

	err = ioutil.WriteFile(path.Join(baseDir, th.language, "main.go"), []byte(mainCode), 0755)
	if err != nil {
		return
	}

	return
}

func parseGoCode(code string) (newCode string, ok bool) {
	ok = true

	start := strings.Index(code, `func `)
	comment := code[:start]

	var structCode string

	if len(strings.TrimSpace(comment)) > 0 {
		typeIndex := strings.Index(code, `type`)
		commentRightParentheses := strings.Index(code, `}`)
		if typeIndex < commentRightParentheses {
			sc := comment[typeIndex:commentRightParentheses]
			for _, c := range strings.Split(sc, "\n") {
				structCode += "\n" + strings.Replace(strings.TrimSpace(c), "*", "", 1)
			}

			structCode += "}\n\n"
		}
	}

	code = code[start:]
	start = 0
	leftParentheses := strings.Index(code, `(`)

	if strings.Count(code, `func `) != 1 {
		ok = false
		return
	}

	funcName := code[start+5 : leftParentheses]

	rightParentheses := strings.Index(code, `)`)
	paramsStr := code[leftParentheses+1 : rightParentheses]
	paramsList := strings.Split(paramsStr, ",")

	var params []string
	for _, s := range paramsList {
		s = strings.TrimSpace(s)
		p := strings.Split(s, " ")
		params = append(params, p[0])
	}

	leftCurly := strings.Index(code, `{`)
	resultParams := code[rightParentheses+1 : leftCurly]

	exportFunction := fmt.Sprintf(`%s(%s)`, funcName, strings.Join(params, ", "))
	if len(strings.TrimSpace(resultParams)) > 0 {
		exportFunction = fmt.Sprintf(`return %s`, exportFunction)
	}

	code = structCode + code[:leftCurly+2] + "\t" + exportFunction + code[leftCurly+2:]

	newCode = strings.Replace(code, funcName, "Export", 1)
	return
}
