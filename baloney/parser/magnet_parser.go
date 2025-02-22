package parser

import (
	"fmt"
	"net/url"
	"strings"
)

type MagnetMeta struct {
	InfoHash string
	Name     string
	Trackers []string
	Size     string
}

func ParseMagnetLink(magnetURI string) (*MagnetMeta, error) {
	parsedURL, err := url.Parse(magnetURI)
	if err != nil {
		return nil, fmt.Errorf("invalid magnet link: %w", err)
	}

	if parsedURL.Scheme != "magnet" {
		return nil, fmt.Errorf("not a valid magnet link")
	}

	params := parsedURL.Query()
	magnetData := &MagnetMeta{}

	xt := params.Get("xt")
	if strings.HasPrefix(xt, "urn:btih:") {
		magnetData.InfoHash = strings.TrimPrefix(xt, "urn:btih:")
	} else {
		return nil, fmt.Errorf("invalid magnet link: missing info hash")
	}

	magnetData.Name = params.Get("dn")
	magnetData.Trackers = params["tr"]
	magnetData.Size = params.Get("xl")

	return magnetData, nil
}
