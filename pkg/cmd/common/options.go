package common

import (
	"github.com/Benbentwo/utils/log"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"io"
	"os"
	"strconv"
)

const (
	OptionBatchMode = "batch-mode"
	OptionVerbose   = "verbose"
)

type CommonOptions struct {
	Cmd       *cobra.Command
	Args      []string
	BatchMode bool
	Verbose   bool
	In        terminal.FileReader
	Out       terminal.FileWriter
	Err       io.Writer
}

// AddBaseFlags adds the base flags for all commands
func (o *CommonOptions) AddBaseFlags(cmd *cobra.Command) {
	defaultBatchMode := false
	if os.Getenv("BATCH_MODE") == "true" {
		defaultBatchMode = true
	}
	cmd.PersistentFlags().BoolVarP(&o.BatchMode, OptionBatchMode, "b", defaultBatchMode, "Runs in batch mode without prompting for user input")
	cmd.PersistentFlags().BoolVarP(&o.Verbose, OptionVerbose, "", false, "Enables verbose output")

	o.Cmd = cmd
}

func SetLoggingLevel(cmd *cobra.Command, args []string) {
	verbose, err := strconv.ParseBool(cmd.Flag(OptionVerbose).Value.String())
	if err != nil {
		util.Logger().Errorf("Unable to determine log level")
	}

	if verbose {
		err := log.SetLevel("debug")
		if err != nil {
			util.Logger().Errorf("Unable to set log level to debug")
		}
	} else {
		err := log.SetLevel("info")
		if err != nil {
			util.Logger().Errorf("Unable to set log level to info")
		}
	}
}
