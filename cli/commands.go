package cli

import (
	"github.com/genome-engine/genome/script"
	"github.com/spf13/cobra"
	"log"
)

//Commands
const (
	Exec    string = "exec" //genome exec {script_name}: take the script with the specified name from the /scripts folder
	Welcome string = `

						    ____     _____      _   _      ___      __  __     _____ 
						  / ___|    | ___ |    | \ | |    / _ \    |  \/  |   | ___ |
						 | |  _     |  _|      |  \| |   | | | |   | |\/| |   |  _| 
						 | |_| |    | |___     | |\  |   | |_| |   | |  | |   | |___ 
						  \____|    |_____|    |_| \_|    \___/    |_|  |_|   |_____|


									The genome welcomes you.


							     +---------------------------+
							     |    Available commands:    |
							     +---------------------------+
		
			+----+
			|exec|	[script_path]: runs a yaml script.
			+----+
`
)

var (
	rootCmd = &cobra.Command{
		Use: "genome",
		Run: func(cmd *cobra.Command, args []string) { println(Welcome) },
	}
	commands = []*cobra.Command{
		{
			Use: Exec,
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) == 0 {
					log.Print("You need to pass the path to the yaml file with the script")
					return
				}

				s, err := script.NewScript(args[0])
				if err != nil {
					log.Print(err.Error())
					return
				}

				if err = s.Execute(); err != nil {
					log.Print(err.Error())
					return
				}

				log.Print("The use of the scrapbook ended successfully.")
			},
		},
	}
)
