package completion

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var completionWriter io.Writer = os.Stdout

func Run(cmd *cobra.Command, completionFlag string) error {
	switch completionFlag {
	case "bash":
		if err := cmd.Root().GenBashCompletion(completionWriter); err != nil {
			return err
		}
	case "zsh":
		if err := cmd.Root().GenZshCompletion(completionWriter); err != nil {
			return err
		}
	case "fish":
		if err := cmd.Root().GenFishCompletion(completionWriter, true); err != nil {
			return err
		}
	case "powershell":
		if err := cmd.Root().GenPowerShellCompletionWithDesc(completionWriter); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%v: invalid shell", completionFlag)
	}
	return nil
}
