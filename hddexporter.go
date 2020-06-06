package main

import (
	"io/ioutil"
	"net"
	"strings"
	"strconv"
	"log"
	"time"
	"net/http"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (temperatureGauge = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "hddTemp",
				Name:      "temperature_celsius",
				Help:      "Disk temperature in celsius",
			},
			[]string{"device"},
		)
);

func main() {
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(temperatureGauge)
	go func() {
		for {
			conn, err := net.DialTimeout("tcp", "localhost:7634", time.Second*3)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			data, err := ioutil.ReadAll(conn)
			if err != nil {
				log.Fatal(err)
			}

			fields := strings.Split(string(data), "|")
			for index := 0; index < len(fields)/5; index++ {
				offset := index * 5
				device := fields[offset+1]
				device = device[strings.LastIndex(device, "/")+1:]
				temperatureField := fields[offset+3]
				temperature, err := strconv.ParseFloat(temperatureField, 64)
				if err != nil {
					fmt.Println(err)
					continue
				}
				temperatureGauge.WithLabelValues(device).Set(temperature)
			}
			time.Sleep(time.Second * 10)
     		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
