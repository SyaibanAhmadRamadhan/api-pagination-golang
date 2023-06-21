package encodedecode

import (
	"encoding/base64"
	"encoding/json"
)

func Encode(data any) string {
	byte, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	encode := base64.StdEncoding.EncodeToString(byte)
	return encode
}

func Decode(data string) ([]byte, error) {
	decode, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return decode, nil
}
