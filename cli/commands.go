package cli

import (
	"fmt"
	"github.com/genome-engine/genome/script"
	"github.com/spf13/cobra"
)

//Commands
const (
	Run     string = "run" //genome exec {script_name}: take the script with the specified name from the /scripts folder
	Welcome string = `

	   ____     _____      _   _      ___      __  __     _____ 	
	 / ___|    | ___ |    | \ | |    / _ \    |  \/  |   | ___ |	
	| |  _     |  _|      |  \| |   | | | |   | |\/| |   |  _| 	
	| |_| |    | |___     | |\  |   | |_| |   | |  | |   | |___ 	
	 \____|    |_____|    |_| \_|    \___/    |_|  |_|   |_____|

	+---------------------------+
	|    Available commands:    |
	+---------------------------+
		
	+---+
	|run|	script_path.yaml(can be unsuffixed) 
	+---+ 
`
)

var (
	rootCmd = &cobra.Command{
		Use: "genome",
		Run: func(cmd *cobra.Command, args []string) { println(Welcome) },
	}
	runCmd = &cobra.Command{
		Use: Run,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("You need to pass the path to the yaml file with the script")
				return
			}
			for _, arg := range args {
				fmt.Printf("%v script execution running.\n", arg)
				s, err := script.NewScript(arg)
				if err != nil {
					fmt.Print(err.Error())
					return
				}

				if err = s.Execute(); err != nil {
					fmt.Print(err.Error())
					return
				}
				fmt.Printf("\n%v script execution finished\n\n", arg)
			}
		},
	}
)
