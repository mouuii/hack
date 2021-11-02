package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "system",
	Short: "system",
	Long:  "system is root command.",
}

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	rootCmd.PersistentFlags().StringP("features", "f", "", "feature list")
	_ = viper.BindPFlag("features", rootCmd.PersistentFlags().Lookup("features"))
	rootCmd.Execute()
	cc := viper.GetString("features")
	fmt.Println(cc)
}
