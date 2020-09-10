package databases

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/jszwec/csvutil"
)

type Service struct {
	ServiceName             string `csv:"Service Name"`
	PortNumber              string `csv:"Port Number"`
	TransportProtocol       string `csv:"Transport Protocol"`
	Description             string `csv:"Description"`
	Assignee                string `csv:"Assignee"`
	Contact                 string `csv:"Contact"`
	RegistrationDate        string `csv:"Registration Date"`
	ModificationDate        string `csv:"Modification Date"`
	Reference               string `csv:"Reference"`
	ServiceCode             string `csv:"Service Code"`
	UnauthorizedUseReported string `csv:"Unauthorized Use Reported"`
	AssignmentNotes         string `csv:"Assignment Notes"`
}

type ServiceNamesPortNumbersDatabase struct {
	dbPath   string
	services map[int]*Service
}

func NewServiceNamesPortNumbersDatabase(dbPath string) *ServiceNamesPortNumbersDatabase {
	return &ServiceNamesPortNumbersDatabase{dbPath, make(map[int]*Service)}
}

func (d *ServiceNamesPortNumbersDatabase) Open() error {
	// Read CSV file
	contents, err := ioutil.ReadFile(d.dbPath)
	if err != nil {
		return err
	}

	var rawServices []Service
	if err := csvutil.Unmarshal(contents, &rawServices); err != nil {
		return err
	}

	// Take into account port ranges
	// TODO: Make the value an array for the different protocols
	for _, service := range rawServices {
		// Ignore empty port numbers
		if service.PortNumber == "" {
			continue
		}

		rangeEnd := strings.Split(service.PortNumber, "-")

		startPort, err := strconv.Atoi(rangeEnd[0])
		if err != nil {
			log.Fatal(err)
		}

		d.services[startPort] = &service

		if len(rangeEnd) > 1 {
			endPort, err := strconv.Atoi(rangeEnd[1])
			if err != nil {
				log.Fatal(err)
			}

			delta := endPort - startPort

			for i := 1; i <= delta; i++ {
				d.services[startPort+i] = &service
			}
		}
	}

	return nil
}

func (d *ServiceNamesPortNumbersDatabase) GetService(port int) (*Service, error) {
	service := d.services[port]

	if service == nil {
		return nil, errors.Errorf("could not find service for port %v", port)
	}

	return service, nil
}
