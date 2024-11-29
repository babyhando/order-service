package cache

import "encoding/json"

func (c *ObjectCacher[T]) unmarshal(data []byte, out any) error {
	if c.serializationType == SerializationTypeJSON {
		return json.Unmarshal(data, out)
	}

	// implement rest of them
	return nil
}

func (c *ObjectCacher[T]) Marshal(in any) ([]byte, error) {
	if c.serializationType == SerializationTypeJSON {
		return json.Marshal(in)
	}

	// implement rest of them
	return nil, nil
}
