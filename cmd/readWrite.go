/*
Copyright © 2022 fahrizalfarid
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/fahrizalfarid/read-binary-golang/src/usecase"
	"github.com/spf13/cobra"
)

var (
	source, output string
	bufferSize     int

	reader = usecase.NewReader()
	writer = usecase.NewWriter()
)

// readWriteCmd represents the readWrite command
var readWriteCmd = &cobra.Command{
	Use:   "readWrite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readWrite called")

		buffers, err := reader.ReadFile(source, bufferSize)
		if err != nil {
			panic(err)
		}

		err = writer.WriteToFile(buffers, output)
		if err != nil {
			panic(err)
		}
		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(readWriteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readWriteCmd.PersistentFlags().String("foo", "", "A help for foo")

	readWriteCmd.Flags().StringVarP(&source, "inputFile", "i", "img.jpg", "Your filepath")
	readWriteCmd.MarkFlagRequired("inputFile")

	readWriteCmd.Flags().StringVarP(&output, "outputFile", "o", "img.png", "Output file")
	readWriteCmd.MarkFlagRequired("outputFile")

	readWriteCmd.Flags().IntVarP(&bufferSize, "bufferSize", "b", 1024, "Buffer size for reading, 1024 byte = 1KB")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readWriteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
