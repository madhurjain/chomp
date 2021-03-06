package packets

import (
	"bytes"
	"fmt"
	"github.com/madhurjain/chomp/webservice/Godeps/_workspace/src/github.com/pborman/uuid"
	"io"
)

//SubscribePacket is an internal representation of the fields of the
//Subscribe MQTT packet
type SubscribePacket struct {
	FixedHeader
	MessageID uint16
	Topics    []string
	Qoss      []byte
	uuid      uuid.UUID
}

func (s *SubscribePacket) String() string {
	str := fmt.Sprintf("%s\n", s.FixedHeader)
	str += fmt.Sprintf("MessageID: %d topics: %s", s.MessageID, s.Topics)
	return str
}

func (s *SubscribePacket) Write(w io.Writer) error {
	var body bytes.Buffer
	var err error

	body.Write(encodeUint16(s.MessageID))
	for i, topic := range s.Topics {
		body.Write(encodeString(topic))
		body.WriteByte(s.Qoss[i])
	}
	s.FixedHeader.RemainingLength = body.Len()
	packet := s.FixedHeader.pack()
	packet.Write(body.Bytes())
	_, err = packet.WriteTo(w)

	return err
}

//Unpack decodes the details of a ControlPacket after the fixed
//header has been read
func (s *SubscribePacket) Unpack(b io.Reader) {
	s.MessageID = decodeUint16(b)
	payloadLength := s.FixedHeader.RemainingLength - 2
	for payloadLength > 0 {
		topic := decodeString(b)
		s.Topics = append(s.Topics, topic)
		qos := decodeByte(b)
		s.Qoss = append(s.Qoss, qos)
		payloadLength -= 2 + len(topic) + 1 //2 bytes of string length, plus string, plus 1 byte for Qos
	}
}

//Details returns a Details struct containing the Qos and
//MessageID of this ControlPacket
func (s *SubscribePacket) Details() Details {
	return Details{Qos: 1, MessageID: s.MessageID}
}

//UUID returns the unique ID assigned to the ControlPacket when
//it was originally received. Note: this is not related to the
//MessageID field for MQTT packets
func (s *SubscribePacket) UUID() uuid.UUID {
	return s.uuid
}
