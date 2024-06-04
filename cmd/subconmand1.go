/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subconmand1Cmd represents the subconmand1 command
var subconmand1Cmd = &cobra.Command{
	Use:   "subconmand1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subconmand1 called")
		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(author)
		fmt.Println(args)
	},
}

// func (c *Command) SetHelpCommand(cmd *Command)
// func (c *Command) SetHelpCommandGroupID(groupID string)
// func (c *Command) SetHelpFunc(f func(*Command, []string))
// func (c *Command) SetHelpTemplate(s string)

func init() {
	rootCmd.AddCommand(subconmand1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subconmand1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subconmand1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	subconmand1Cmd.Flags().StringP("author", "a", "lanzhiwang@163.com", "Help message for Author")
}
