package utils

import (
	"bufio"
	"bytes"
	"io"
)

type Event struct {
	Type string
	Data []byte
}

// A base implementation of a Decoder for text/event-stream.
type EventStreamDecoder struct {
	evt    Event
	rc     io.ReadCloser
	reader *bufio.Reader
	err    error
}

func NewEventStreamDecoder(rc io.ReadCloser) *EventStreamDecoder {
	return &EventStreamDecoder{
		rc:     rc,
		reader: bufio.NewReader(rc),
	}
}

func (s *EventStreamDecoder) Next() bool {
	if s.err != nil {
		return false
	}

	event := ""
	data := bytes.NewBuffer(nil)

	for {
		txt, err := s.reader.ReadBytes('\n')
		if err != nil {
			s.err = err
			return false
		}

		// txt is trimed with leading space
		txt = bytes.TrimSpace(txt)

		// Dispatch event on an empty line
		if len(txt) == 0 {
			s.evt = Event{
				Type: event,
				Data: data.Bytes(),
			}
			return true
		}

		// Split a string like "event: bar" into name="event" and value=" bar".
		name, value, _ := bytes.Cut(txt, []byte(":"))

		// Consume an optional space after the colon if it exists.
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		switch string(name) {
		case "":
			// An empty line in the for ": something" is a comment and should be ignored.
			continue
		case "event":
			event = string(value)
		case "data":
			_, s.err = data.Write(value)
			if s.err != nil {
				break
			}
			_, s.err = data.WriteRune('\n')
			if s.err != nil {
				break
			}
		}
	}
}

func (s *EventStreamDecoder) Event() Event {
	return s.evt
}

func (s *EventStreamDecoder) Close() error {
	return s.rc.Close()
}

func (s *EventStreamDecoder) Err() error {
	return s.err
}
