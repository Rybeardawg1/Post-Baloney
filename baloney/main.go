package main

import (
	"flag"
	"fmt"

	"github.com/Rybeardawg1/Post-Baloney/baloney/config"
	"github.com/Rybeardawg1/Post-Baloney/baloney/parser"
	"github.com/Rybeardawg1/Post-Baloney/baloney/utils"
)

const version = "1.0.0"

func main() {
	cfg := config.LoadConfig()

	showVersion := flag.Bool("v", false, "Show version")
	torrentFile := flag.String("t", "", ".torrent file path")
	magnetLink := flag.String("m", "", "Magnet link")
	downloadPath := flag.String("p", "", "Download path")
	setPath := flag.String("d", "", "Set a new default path")
	showPath := flag.Bool("show", false, "Show current default path")

	flag.Parse()

	if *showPath {
		fmt.Println("Current default path:", cfg.DefaultPath)
		return
	}

	if *setPath != "" {
		if utils.IsValidPath(*setPath) {
			cfg.DefaultPath = *setPath
			config.SaveConfig(cfg)
			fmt.Println("New default path set to:", *setPath)
		} else {
			fmt.Println("Invalid directory path.")
		}
		return
	}

	if *showVersion {
		fmt.Printf("baloney version %s\n", version)
		return
	}

	if *torrentFile != "" {
		meta, err := parser.ParseTorrentFile(*torrentFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Torrent Name:", meta.Info.Name)
		fmt.Println("Announce URL:", meta.Announce)
		return
	}

	if *magnetLink != "" {
		meta, err := parser.ParseMagnetLink(*magnetLink)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Magnet Info Hash:", meta.InfoHash)
		fmt.Println("Trackers:", meta.Trackers)
		return
	}

	fmt.Println("Error: Provide a .torrent file (-t) or magnet link (-m).")
}
