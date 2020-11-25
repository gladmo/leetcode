package leet

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type TestLog struct {
	table  *tablewriter.Table
	length int
}

var (
	failColor = []tablewriter.Colors{{}, {}, {tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgRedColor}}
	passColor = []tablewriter.Colors{{}, {}, {tablewriter.FgHiGreenColor, tablewriter.Bold, tablewriter.BgGreenColor}}
)

func NewTestLog(length int) *TestLog {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"序号", "用例名称", "测试状态", "失败原因"})

	return &TestLog{table: table, length: length}
}

func (th *TestLog) SetMaxLength(length int) {
	th.length = length
}

func (th *TestLog) Fail(index, name, reason interface{}) {
	info := []string{fmt.Sprintf("%s/%d", fmt.Sprint(index), th.length), fmt.Sprint(name), "FAILED", fmt.Sprint(reason)}

	th.table.Rich(info, failColor)
}

func (th *TestLog) Pass(index, name interface{}) {
	info := []string{fmt.Sprintf("%s/%d", fmt.Sprint(index), th.length), fmt.Sprint(name), "PASS", ""}

	th.table.Rich(info, passColor)
}

func (th *TestLog) Render() {
	th.table.Render()
}
