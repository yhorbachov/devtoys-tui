package cmd

import (
	"devtoys-tui/ui/toyselection"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "toys",
	Run: func(cmd *cobra.Command, args []string) {
		model := toyselection.NewModel()

		if _, err := tea.NewProgram(model).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {}
