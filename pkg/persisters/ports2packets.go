package persisters

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

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

type Ports2PacketPersister struct {
	*ExternalSource
	dbPath  string
	packets map[int]*Packet
}

func NewPorts2PacketPersister(dbPath string, sourceURL string) *Ports2PacketPersister {
	return &Ports2PacketPersister{
		ExternalSource: &ExternalSource{
			SourceURL:       sourceURL,
			DestinationPath: dbPath,
		},
		dbPath:  dbPath,
		packets: make(map[int]*Packet),
	}
}

func (d *Ports2PacketPersister) Open() error {
	// If CSV file does not exist, download & create it
	if err := d.ExternalSource.PullIfNotExists(); err != nil {
		return err
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

func (d *Ports2PacketPersister) GetPacket(port int) (*Packet, error) {
	packet := d.packets[port]

	if packet == nil {
		return nil, fmt.Errorf("could not find packet for port %v", port)
	}

	return packet, nil
}
