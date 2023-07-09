/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-grpc/simple-client-server/post-server"
	"go-grpc/simple-client-server/user-server"
	gateway "go-grpc/simple-gateway-server"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run SUBCOMMAND",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User gRPC Server",
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside subCmd Run with args: %v\n", args)
		user.Execute()
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	// },
	// PersistentPostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
	// },
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post gRPC Server",
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside subCmd Run with args: %v\n", args)
		post.Execute()
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	// },
	// PersistentPostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
	// },
}

var gwCmd = &cobra.Command{
	Use:   "gw",
	Short: "gRPC Gateway Server",
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside subCmd Run with args: %v\n", args)
		gateway.Execute()
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	// },
	// PersistentPostRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
	// },
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.AddCommand(userCmd)
	runCmd.AddCommand(postCmd)
	runCmd.AddCommand(gwCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
