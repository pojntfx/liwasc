package databases

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

type Ports2PacketDatabase struct {
	dbPath  string
	packets map[int]*Packet
}

func NewPorts2PacketDatabase(dbPath string) *Ports2PacketDatabase {
	return &Ports2PacketDatabase{dbPath, make(map[int]*Packet)}
}

func (d *Ports2PacketDatabase) Open() error {
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
