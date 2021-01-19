package checks

import (
	"crypto/tls"
	"crypto/x509"
	rabbithole "github.com/Serviceware/rabbit-hole/v2"
	"io/ioutil"
	"log"
	"net/http"
)

type CLientConfig struct {
	Address string

	CaCert     string
	ClientCert string
	ClientKey  string

	Username string
	Password string
}

func (c *CLientConfig) NewRabbitMQClient() (*rabbithole.Client, error) {
	if c.ClientCert != "" {
		tlsConfig := &tls.Config{
			Certificates: c.certificates(),
			RootCAs:      c.rootCAs(),
		}

		transport := &http.Transport{TLSClientConfig: tlsConfig}
		return rabbithole.NewTLSClient(c.Address, c.Username, c.Password, transport)
	}

	return rabbithole.NewClient(c.Address, c.Username, c.Password)
}

func (c *CLientConfig) certificates() []tls.Certificate {
	cer, err := tls.LoadX509KeyPair(c.ClientCert, c.ClientKey)

	if err != nil {
		log.Fatal(err.Error())
	}

	return []tls.Certificate{cer}
}

func (c *CLientConfig) rootCAs() *x509.CertPool {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	certs, err := ioutil.ReadFile(c.CaCert)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	return rootCAs
}
