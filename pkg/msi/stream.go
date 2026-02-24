package msi

import "github.com/cuhsat/go-mscfb/pkg/mscfb"

type Streams struct {
	Entries *mscfb.Entries
}

func NewStreams(entries *mscfb.Entries) *Streams {
	return &Streams{
		Entries: entries,
	}
}

func (s *Streams) Next() string {
	for {
		entry := s.Entries.Next()
		if entry == nil {
			return ""
		}

		if !entry.IsStream() ||
			entry.Name == DigitalSignatureStreamName ||
			entry.Name == SummaryInfoStreamName ||
			entry.Name == DigitalSignatureExStreamName {
			continue
		}

		name, isTable := NameDecode(entry.Name)
		if !isTable {
			return name
		}
	}
}
