package cli

import (
	"fmt"
	"github.com/genome-engine/genome/script"
	"github.com/spf13/cobra"
	"os"
)

//Commands
const (
	Run       string = "run" //genome exec {script_name}: take the script with the specified name from the /scripts folder
	GetScript string = "get-script"
	Welcome   string = `

	   ____     _____      _   _      ___      __  __     _____ 	
	 / ___|    | ___ |    | \ | |    / _ \    |  \/  |   | ___ |	
	| |  _     |  _|      |  \| |   | | | |   | |\/| |   |  _| 	
	| |_| |    | |___     | |\  |   | |_| |   | |  | |   | |___ 	
	 \____|    |_____|    |_| \_|    \___/    |_|  |_|   |_____|

	+---------------------------+
	|    Available commands:    |
	+---------------------------+
		
	+---+
	|run|		script_path.yaml(can be unsuffixed) 
	+---+
	+-----------+
	|get-script |	creates a script file.
	| / gs      |
	+-----------+
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
					fmt.Println(err.Error())
					return
				}

				if err = s.Execute(); err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Printf("\n%v script execution finished\n\n", arg)
			}
		},
	}

	getScriptCmd = &cobra.Command{
		Use:     GetScript,
		Aliases: []string{"gs"},
		Run: func(cmd *cobra.Command, args []string) {
			var filename string
			var source = `parses:
 - path: analyze_path

 #If you want to use global templates and the global variable [GENOME_TEMPS] 
 #is declared, unlock the glob_temps flag.
#glob_temps:true
templates: 
 - path: template_path

generate:
 path: path_to_out

  #insert mode is used to add new text to a file with existing text.
 #mode: insert 

  #add a mark to the file (if mode: insert) 
  #insert:label_name; #insert-end:label_name
 #label: label_name

 #delimiters: 0 - {{}}, 1 - <>
#delimiters: 1 

 #if you want to see the generation process displayed in the console.
#logs: true
`

			switch len(args) > 0 {
			case true:
				filename = args[0] + ".yml"
			case false:
				filename = "script.yml"
				break
			}

			file, err := os.Create(filename)
			if err != nil && !os.IsExist(err) {
				fmt.Println(err.Error())
				return
			}

			if _, err = file.WriteString(source); err != nil {
				fmt.Println(err.Error())
				return
			}
			_ = file.Close()
		},
	}
)
