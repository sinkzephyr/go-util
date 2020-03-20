package util

func FixUrlPath(path string) string {
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func FixUrlHost(host string) string {
	if host[0:3] != "http" {
		return "http://" + host
	}
	return host
}
