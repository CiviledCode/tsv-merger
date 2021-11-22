package merger

type StringMerger struct {
	Config TSVConfig

	Content map[string]string
}

// Put ...
func (s StringMerger) Put(key string, value interface{}) {
	s.Content[key] += value.(string)
}

func (s StringMerger) Output(file string) error {
	return nil
}
