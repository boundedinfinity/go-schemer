package schema

func (s SharedSchema) CopyT() SharedSchema {
	return s.copy()
}

func (t *SharedSchema) MergeT(other SharedSchema) error {
	return t.merge(other)
}
