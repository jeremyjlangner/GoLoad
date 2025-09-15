package executor

import (
    "fmt"
    "io"
    "net/http"
    "goload/models"
	"errors"
	"strings"
	"time"
)

var client = &http.Client{}

func buildRequest(job models.Job) (*http.Request, error) {
	req, err := http.NewRequest(job.API.Method, job.API.URL, strings.NewReader(job.API.Body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return req, err
	}

	for i := range job.API.Headers{
		if len(job.API.Headers[i]) == 2 {
			req.Header.Add(job.API.Headers[i][0], job.API.Headers[i][1])
		} else {
			fmt.Println("Error with headers request:")
			return req, errors.New("Invalid request headers")
		}
	}

	return req, nil
}

func execRequest(request *http.Request) (status int, body []byte, err error, latency time.Duration) {
	start := time.Now()
	resp, err := client.Do(request)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	status = resp.StatusCode
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, fmt.Errorf("error: received status code %d", resp.StatusCode), elapsed
	}

	body, err = io.ReadAll(resp.Body)
    if err != nil {
        return status, nil, err, elapsed
    }

    return status, body, nil, elapsed
}

func HttpRequest(job models.Job) (int, []byte, error, time.Duration) {
    req, err := buildRequest(job)
    if err != nil {
        return http.StatusInternalServerError, nil, err, 0
    }

    return execRequest(req)
}
