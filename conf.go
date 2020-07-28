package main

// Conf type
type Conf interface {
	read(path string) (interface{}, error)
}

func confRead(conf Conf, path string) (interface{}, error) {
	return conf.read(path)
}
