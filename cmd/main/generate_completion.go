package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GenerateCompletion(flag *pflag.FlagSet) error {
	command := &cobra.Command{
		Use: "completions",
	}

	command.Flags().AddFlagSet(flag)
	os.Mkdir("completions/", 0755)
	os.MkdirAll("completions/bash", 0755)
	os.MkdirAll("completions/zsh", 0755)
	os.MkdirAll("completions/fish", 0755)
	os.MkdirAll("completions/ps1", 0755)
	command.GenBashCompletionFileV2("completions/bash/tallylingo", true)
	command.GenZshCompletionFile("completions/zsh/tallylingo")
	command.GenFishCompletionFile("completions/fish/tallylingo", true)
	command.GenPowerShellCompletionFile("completions/ps1/tallylingo")
	return nil
}
