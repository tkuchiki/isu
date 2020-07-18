package printer

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type Printer struct {
	table *tablewriter.Table
	out   io.Writer
}

func New(out io.Writer) *Printer {
	return &Printer{
		out: out,
	}
}

func (p *Printer) PrintTable(headers []string, data [][]string) {
	table := tablewriter.NewWriter(p.out)
	table.SetHeader(headers)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
