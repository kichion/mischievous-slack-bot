package s3

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/xerrors"
)

// GetSecret はgoogle service acountのシークレットjsonを取得します
func (s S3) GetSecret() ([]byte, error) {
	params := s.buildParams("google/secret.json")
	resp, err := s.SelectObjectContent(params)

	if err != nil {
		return []byte(""), xerrors.Errorf("s3 GetSecret error: %v", err)
	}
	defer resp.EventStream.Close() // nolint:errcheck

	for event := range resp.EventStream.Events() {
		switch v := event.(type) {
		case *s3.RecordsEvent:
			return v.Payload, nil
		}
	}

	return []byte(""), xerrors.Errorf("s3 GetSecret error: %v", errors.New("NOT FOUND"))
}
