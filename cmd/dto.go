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

// dtoCmd represents the dto command
var dtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runDtoCode()
	},
}

func init() {
	rootCmd.AddCommand(dtoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dtoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dtoCmd.Flags().String("dto", "d", "Help message for toggle")
}

func runDtoCode() {
	dtoTem, err := os.ReadFile("./templates/dto.template")
	if err != nil {
		panic(err)
	}

	checkModel(models)
	var ms []dt.ModelInfo

	for _, model := range models {
		strDto := compactStr("./dto/", strings.ToLower(model.Name), "_", "dto.go")
		if _, err = os.Stat(strDto); os.IsNotExist(err) {
			ms = append(ms, model)
		}
	}

	checkModel(ms)

	if err = generateDtoCode(ms, string(dtoTem)); err != nil {
		panic(err)
	}
}

func generateDtoCode(models []dt.ModelInfo, dtoTemplate string) error {
	dtoTmpl, err := template.New("dto").Parse(dtoTemplate)
	if err != nil {
		return err
	}

	for _, model := range models {
		_, err = os.Stat("dto")
		if os.IsNotExist(err) {
			if err = os.Mkdir("dto", 0777); err != nil {
				panic(err)
			}
		}

		dtoFileName := fmt.Sprintf("dto/%s_dto.go", strings.ToLower(model.Name))
		if err = do.GenerateFile(dtoTmpl, dtoFileName, model); err != nil {
			return err
		}
	}

	return nil
}
