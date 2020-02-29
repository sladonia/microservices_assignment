package storage

import (
	"port_domain_service/src/portpb"
	"sync"
)

var Storage StorageInterface = &inmemoryStorage{Ports: make(map[string]*portpb.Port)}

type StorageInterface interface {
	Get(portAbbreviation string) (*portpb.Port, error)
	Upsert(port *portpb.Port) (inserted, updated int32)
}

type inmemoryStorage struct {
	sync.RWMutex
	Ports map[string]*portpb.Port
}

func (s *inmemoryStorage) Upsert(port *portpb.Port) (int32, int32) {
	s.Lock()
	defer s.Unlock()

	inserted := int32(0)
	updated := int32(0)
	_, ok := s.Ports[port.Abbreviation]
	if ok {
		updated = 1
	} else {
		inserted = 1
	}
	s.Ports[port.Abbreviation] = port
	return inserted, updated
}

func (s *inmemoryStorage) Get(portAbbreviation string) (*portpb.Port, error) {
	s.RLock()
	defer s.RUnlock()

	port, ok := s.Ports[portAbbreviation]
	if !ok {
		return nil, NewItemNotFoundError(portAbbreviation)
	}
	return port, nil
}
