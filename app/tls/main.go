package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//https://www.cnblogs.com/dust90/p/11207219.html

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	// 启动https方式访问
	cert, err := tls.LoadX509KeyPair("./configs/cert/server.pem", "./configs/cert/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./configs/cert/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	ReadTimeout := time.Duration(60) * time.Second
	WriteTimeout := time.Duration(60) * time.Second

	s := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    certPool,
		},
	}

	s.ListenAndServeTLS("./configs/cert/server.pem", "./configs/cert/server.key")
}
