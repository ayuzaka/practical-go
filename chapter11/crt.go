package chapter11

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 証明書を読み込む
	cert, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	cfg := &tls.Config{
		RootCAs: certPool,
	}
	cfg.BuildNameToCertificate()

	// クライアントを作成
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
			Proxy:           http.ProxyFromEnvironment,
		},
	}

	resp, err := hc.Get("https://www.oreilly.co.jp/index.shtml")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
