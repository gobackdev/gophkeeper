package main

import (
	"flag"
	"fmt"

	"github.com/gobackev/gophkeeper/internal/buildinfo"
	agentcfg "github.com/gobackev/gophkeeper/internal/config/agent"
)

func main() {
	showVersion := flag.Bool("version", false, "print client version and exit")
	configPath := flag.String("config", "", "path to agent JSON config")
	flag.Parse()

	if *showVersion {
		cfg, _ := agentcfg.LoadAgentConfig(*configPath)
		version := buildinfo.Version
		date := buildinfo.Date

		if cfg != nil && cfg.Version.Version != "" {
			version = cfg.Version.Version
		}
		if cfg != nil && cfg.Version.Date != "" {
			date = cfg.Version.Date
		}

		fmt.Printf("version=%s date=%s\n", version, date)
		return
	}
}
