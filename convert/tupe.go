package convert

type Tuple struct {
	elements []interface{}
}

func NewTuple(elements ...interface{}) *Tuple {
	return &Tuple{elements: elements}
}

func (t *Tuple) Get(index int) interface{} {
	if t == nil || len(t.elements) <= index {
		return nil
	}
	return t.elements[index]
}

func (t *Tuple) GetString(index int) string {
	return ToString(t.Get(index))
}

func (t *Tuple) GetInt(index int) int {
	return ToInt(t.Get(index))
}

func (t *Tuple) GetInt64(index int) int64 {
	return ToInt64(t.Get(index))
}

func (t *Tuple) GetInt32(index int) int32 {
	return ToInt32(t.Get(index))
}

func (t *Tuple) Len() int {
	return len(t.elements)
}
