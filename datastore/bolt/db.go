package bolt

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
)

type dbKeys struct {
	user    []byte
	session []byte
	repo    []byte
}

var keys = dbKeys{
	user:    []byte("users"),
	session: []byte("sessions"),
	repo:    []byte("repo"),
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func gobEncode(p interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(p)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func gobDecode(data []byte, target *interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(target)

	if err != nil {
		return err
	}

	return nil
}

func encode(p interface{}) ([]byte, error) {
	enc, err := json.Marshal(p)

	if err != nil {
		return nil, err
	}

	return enc, nil
}

func decode(data []byte, target interface{}) error {
	err := json.Unmarshal(data, &target)

	if err != nil {
		return err
	}

	return nil
}
