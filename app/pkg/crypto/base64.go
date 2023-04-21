package crypto

import (
	"context"
	"encoding/base64"
)

func DecodeBase64(ctx context.Context, token string) (payload string, err error) {
	bpayload, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return
	}
	payload = string(bpayload)
	return
}