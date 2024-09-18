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

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runServiceCode()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serviceCmd.Flags().String("service", "s", "generate service code")
}

func runServiceCode() {
	serviceTem, err := os.ReadFile("./templates/service.template")
	if err != nil {
		panic(err)
	}

	checkModel(models)
	var ms []dt.ModelInfo
	for _, model := range models {
		strService := compactStr("./service/", strings.ToLower(model.Name), "_", "service.go")
		if _, err := os.Stat(strService); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}
	checkModel(ms)

	if err = generateServiceCode(ms, string(serviceTem)); err != nil {
		panic(err)
	}
}

func generateServiceCode(models []dt.ModelInfo, serviceTemplate string) error {
	serviceTmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return err
	}

	for _, model := range models {
		_, err = os.Stat("service")
		if os.IsNotExist(err) {
			if err = os.Mkdir("service", 0777); err != nil {
				panic(err)
			}
		}

		serviceFileName := fmt.Sprintf("service/%s_service.go", strings.ToLower(model.Name))
		if err = do.GenerateFile(serviceTmpl, serviceFileName, model); err != nil {
			return err
		}
	}

	return nil
}
