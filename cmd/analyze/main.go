package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/adject1/macondo/analyzer"
	"github.com/adject1/macondo/config"
)

func main() {
	// Determine the directory of the executable. We will use this
	// directory to find the data files if an absolute path is not
	// provided for these!
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	cfg := &config.Config{}
	cfg.Load([]string{})
	log.Info().Msgf("Loaded config: %v", cfg)
	cfg.AdjustRelativePaths(exPath)
	cfg.Debug = false

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	an := analyzer.NewAnalyzer(cfg)
	err = an.RunTest()
	if err != nil {
		fmt.Println(err)
	}
}
