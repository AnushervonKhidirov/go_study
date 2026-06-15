package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"task_tracker/pkg/validation"
)

const (
	IDParseBase = 10
	IDBitSize   = 32
)

func ParseID(idStr string) (int, error) {
	id, err := strconv.ParseInt(idStr, IDParseBase, IDBitSize)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %w", err)
	}
	return int(id), nil
}

func ValidateAndDecode[Type any](body io.Reader, structure *Type) error {
	err := json.NewDecoder(body).Decode(&structure)

	if err != nil {
		return err
	}

	err = validation.Validate(structure)

	return nil
}
