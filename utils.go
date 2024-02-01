package lrpc

import "github.com/google/uuid"

func uuidArray(in []string) ([]uuid.UUID, error) {
	array := make([]uuid.UUID, len(in))
	for i, v := range in {
		id, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		array[i] = id
	}
	return array, nil
}
