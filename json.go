package crypto

import "encoding/json"

func JsonDecode(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

func JsonEncode(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
