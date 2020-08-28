package terminal

import (
	"os"
    "code.cloudfoundry.org/cli/cf/terminal"
    "code.cloudfoundry.org/cli/cf/trace"
    "bytes"
)

var ui terminal.UI

func init() {
    var b bytes.Buffer
    ui = terminal.NewUI(os.Stdin, os.Stdout, terminal.NewTeePrinter(os.Stdout), trace.NewWriterPrinter(&b, false))
}

func Fail(mssg string) {
    ui.Failed(mssg)
    os.Exit(1)
}

func Table(headers []string) *terminal.UITable{
	return ui.Table(headers)
}

func Add(table *terminal.UITable, row ...string) {
    table.Add(row...)
}

func PrintTable(table *terminal.UITable) {
    table.Print()
}