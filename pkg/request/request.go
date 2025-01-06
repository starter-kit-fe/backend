package request

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Common errors
var (
	ErrEmptyMethod     = errors.New("method cannot be empty")
	ErrEmptyURL        = errors.New("URL cannot be empty")
	ErrInvalidURL      = errors.New("invalid URL")
	ErrBodyTooLarge    = errors.New("response body exceeds maximum size")
	ErrRequestTimeout  = errors.New("request timeout")
	ErrContextCanceled = errors.New("request context canceled")
)

const (
	defaultTimeout     = 30 * time.Second
	defaultMaxBodySize = 10 << 20 // 10MB
	defaultRetries     = 2
	defaultBufferSize  = 32 * 1024 // 32KB for buffer pool
)

// ResponseHandler defines how to handle HTTP responses
type ResponseHandler func(*http.Response) error

// HttpClient represents an HTTP client with configuration
type HttpClient struct {
	client      *http.Client
	timeout     time.Duration
	userAgent   string
	maxBodySize int64
	retries     int
	bufferPool  sync.Pool
	hooks       []ResponseHandler
}

// Option defines the method to customize HttpClient
type Option func(*HttpClient)

// WithTimeout sets custom timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *HttpClient) {
		c.timeout = timeout
	}
}

// WithUserAgent sets custom user-agent
func WithUserAgent(userAgent string) Option {
	return func(c *HttpClient) {
		c.userAgent = userAgent
	}
}

// WithMaxBodySize sets maximum response body size
func WithMaxBodySize(size int64) Option {
	return func(c *HttpClient) {
		c.maxBodySize = size
	}
}

// WithRetries sets the number of retry attempts
func WithRetries(retries int) Option {
	return func(c *HttpClient) {
		c.retries = retries
	}
}

// WithResponseHook adds a response hook
func WithResponseHook(hook ResponseHandler) Option {
	return func(c *HttpClient) {
		c.hooks = append(c.hooks, hook)
	}
}

// WithCustomTransport sets a custom transport
func WithCustomTransport(transport http.RoundTripper) Option {
	return func(c *HttpClient) {
		c.client.Transport = transport
	}
}

// NewHttpClient creates a new HttpClient instance with options
func NewHttpClient(opts ...Option) *HttpClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxConnsPerHost:       100,
	}

	client := &HttpClient{
		timeout:     defaultTimeout,
		maxBodySize: defaultMaxBodySize,
		retries:     defaultRetries,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, defaultBufferSize))
			},
		},
	}

	client.client = &http.Client{
		Transport: transport,
		Timeout:   client.timeout,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Request encapsulates HTTP request configuration
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    interface{}
	Context context.Context // Embed context directly in request
}

// Validate checks if the request configuration is valid
func (r *Request) Validate() error {
	if r.Method == "" {
		return ErrEmptyMethod
	}
	if r.URL == "" {
		return ErrEmptyURL
	}
	if _, err := url.Parse(r.URL); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidURL, err)
	}
	return nil
}

// DoWithBody executes HTTP request and returns raw response
func (hc *HttpClient) DoWithBody(reqConfig *Request) (*http.Response, error) {
	if err := reqConfig.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request config: %w", err)
	}

	ctx := reqConfig.Context
	if ctx == nil {
		ctx = context.Background()
	}

	var lastErr error
	for attempt := 0; attempt <= hc.retries; attempt++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("%w: %v", ErrContextCanceled, ctx.Err())
		default:
		}

		resp, err := hc.executeRequest(ctx, reqConfig)
		if err == nil {
			return resp, nil
		}

		lastErr = err
		if !hc.shouldRetry(err) {
			break
		}

		// Exponential backoff
		if attempt < hc.retries {
			time.Sleep(time.Duration(1<<uint(attempt)) * time.Second)
		}
	}

	return nil, fmt.Errorf("all attempts failed: %w", lastErr)
}

// executeRequest performs a single HTTP request
func (hc *HttpClient) executeRequest(ctx context.Context, reqConfig *Request) (*http.Response, error) {
	body, err := hc.prepareRequestBody(reqConfig.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, reqConfig.Method, reqConfig.URL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	hc.setRequestHeaders(req, reqConfig.Headers)

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Execute response hooks
	for _, hook := range hc.hooks {
		if err := hook(resp); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("response hook failed: %w", err)
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		buf := hc.bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		defer hc.bufferPool.Put(buf)

		_, _ = io.CopyN(buf, resp.Body, 1024) // Read first 1KB of error response
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, buf.String())
	}

	return resp, nil
}

// Do executes HTTP request and returns response body bytes
func (hc *HttpClient) Do(reqConfig *Request) ([]byte, error) {
	resp, err := hc.DoWithBody(reqConfig)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Use buffer pool for reading response
	buf := hc.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer hc.bufferPool.Put(buf)

	written, err := io.CopyN(buf, resp.Body, hc.maxBodySize+1)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if written > hc.maxBodySize {
		return nil, ErrBodyTooLarge
	}

	return buf.Bytes(), nil
}

// shouldRetry determines if the request should be retried based on the error
func (hc *HttpClient) shouldRetry(err error) bool {
	if err == nil {
		return false
	}

	// Retry on network errors and server errors (5xx)
	var netErr net.Error
	if errors.As(err, &netErr) {
		return true
	}

	if errors.Is(err, io.EOF) {
		return true
	}

	var urlErr *url.Error
	if errors.As(err, &urlErr) && urlErr.Timeout() {
		return true
	}

	return false
}

// prepareRequestBody converts the body interface to io.Reader
func (hc *HttpClient) prepareRequestBody(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	switch b := body.(type) {
	case string:
		return bytes.NewBufferString(b), nil
	case []byte:
		return bytes.NewBuffer(b), nil
	case url.Values:
		return bytes.NewBufferString(b.Encode()), nil
	case io.Reader:
		return b, nil
	default:
		buf := hc.bufferPool.Get().(*bytes.Buffer)
		buf.Reset()

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			hc.bufferPool.Put(buf)
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}

		return buf, nil
	}
}

// setRequestHeaders sets the request headers
func (hc *HttpClient) setRequestHeaders(req *http.Request, headers map[string]string) {
	if hc.userAgent != "" {
		req.Header.Set("User-Agent", hc.userAgent)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
