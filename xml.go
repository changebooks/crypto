package crypto

import "encoding/xml"

func XmlDecode(s string, v interface{}) error {
	return xml.Unmarshal([]byte(s), v)
}

func XmlEncode(v interface{}) (string, error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
