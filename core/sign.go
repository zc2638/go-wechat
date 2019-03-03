package core

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"sort"
)

func Sign(params map[string]string, apiKey string, h hash.Hash) string {
	if h == nil {
		h = md5.New()
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bufw := bufio.NewWriterSize(h, 128)
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		_, _ = bufw.WriteString(k + "=" + v + "&")
	}
	_, _ = bufw.WriteString("key=" + apiKey)
	_ = bufw.Flush()

	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}
