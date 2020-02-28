package json_parser

import (
	"client_api/src/domains"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

var testPortData = strings.NewReader(`{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  }
}`)

func getNPortsReader(n int) io.Reader {
	portString := `"AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"},`

	builder := strings.Builder{}
	builder.WriteString("{")
	for i := 0; i < n; i++ {
		builder.WriteString(portString)
	}
	st := builder.String()
	return strings.NewReader(st[:len(st)-1] + "}")
}

func TestGetPortsChannel(t *testing.T) {

	t.Run("success", func(tt *testing.T) {
		portsCh, err := GetPortsChannel(testPortData)
		assert.Nil(tt, err)
		assert.NotNil(tt, portsCh)
		ports := make([]domains.Port, 0, 3)
		for port := range portsCh {
			ports = append(ports, port)
		}
		assert.Len(tt, ports, 3)
	})

	t.Run("failure invalid json body", func(tt *testing.T) {
		portsCh, err := GetPortsChannel(strings.NewReader("fur-fur-fut"))
		assert.NotNil(tt, err)
		assert.Nil(tt, portsCh)
	})
}

func BenchmarkGetPortsChannel(b *testing.B) {
	b.Run("small json", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {

			f, err := os.Open("../../../small_ports.json")
			if err != nil {
				panic(err)
			}
			portsCh, err := GetPortsChannel(f)
			for port := range portsCh {
				_ = port
			}
		}
	})

	b.Run("large json", func(bb *testing.B) {
		for i := 0; i < bb.N; i++ {
			f, err := os.Open("../../../ports.json")
			if err != nil {
				panic(err)
			}
			portsCh, err := GetPortsChannel(f)
			for port := range portsCh {
				_ = port
			}
		}
	})
}
