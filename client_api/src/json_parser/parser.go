package json_parser

import (
	"client_api/src/domains"
	"client_api/src/logger"
	"encoding/json"
	"io"
)

// GetPortsChannel reads JSON stream object by object and puts them down the channel
// It returns the channel and closes it when the stream is empty
func GetPortsChannel(portsJsonStream io.Reader) (<-chan domains.Port, error) {
	portsCh := make(chan domains.Port)
	decoder := json.NewDecoder(portsJsonStream)

	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}

	// read and parse json object by object
	// put Ports down the channel
	// close the channel when end reading
	go func() {
		port := domains.Port{}
		for decoder.More() {
			key, err := decoder.Token()
			if err != nil {
				logger.Logger.Debugw("unable to decode json key", "error", err)
				continue
			}

			s, ok := key.(string)
			if !ok {
				logger.Logger.Debugf("unable to pass key to string. key: %v", s)
				continue
			}

			err = decoder.Decode(&port)
			if err != nil {
				logger.Logger.Debugw("unable to decode json object", "error", err)
				continue
			}
			port.Abbreviation = s
			portsCh <- port
		}
		close(portsCh)
	}()

	return portsCh, nil
}
