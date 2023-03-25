package bridgeservice

import (
	"errors"
	"unicode/utf8"

	v1 "github.com/Big-Kotik/transparent-data-bridge-api/bridge/api/v1"
)

func validateRequest(req *v1.SendFileRequest) error {
	if req == nil {
		return errors.New("empty req")
	}
	
	if !utf8.Valid([]byte(req.FileName)) {
		return errors.New("file name must be utf8 encoded string")
	}

	length := utf8.RuneCountInString(req.FileName)
	if length == 0 || length >= 100 {
		return errors.New("length must be > 0 and < 100")
	}

	return nil
}
