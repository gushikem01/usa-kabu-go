package httpclient

import (
	"bytes"
	"time"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type HTTP struct {
	Client *fasthttp.Client
	logger *zap.Logger
}

// HTTP HTTPクライアント
func NewHTTP(l *zap.Logger) *HTTP {
	h := &HTTP{Client: &fasthttp.Client{
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
	}, logger: l}
	return h
}

// Get Getリクエスト
func (h *HTTP) Get(url string, header map[string]string) ([]byte, error) {
	b, err := h.request(fasthttp.MethodGet, url, header)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// request ...
// func (h *HTTP) request(method, url string, header map[string]string, b io.Reader, query io.Reader) (*fasthttp.Response, error) {
func (h *HTTP) request(method, url string, header map[string]string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)

	// fasthttp does not automatically request a gzipped response.
	// We must explicitly ask for it.
	for k, v := range header {
		req.Header.Set(k, v)
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, err
	}
	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		return nil, err
	}
	contentEncoding := resp.Header.Peek("Content-Encoding")
	var body []byte
	if bytes.EqualFold(contentEncoding, []byte("gzip")) {
		body, _ = resp.BodyGunzip()
	} else {
		body = resp.Body()
	}
	return body, nil
}
