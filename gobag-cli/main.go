package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// gobag-cli usage:
func main() {
	app := cli.NewApp()
	app.Author = "Starship Technologies OÜ"
	app.Email = "technology@starship.xyz"
	app.Name = "gobag-cli"
	app.Usage = "CLI tool to explore ROS bags"
	app.Version = "1.0.0"
	app.After = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:        "dump",
			Description: "Dump indicated content of the bag ",
			Subcommands: []cli.Command{
				{
					Name:        "chunks",
					Description: "dump uncompressed chunks",
					ArgsUsage:   "<filename>",
					Action: func(c *cli.Context) error {
						filename := c.Args().Get(0)
						if filename == "" {
							cli.ShowCommandHelpAndExit(c, "chunks", -1)
						}
						err := dumpChunks(filename)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
				{
					Name:        "chunksinfo",
					Description: "dump chunk information",
					ArgsUsage:   "<filename>",
					Action: func(c *cli.Context) error {
						filename := c.Args().Get(0)
						if filename == "" {
							cli.ShowCommandHelpAndExit(c, "chunks", -1)
						}
						err := dumpChunkInfo(filename)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
				{
					Name:        "messagedefinitions",
					Description: "dump message definitions",
					ArgsUsage:   "<filename>",
					Action: func(c *cli.Context) error {
						filename := c.Args().Get(0)
						if filename == "" {
							cli.ShowCommandHelpAndExit(c, "messagedefinitions", -1)
						}
						err := dumpMessageDefinitions(filename)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
				{
					Name:        "tabledefinitions",
					Description: "dump HIVE DDL table definitions for all topics into separate .sql files",
					ArgsUsage:   "<filename>",
					Action: func(c *cli.Context) error {
						filename := c.Args().Get(0)
						if filename == "" {
							cli.ShowCommandHelpAndExit(c, "tabledefinitions", -1)
						}
						err := dumpTableDefinitions(filename)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
				{
					Name:        "json",
					Description: "dump full bag to JSON",
					ArgsUsage:   "<inputfilename> <outfilename>",
					Action: func(c *cli.Context) error {
						inputfilename := c.Args().Get(0)
						if inputfilename == "" {
							cli.ShowCommandHelpAndExit(c, "json", -1)
						}
						outputfilename := c.Args().Get(1)
						if outputfilename == "" {
							cli.ShowCommandHelpAndExit(c, "json", -1)
						}
						err := dumpJSON(inputfilename, outputfilename)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
				{
					Name:        "topics",
					Description: "dump bag messages to JSON by topic",
					ArgsUsage:   "<inputfilename> <outputdirname>",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "time",
							Value: "",
							Usage: "Comma separated Unix epoch timestamps for start and end to filter by time",
						},
						cli.StringFlag{
							Name:  "filter",
							Value: "",
							Usage: "Comma separated list of topics to limit output (including '/' prefix if needed)",
						},
					},
					Action: func(c *cli.Context) error {
						var (
							startTime    int64
							endTime      int64
							topicsFilter []string
							err          error
						)
						inputfilename := c.Args().Get(0)
						if inputfilename == "" {
							cli.ShowCommandHelpAndExit(c, "topics", -1)
						}
						outputdirname := c.Args().Get(1)
						if outputdirname == "" {
							cli.ShowCommandHelpAndExit(c, "topics", -1)
						}
						if len(c.String("time")) > 0 {
							timeSlices := strings.Split(c.String("time"), ",")
							if len(timeSlices) != 2 {
								err := fmt.Errorf("time requires 2 comma separated timestamps")
								return cli.NewExitError(err, -1)
							}
							startTime, err = strconv.ParseInt(timeSlices[0], 10, 64)
							if err != nil {
								return cli.NewExitError(err, -1)
							}
							endTime, err = strconv.ParseInt(timeSlices[1], 10, 64)
							if err != nil {
								return cli.NewExitError(err, -1)
							}
						}
						if len(c.String("filter")) > 0 {
							topicsFilter = strings.Split(c.String("filter"), ",")
						}

						err = dumpTopicsJSON(inputfilename, outputdirname, startTime, endTime, topicsFilter)
						if err != nil {
							return cli.NewExitError(err, -1)
						}
						return nil
					},
				},
			},
		},
		{
			Name:        "docs",
			Usage:       "gobag-cli docs > documentation.md",
			Description: "Generate documentation in markdown format and print to standard out.",
			Action: func(c *cli.Context) error {
				docs := GenerateDocs(c.App)
				fmt.Print(docs)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
