package do

import (
	"fmt"
	"github.com/jeffcail/cgncode/dt"
	"os"
	"text/template"
)

func GenerateFile(tmpl *template.Template, fileName string, model dt.ModelInfo) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := tmpl.Execute(file, model); err != nil {
		return err
	}
	fmt.Printf("Generated file: %s\n", fileName)
	return nil
}
