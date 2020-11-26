package leet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func (th SaveOption) golang() (err error) {
	baseDir := th.SaveDir
	solutionDir := path.Join(baseDir, "solution")
	err = os.MkdirAll(solutionDir, 0755)
	if err != nil {
		return
	}

	code, ok := parseGoCode(th.CodeSnippet)
	if !ok {
		code = `
func Export() {

}`
	}

	solution := fmt.Sprintf(`package solution

%s

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

%s
`, code, th.CodeSnippet)

	err = ioutil.WriteFile(
		path.Join(solutionDir, fmt.Sprintf("%s.go", th.TitleSlug)),
		[]byte(solution),
		0755)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(
		path.Join(solutionDir, fmt.Sprintf("%s.go.clean", th.TitleSlug)),
		[]byte(solution), 0755)
	if err != nil {
		return
	}

	mainCode := fmt.Sprintf(`package main

import (
	"github.com/gladmo/leetcode/%s"
	"github.com/gladmo/leetcode/leet"
)

func main() {
	/*
     
	%s

    */

	tests := []struct {
		name  string
		input [][]int
		want  bool
	}{
		{
			name: "test-[[1],[2],[3],[]]",
			input: [][]int{
				{1},
				{2},
				{3},
				{},
			},
			want: true,
		},
		{
			name: "test-[[1,3],[3,0,1],[2],[0]]",
			input: [][]int{
				{1, 3},
				{3, 0, 1},
				{2},
				{0},
			},
			want: false,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
`, solutionDir, th.SampleTestCase)

	err = ioutil.WriteFile(path.Join(baseDir, "main.go"), []byte(mainCode), 0755)
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

func GolangClear(dir string) (err error) {
	goSolutionDir := path.Join(dir, "golang", "solution")
	_, err = os.Stat(goSolutionDir)
	if err != nil {
		return
	}

	f, err := os.Open(goSolutionDir)
	if err != nil {
		return
	}

	fileInfo, err := f.Readdir(10)
	if err != nil {
		return
	}

	var goFiles []string
	for _, info := range fileInfo {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path.Join(goSolutionDir, info.Name()))
		}
	}

	for _, fileDir := range goFiles {
		written, err := CopyFile(fileDir, fileDir+".clean")
		if err != nil {
			fmt.Println(err.Error())
		}

		if written > 0 {
			fmt.Println(fileDir, "cleared")
		}
	}

	return
}
