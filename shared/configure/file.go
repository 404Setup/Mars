package configure

import (
	_ "embed"
	"log/slog"
	"os"
)

//go:embed config.yml
var fileData []byte

func newFile() {
	file, err := os.Create("config.yml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(fileData)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
