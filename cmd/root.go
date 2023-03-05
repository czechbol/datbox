package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	cmdUse   = `datbox <path/to/file.xml>`
	cmdShort = "Transforms XML data from https://www.mojedatovaschranka.cz/sds/welcome?part=opendata to other formats"
	cmdLong  = `Transforms XML data from https://www.mojedatovaschranka.cz/sds/welcome?part=opendata to other formats

Drops optional parameters when unused to save on disk space.

It can also read from the STDIN stream so you can pipe data to into it.`
)

type Datbox struct {
	Format     string
	OutputFile string
}

// rootCmd represents the base command when called without any subcommands
func (d *Datbox) CobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   cmdUse,
		Short: cmdShort,
		Long:  cmdLong,
		Args:  cobra.ExactArgs(1),
		Run:   d.Run,
	}
	f := cmd.Flags()
	f.StringVarP(&d.Format, "format", "f", "json", "Output format. Allows: \"json\",\"yaml\",\"xml\"")
	f.StringVarP(&d.OutputFile, "output-file", "o", "", "Output file. Otherwise the results will be printed to stdout.")

	return cmd
}

func (d *Datbox) Run(cmd *cobra.Command, args []string) {
	stat, _ := os.Stdin.Stat()
	var byteData []byte

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			byteData = append(byteData, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		filename := args[0]
		if _, err := os.Stat(filename); err != nil {
			fmt.Println(fmt.Errorf("%w: file not found", errors.New("Invalid argument")))
			os.Exit(1)
		}
		if err := ReadXML(filename, &byteData); err != nil {
			fmt.Println(fmt.Errorf("%w: unable to read file", errors.New("File read error")))
			os.Exit(1)
		}
	}
	var boxList List
	if err := ParseXML(byteData, &boxList); err != nil {
		fmt.Println(fmt.Errorf("%w: unable to parse file", errors.New("File read error")))
		os.Exit(1)
	}
	var outData []byte

	switch {
	case strings.ToLower(d.Format) == "json":
		if err := GetJSON(&boxList, &outData); err != nil {
			fmt.Println(fmt.Errorf("%w: JSON output failed", errors.New("Output error")))
			os.Exit(1)
		}
	case strings.ToLower(d.Format) == "yaml":
		if err := GetYAML(&boxList, &outData); err != nil {
			fmt.Println(fmt.Errorf("%w: YAML output failed", errors.New("Output error")))
			os.Exit(1)
		}
	case strings.ToLower(d.Format) == "xml":
		if err := GetXML(&boxList, &outData); err != nil {
			fmt.Println(fmt.Errorf("%w: XML output failed", errors.New("Output error")))
			os.Exit(1)
		}
	default:
		fmt.Println(fmt.Errorf("%w: --output, -o Allows only: \"json\",\"yaml\",\"xml\" (default json)", errors.New("Invalid argument")))
		os.Exit(1)
	}
	if d.OutputFile != "" {
		WriteToFile(&outData, d.OutputFile)
		os.Exit(0)
	}
	fmt.Println(string(outData))
	os.Exit(0)

}
