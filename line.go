package expression


type Line string

func (l Line) Get(key string) string {
	return string(l)
}
