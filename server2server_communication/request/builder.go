package request

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type RequestBuilder struct {
	client  *http.Client
	ctx     context.Context
	headers map[string]string
	body    []byte
	result  interface{}
}

func NewHttpRequestBuilder(client *http.Client) *RequestBuilder {
	return &RequestBuilder{
		client:  client,
		headers: make(map[string]string),
	}
}

func (rb *RequestBuilder) NewRequestWithContext(ctx context.Context) *RequestBuilder {
	rb.ctx = ctx
	return rb
}

func (rb *RequestBuilder) AddHeaders(headers map[string]string) *RequestBuilder {
	for k, v := range headers {
		rb.headers[k] = v
	}
	return rb
}

func (rb *RequestBuilder) WithJSONBody(data interface{}) *RequestBuilder {
	jsonBytes, _ := json.Marshal(data)
	rb.body = jsonBytes
	return rb
}

func (rb *RequestBuilder) ResponseAs(result interface{}) *RequestBuilder {
	rb.result = result
	return rb
}

func (rb *RequestBuilder) Post(url string) error {
	req, err := http.NewRequestWithContext(rb.ctx, http.MethodPost, url, bytes.NewReader(rb.body))
	if err != nil {
		return err
	}
	for k, v := range rb.headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := rb.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if rb.result != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(bodyBytes, rb.result)
	}
	return err
}
