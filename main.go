package main

import (
	"bytes"
	"fmt"
	"github.com/mpetavy/common"
	"os"
	"time"
)

func init() {
	common.Init("test", "0.0.0", "", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

func readSourcecode() (string, error) {
	files, err := common.ListFiles("*.js", false)
	if common.Error(err) {
		return "", err
	}

	var buf bytes.Buffer

	for _, file := range files {
		src, err := os.ReadFile(file)
		if common.Error(err) {
			return "", err
		}

		_, err = buf.Write(src)
		if common.Error(err) {
			return "", err
		}
	}

	return buf.String(), nil
}

func run() error {
	src, err := readSourcecode()
	if common.Error(err) {
		return err
	}

	message, err := os.ReadFile("adt_a01.hl7")
	if common.Error(err) {
		return err
	}

	engine, err := common.NewGojaEngine(string(src))
	if common.Error(err) {
		return err
	}

	output, err := engine.Run(time.Hour, "main", string(message))
	if common.Error(err) {
		return err
	}

	common.Info(output)

	return nil
}

func main() {
	common.Run(nil)
}
