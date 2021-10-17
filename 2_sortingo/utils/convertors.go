package utils

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
	"path/filepath"
)

func ToTable(arr []Resources, tableName string) string {
	t := table.NewWriter()
	t.SetTitle(tableName)
	t.SetAutoIndex(true)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Size", "Swaps", "Comparisons", "Time"})

	for i := 0; i < len(arr); i++ {
		t.AppendRows([]table.Row{
			{arr[i].Size, arr[i].Swaps, arr[i].Comparisons, arr[i].Time},
		})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
	t.SetOutputMirror(nil)
	return t.RenderCSV()
}

func ToCSV(table string, filename string) {
	pathname := "AlgResources/"
	newpath := filepath.Join(".", pathname)
	os.MkdirAll(newpath, os.ModePerm)

	file, err := os.Create(pathname + filename + ".csv")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	file.WriteString(table)
	fmt.Println("Done")
}
