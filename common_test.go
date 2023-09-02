package card

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {
	var x JSONObject
	err := json.Unmarshal([]byte(""), &x)
	t.Log(err)
}
