package card

type JSONObject struct {
}

func (o *JSONObject) Get(key string) *JSONObject {
	return nil
}
func (o *JSONObject) Index(idx int) *JSONObject {
	return nil
}

func (o *JSONObject) UnmarshalJSON(bs []byte) error {
	if string(bs) == "null" {
		return nil
	}
	return nil
}
