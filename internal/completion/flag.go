package completion

import (
	"github.com/spf13/cobra"
)

func CompletionFlag(cmd *cobra.Command, b *string) {
	cmd.Flags().StringVar(b, "completion", "", "Output command-line completion code for the specified shell (zsh|bash|fish|powershell)")
	err := cmd.RegisterFlagCompletionFunc("completion", completionCompletion)
	if err != nil {
		panic(err)
	}
}

func completionCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"bash", "zsh", "fish", "powershell"}, cobra.ShellCompDirectiveNoFileComp
}
