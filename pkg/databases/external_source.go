package databases

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ExternalSource struct {
	SourceURL       string
	DestinationPath string
}

func (s *ExternalSource) PullIfNotExists() error {
	// If CSV file does not exist, download & create it
	if _, err := os.Stat(s.DestinationPath); os.IsNotExist(err) {
		// Create leading directories
		leadingDir, _ := filepath.Split(s.DestinationPath)
		if err := os.MkdirAll(leadingDir, os.ModeDir); err != nil {
			return err
		}

		// Create file
		out, err := os.Create(s.DestinationPath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Download file
		res, err := http.Get(s.SourceURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// Write to file
		if _, err := io.Copy(out, res.Body); err != nil {
			return err
		}
	}

	return nil
}
