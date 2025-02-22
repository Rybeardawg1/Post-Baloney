package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/anacrolix/torrent/bencode"
)

type TorrentMeta struct {
	Announce     string     `bencode:"announce"`
	AnnounceList [][]string `bencode:"announce-list"`
	Info         struct {
		Name        string `bencode:"name"`
		PieceLength int    `bencode:"piece length"`
		Length      int64  `bencode:"length"`
		Files       []struct {
			Length int64    `bencode:"length"`
			Path   []string `bencode:"path"`
		} `bencode:"files"`
	} `bencode:"info"`
}

func ParseTorrentFile(filePath string) (*TorrentMeta, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var meta TorrentMeta
	err = bencode.Unmarshal(bytes.NewReader(data), &meta)
	if err != nil {
		return nil, fmt.Errorf("error parsing torrent file: %w", err)
	}

	return &meta, nil
}
