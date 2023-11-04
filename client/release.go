package client

import "net"

var (
	rootCrtPaths = []string{"/Users/liushiyu/gopher/my-x-mirror/example/san/server.pem"}
)

// server name in release builds is the host part of the server address
func serverName(addr string) string {
	host, _, err := net.SplitHostPort(addr)

	// should never panic because the config parser calls SplitHostPort first
	if err != nil {
		panic(err)
	}

	return host
}
