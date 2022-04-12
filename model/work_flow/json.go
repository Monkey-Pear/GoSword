package work_flow

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSON json.RawMessage

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
