package main

import (
	"fmt"
	"github.com/mpetavy/common"
	"os"
	"time"
)

func init() {
	common.Init("test", "0.0.0", "", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

func run() error {
	src, err := os.ReadFile("api.js")
	if common.Error(err) {
		return err
	}

	engine, err := common.NewGojaEngine(string(src))
	if common.Error(err) {
		return err
	}

	output, err := engine.Run(time.Hour, "")
	if common.Error(err) {
		return err
	}

	common.Info(output)

	return nil
}

func main() {
	common.Run(nil)
}
