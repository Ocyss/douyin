package main

import (
	"github.com/Ocyss/douyin/cmd"
)

func main() {
	// trace.Start(os.Stderr)
	// defer trace.Stop()
	cmd.Execute()
}
