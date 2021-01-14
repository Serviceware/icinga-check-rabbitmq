package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"net/http"
)

type AlivenessCheck struct {
	client *rabbithole.Client
}

type Aliveness struct {
	Status string `json:"status"`
}

func NewAlivenessCheck(client *rabbithole.Client) Check {
	return &AlivenessCheck{client: client}
}

func (c *AlivenessCheck) DoCheck() int {
	aliveness, err := c.Aliveness()

	if err != nil {
		println(err.Error())
		return 2
	}

	println(aliveness.Status)
	return 0
}

// Overview returns an overview of cluster state with some key aggregated metrics.
func (c *AlivenessCheck) Aliveness() (rec *Aliveness, err error) {
	req, err := c.newGETRequest(c.client, "aliveness-test/labs")
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(c.client, req, &rec); err != nil {
		return nil, err
	}

	return rec, nil
}

func (c *AlivenessCheck) newGETRequest(client *rabbithole.Client, path string) (*http.Request, error) {
	s := client.Endpoint + "/api/" + path
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.SetBasicAuth(client.Username, client.Password)

	return req, err
}

func executeAndParseRequest(client *rabbithole.Client, req *http.Request, rec interface{}) (err error) {
	res, err := executeRequest(client, req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&rec); err != nil {
		return err
	}

	return nil
}

func executeRequest(client *rabbithole.Client, req *http.Request) (resp *http.Response, err error) {
	httpc := &http.Client{
		//Timeout: client.timeout,
	}
	//if client.transport != nil {
	//	httpc.Transport = client.transport
	//}
	resp, err = httpc.Do(req)
	if err != nil {
		return nil, err
	}

	if err = parseResponseErrors(resp); err != nil {
		if resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}

	return resp, err
}

func parseResponseErrors(res *http.Response) (err error) {
	if res.StatusCode == http.StatusUnauthorized {
		return errors.New("Error: API responded with a 401 Unauthorized")
	}

	// handle a "404 Not Found" response for a DELETE request as success.
	if res.Request.Method == http.MethodDelete && res.StatusCode == http.StatusNotFound {
		return nil
	}

	if res.StatusCode >= http.StatusBadRequest {
		rme := rabbithole.ErrorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&rme); err != nil {
			rme.Message = fmt.Sprintf("Error %d from RabbitMQ: %s", res.StatusCode, err)
		}
		rme.StatusCode = res.StatusCode
		return rme
	}
	return nil
}
