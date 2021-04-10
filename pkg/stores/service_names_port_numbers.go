package stores

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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
	*ExternalSource
	dbPath   string
	services map[int][]Service
}

func NewServiceNamesPortNumbersDatabase(dbPath string, sourceURL string) *ServiceNamesPortNumbersDatabase {
	return &ServiceNamesPortNumbersDatabase{
		ExternalSource: &ExternalSource{
			SourceURL:       sourceURL,
			DestinationPath: dbPath,
		},
		dbPath:   dbPath,
		services: make(map[int][]Service),
	}
}

func (d *ServiceNamesPortNumbersDatabase) Open() error {
	// If CSV file does not exist, download & create it
	if err := d.ExternalSource.PullIfNotExists(); err != nil {
		return err
	}

	// Read CSV file
	contents, err := ioutil.ReadFile(d.dbPath)
	if err != nil {
		return err
	}

	var rawServices []Service
	if err := csvutil.Unmarshal(contents, &rawServices); err != nil {
		return err
	}

	for _, service := range rawServices {
		rangePoints := strings.Split(service.PortNumber, "-")

		// Skip services with empty ports
		rawStartPort := rangePoints[0]
		if rawStartPort == "" {
			continue
		}

		startPort, err := strconv.Atoi(rawStartPort)
		if err != nil {
			return err
		}

		d.services[startPort] = append(d.services[startPort], service)

		// Port range
		if len(rangePoints) > 1 {
			rawEndPort := rangePoints[1]

			endPort, err := strconv.Atoi(rawEndPort)
			if err != nil {
				return err
			}

			for currentPort := startPort + 1; currentPort <= endPort; currentPort++ {
				d.services[currentPort] = append(d.services[currentPort], service)
			}
		}
	}

	return nil
}

// GetService returns the services that match the port and protocol given
// Use "*" as the protocol to find all services on the port independent of protocol
func (d *ServiceNamesPortNumbersDatabase) GetService(port int, protocol string) ([]Service, error) {
	allServicesForProtocol := d.services[port]
	if allServicesForProtocol == nil {
		return nil, fmt.Errorf("could not find service(s) for port %v", port)
	}

	outServices := make([]Service, 0)
	for _, service := range allServicesForProtocol {
		if service.TransportProtocol == protocol || protocol == "*" {
			outServices = append(outServices, service)
		}
	}

	if len(outServices) < 1 {
		return nil, fmt.Errorf("could find service(s) for port %v, but not for protocol %v on that port", port, protocol)
	}

	return outServices, nil
}
