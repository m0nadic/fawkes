package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


func writeEnv(c *gin.Context) {
	for _, k := range os.Environ() {
		fmt.Fprintln(c.Writer, k)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(getStatus(isHealthy()), gin.H{
		"healthy": isHealthy(),
	})
}

func die(c *gin.Context) {
	log.Println("about to die")
	go func() {
		time.Sleep(100 * time.Millisecond)
		os.Exit(1)
	}()
	c.Data(http.StatusOK, "text/plain", []byte("OK"))
}

func getStatus(h bool) int {
	if h {
		return http.StatusOK
	} else {
		return http.StatusInternalServerError
	}
}

func sick(c *gin.Context) {
	makeSick()
	c.Data(http.StatusOK, "text/plain", []byte("OK"))
}

func recover(c *gin.Context) {
	makeHealthy()
	c.Data(http.StatusOK, "text/plain", []byte("OK"))
}

func memSpike(c *gin.Context) {
	s := []string{"..."}
	go func() {
		for {
			s = append(s, s...)
		}
	}()
	c.Data(http.StatusOK, "text/plain", []byte("OK"))
}

func cpuSpike(c *gin.Context) {
	go func() {
		time.Sleep(100 * time.Millisecond)
		for {
			_ = 0
		}
	}()
	c.Data(http.StatusOK, "text/plain", []byte("OK"))
}

var bindAddr = flag.String("b", "0.0.0.0:8888", "bind address")

func main() {
	flag.Parse()
	log.Println("starting fawkes ...")

	router := gin.Default()

	router.GET("/env", writeEnv)
	router.GET("/health", healthCheck)
	router.POST("/die", die)
	router.POST("/sick", sick)
	router.POST("/recover", recover)

	router.POST("/spike/mem", memSpike)
	router.POST("/spike/cpu", cpuSpike)

	panic(router.Run(*bindAddr))
}
