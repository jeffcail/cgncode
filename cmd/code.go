/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/jeffcail/cgncode/do"
	"github.com/jeffcail/cgncode/dt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCode()
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	codeCmd.Flags().String("code", "c", "generate code")
}

func runCode() {
	handlerTem, err := os.ReadFile("./templates/handler.template")
	if err != nil {
		panic(err)
	}

	serviceTem, err := os.ReadFile("./templates/service.template")
	if err != nil {
		panic(err)
	}

	dtoTem, err := os.ReadFile("./templates/dto.template")
	if err != nil {
		panic(err)
	}

	checkModel(models)

	var ms []dt.ModelInfo
	for _, model := range models {
		strHandler := compactStr("./handler/", strings.ToLower(model.Name), "_", "handler.go")
		if _, err = os.Stat(strHandler); errors.Is(err, os.ErrNotExist) {
			ms = append(ms, model)
		}
		if _, err = os.Stat(compactStr("./dto/", strings.ToLower(model.Name), "_", "dto.go")); errors.Is(err, os.ErrNotExist) {
			ms = append(ms, model)
		}

		if _, err = os.Stat(compactStr("./service/", strings.ToLower(model.Name), "_", "service.go")); errors.Is(err, os.ErrNotExist) {
			ms = append(ms, model)
		}
	}
	checkModel(ms)

	if err = generateCode(ms, string(handlerTem), string(serviceTem), string(dtoTem)); err != nil {
		panic(err)
	}
}

// generateHandlerCode generates handler code based on the model information
func generateCode(models []dt.ModelInfo, handlerTemplate, serviceTemplate, dtoTemplate string) error {
	handlerTmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	serviceTmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return err
	}

	dtoTmpl, err := template.New("dto").Parse(dtoTemplate)
	if err != nil {
		return err
	}

	for _, model := range models {
		// Generate handler code
		_, err = os.Stat("handler")
		if os.IsNotExist(err) {
			if err = os.Mkdir("handler", 0777); err != nil {
				panic(err)
			}
		}

		handlerFileName := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(model.Name))
		if err := do.GenerateFile(handlerTmpl, handlerFileName, model); err != nil {
			return err
		}

		// Generate service code
		_, err = os.Stat("service")
		if os.IsNotExist(err) {
			if err = os.Mkdir("service", 0777); err != nil {
				panic(err)
			}
		}
		serviceFileName := fmt.Sprintf("service/%s_service.go", strings.ToLower(model.Name))
		if err := do.GenerateFile(serviceTmpl, serviceFileName, model); err != nil {
			return err
		}

		// Generate dto code
		_, err = os.Stat("dto")
		if os.IsNotExist(err) {
			if err = os.Mkdir("dto", 0777); err != nil {
				panic(err)
			}
		}
		dtoFileName := fmt.Sprintf("dto/%s_dto.go", strings.ToLower(model.Name))
		if err := do.GenerateFile(dtoTmpl, dtoFileName, model); err != nil {
			return err
		}
	}
	return nil
}

func checkModel(models []dt.ModelInfo) {
	if len(models) == 0 {
		log.Fatal("no models")
	}
}

func compactStr(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}
