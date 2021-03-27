package databases

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jszwec/csvutil"
)

type RawPacket struct {
	Port   int    `csv:"port"`
	Packet string `csv:"packet"`
}

type Packet struct {
	Port   int
	Packet []byte
}

type Ports2PacketDatabase struct {
	dbPath    string
	sourceURL string
	packets   map[int]*Packet
}

func NewPorts2PacketDatabase(dbPath string, sourceURL string) *Ports2PacketDatabase {
	return &Ports2PacketDatabase{dbPath, sourceURL, make(map[int]*Packet)}
}

func (d *Ports2PacketDatabase) Open() error {
	// If CSV file does not exist, download & create it
	if _, err := os.Stat(d.dbPath); os.IsNotExist(err) {
		// Create leading directories
		leadingDir, _ := filepath.Split(d.dbPath)
		if err := os.MkdirAll(leadingDir, os.ModeDir); err != nil {
			return err
		}

		// Create file
		out, err := os.Create(d.dbPath)
		if err != nil {
			return err
		}
		defer out.Close()

		// Download file
		res, err := http.Get(d.sourceURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// Write to file
		if _, err := io.Copy(out, res.Body); err != nil {
			return err
		}
	}

	// Read CSV file
	contents, err := ioutil.ReadFile(d.dbPath)
	if err != nil {
		return err
	}

	var rawPackets []RawPacket
	if err := csvutil.Unmarshal(contents, &rawPackets); err != nil {
		return err
	}

	// Decode base64 encoded data
	for _, rawPacket := range rawPackets {
		content, err := base64.StdEncoding.DecodeString(rawPacket.Packet)
		if err != nil {
			return err
		}

		d.packets[rawPacket.Port] = &Packet{rawPacket.Port, content}
	}

	return nil
}

func (d *Ports2PacketDatabase) GetPacket(port int) (*Packet, error) {
	packet := d.packets[port]

	if packet == nil {
		return nil, fmt.Errorf("could not find packet for port %v", port)
	}

	return packet, nil
}
