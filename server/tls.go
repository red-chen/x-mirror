package server

import (
	"crypto/tls"
	"io/ioutil"
)

func LoadTLSConfig(crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {
	var (
		crt  []byte
		key  []byte
		cert tls.Certificate
	)

	if crt, err = ioutil.ReadFile(crtPath); err != nil {
		return
	}

	if key, err = ioutil.ReadFile(keyPath); err != nil {
		return
	}

	if cert, err = tls.X509KeyPair(crt, key); err != nil {
		return
	}

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return
}
