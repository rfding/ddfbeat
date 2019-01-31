package beater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMetricNames(t *testing.T) {
	assert := assert.New(t)

	// test basic functionality
	testJSON := []byte("{\"name\":\"john\",\"age\":22,\"class\":\"some class\"}")
	fields, err := parseMetricNames(testJSON)

	assert.Contains(fields, "name", "should contain this field")
	assert.Contains(fields, "age", "should contain this field")
	assert.Contains(fields, "class", "should contain this field")
	assert.Len(fields, 3, "should contain 3 strings")
	assert.Nil(err)

	// test empty JSON
	testJSON = []byte("{}")
	fields, err = parseMetricNames(testJSON)

	assert.Equal(fields, []string{}, "should be empty")
	assert.Nil(err)

	// test error
	testJSON = []byte("{{")
	fields, err = parseMetricNames(testJSON)

	assert.NotNil(err)
}

func TestParseMetric(t *testing.T) {
	assert := assert.New(t)

	// test basic functionality
	testJSON := []byte("{\"data\":[{\"value\":4.2,\"timestamp\":\"Jan 24 2019 11:03:00\"}, {\"value\":3,\"timestamp\":\"Jan 24 2019 11:04:00\"}], \"title\":\"Source Codice Confluence Queries for Jan 24 2019 11:02:56 to Jan 24 2019 11:04:56\",\"totalCount\":0}")
	metrics, err := parseMetricData(testJSON)
	parsedMetric := metric{
		Data: []struct {
			Timestamp string  `json:"timestamp"`;
			Value     float64 `json:"value"`
		}{
			{
				Timestamp: "Jan 24 2019 11:03:00",
				Value: 4.2,
			},
			{
				Timestamp: "Jan 24 2019 11:04:00",
				Value: 3,
			},
		},
	}
	assert.Equal(metrics, parsedMetric)
	assert.Nil(err)

	metrics, err = parseMetricData([]byte("{{"))
	assert.NotNil(err)
}
