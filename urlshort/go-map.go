package main

type gomap struct {
	data map[string]string
}

var dict = map[string]string{
	"/gog": "http://google.com",
	"/mil": "http://milliyet.com",
}

func newMap() (*gomap, error) {
	return &gomap{
		data: dict,
	}, nil
}

func (g *gomap) getData(path string) (string, bool) {
	value, ok := g.data[path]
	return value, ok
}
