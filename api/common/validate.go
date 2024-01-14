package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Validator interface {
	Validate() error
}

func ReadAndValidateRequestBody(request *http.Request, dest Validator) error {
	if body, err := io.ReadAll(request.Body); err != nil {
		return fmt.Errorf(err.Error())
	} else if err := json.Unmarshal(body, &dest); err != nil {
		return fmt.Errorf(err.Error())
	} else if err := dest.Validate(); err != nil {
		return err
	}
	return nil
}
