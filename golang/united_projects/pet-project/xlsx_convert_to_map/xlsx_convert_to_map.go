package xlsxconverttomap

import (
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func Convert() {
	f, err := excelize.OpenFile("mapping.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create("mapping.go")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	fmt.Fprintln(out, "package mapping")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "var ExecutorMap = map[string]string{")
	for _, row := range rows[1:] {
		if len(row) < 2 {
			continue
		}
		from := row[0]
		to := row[1]
		if from == "" || to == "" {
			continue
		}
		fmt.Fprintf(out, "\t%q: %q,\n", from, to)
	}
	fmt.Fprintln(out, "}")

}
