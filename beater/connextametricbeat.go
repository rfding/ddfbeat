package beater

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/rfding/connextametricbeat/config"
)

// Connextametricbeat configuration.
type Connextametricbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

type metric struct {
	Data []struct {
		Timestamp string `json:"timestamp"`
		Value float64 `json:"value"`
	} `json:"data"`
	//Title string `json:"title"`
}

// New creates an instance of connextametricbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Connextametricbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts connextametricbeat.
func (bt *Connextametricbeat) Run(b *beat.Beat) error {
	logp.Info("connextametricbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// create custom http client
	httpClient, err := createCustomClient()
	if err != nil {
		return err
	}

	// get all the metric names
	metricNamesURL := "https://localhost:8993/services/internal/metrics/"
	body, err := getResponse(metricNamesURL, httpClient)
	if err != nil {
		return err
	}
	metricNames, err := parseMetricNames(body)

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		if err != nil {
			return err
		}
		metricData := make(map[string]metric, len(metricNames))
		// for each of the metric names, find the value
		for _, name := range metricNames {
			// dateOffset=60 may return empty data, so use dateOffset=120 to ensure data is returned
			url := "https://localhost:8993/services/internal/metrics/" + name + ".json?dateOffset=120"
			body, err := getResponse(url, httpClient)
			if err != nil {
				return err
			}
			metricData[name], err = parseMetric(body)
			if err != nil {
				return err
			}
		}
		fmt.Println(metricData)

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
			},
		}

		// add metric values to the beat event fields
		for name, data := range metricData {
			latest := len(data.Data) - 1
			if latest >= 0 {
				event.Fields[name] = data.Data[latest].Value
			}
		}

		bt.client.Publish(event)
		logp.Info("Event sent")
	}
}

// Stop stops connextametricbeat.
func (bt *Connextametricbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func createCustomClient() (http.Client, error) {
	var httpClient http.Client

	if strings.ToLower(os.Getenv("DDFBeat_SSLSkip")) == "true" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		httpClient = http.Client {
			Timeout: time.Second * 10,
		}
	} else {
		caCert, err := ioutil.ReadFile(os.Getenv("DDFBeat_CertPath"))
		if err != nil {
			return http.Client{}, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		httpClient = http.Client {
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
			Timeout: time.Second * 10,
		}
	}

	return httpClient, nil
}

// makes a get request to the given url using the client provided
func getResponse(url string, client http.Client) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func parseMetricNames(body []byte) ([]string, error) {
	var metrics map[string]interface{}
	jsonErr := json.Unmarshal(body, &metrics)
	if jsonErr != nil {
		return nil, jsonErr
	}

	metricNames := make([]string, 0, len(metrics))
	for k := range metrics {
		metricNames = append(metricNames, k)
	}

	return metricNames, nil
}

func parseMetric(body []byte) (metric, error) {
	metric1 := metric{}
	jsonErr := json.Unmarshal(body, &metric1)
	if jsonErr != nil {
		return metric{}, jsonErr
	}

	return metric1, nil
}
