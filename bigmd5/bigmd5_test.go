package bigmd5

import (
	"encoding/hex"
	"testing"
)

func TestMD5File(t *testing.T) {
	file := "../smallfile.txt"
	b, err := MD5File(file)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}

	if hex.EncodeToString(b) != "4dcaf0cc26cb2217f4e7cf5ca55dbe4b" {
		t.Errorf("MD5File(%s) failed\n", file)
		return
	}
}
