/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jeffcail/cgncode/do"
	"github.com/jeffcail/cgncode/dt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

// handlerCmd represents the handler command
var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runHandlerCode()
	},
}

func init() {
	rootCmd.AddCommand(handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	handlerCmd.Flags().String("handler", "ha", "generate handler code")
}

func runHandlerCode() {
	handlerTem, err := os.ReadFile("./templates/handler.template")
	if err != nil {
		panic(err)
	}

	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strHandler := compactStr("./handler/", strings.ToLower(model.Name), "_", "handler.go")
		if _, err = os.Stat(strHandler); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}

	checkModel(ms)

	if err = generateHandlerCode(ms, string(handlerTem)); err != nil {
		panic(err)
	}
}

func generateHandlerCode(models []dt.ModelInfo, handlerTemplate string) error {
	handlerTmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		return err
	}
	for _, model := range models {
		_, err = os.Stat("handler")
		if os.IsNotExist(err) {
			if err = os.Mkdir("handler", 0777); err != nil {
				panic(err)
			}
		}

		handlerFileName := fmt.Sprintf("handler/%s_handler.go", strings.ToLower(model.Name))
		if err = do.GenerateFile(handlerTmpl, handlerFileName, model); err != nil {
			return err
		}
	}
	return nil
}
