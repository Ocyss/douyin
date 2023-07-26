package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "douyin",
	Short:   "你所热爱的，就是你的生活。",
	Long:    `抖音让每一个人看见并连接更大的世界，鼓励表达、沟通和记录，激发创造，丰富人们的精神世界，让现实生活更美好。`,
	Version: "v0.7.26",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// no-completion
	rootCmd.AddCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}
