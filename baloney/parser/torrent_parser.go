package parser

import (
	"fmt"
	"os"

	"github.com/anacrolix/torrent/bencode"
)

// TorrentMeta represents the structure of a .torrent file
type TorrentMeta struct {
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	Info         struct {
		Name        string `bencode:"name"`
		PieceLength int    `bencode:"piece length"`
		Length      int64  `bencode:"length"` // Size of the file (if single file)
		Files       []struct {
			Length int64    `bencode:"length"`
			Path   []string `bencode:"path"`
		} `bencode:"files"` // List of files (if multi-file torrent)
	} `bencode:"info"`
}

// ParseTorrentFile reads and parses a .torrent file
func ParseTorrentFile(filePath string) (*TorrentMeta, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var meta TorrentMeta
	err = bencode.Unmarshal(data, &meta) // FIX: Pass []byte directly
	if err != nil {
		return nil, fmt.Errorf("error parsing torrent file: %w", err)
	}

	return &meta, nil
}
