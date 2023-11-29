package httpclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/cache"
	"github.com/matDobek/gov--attendance-check/internal/logger"
)

//===================
// Defaults
//===================

var DefaultOptions = ClientOptions{
	Timeout: 5 * time.Second,
}

//===================
// Types
//===================

type HttpClient struct {
	client *http.Client
	cache  cache.Cache
}

type ClientOptions struct {
	Timeout time.Duration
}

//===================
// Errors
//===================

type ConnectionError struct {
	err error
}

func (e ConnectionError) Error() string {
	return fmt.Sprintf("ConnectionError: %s", e.err.Error())
}

func (e ConnectionError) Is(target error) bool {
	_, ok := target.(ConnectionError)
	return ok
}

//===================
// Functions
//===================

//
//
//

func New(cache cache.Cache) *HttpClient {
	options := DefaultOptions

	return ConfigurableNew(cache, options)
}

//
//
//

func ConfigurableNew(cache cache.Cache, options ClientOptions) *HttpClient {
	return &HttpClient{
		cache: cache,
		client: &http.Client{
			Timeout: options.Timeout,
		},
	}
}

//
//
//

func (c *HttpClient) CachedGet(url string) ([]byte, error) {
	cachedResponse, err := c.cache.Get(url)

	switch {
	case err == nil:
		return []byte(cachedResponse), nil
	case errors.Is(err, cache.ErrCouldNotRead):
		response, err := c.Get(url)
		if err != nil {
			return []byte{}, err
		}

		err = c.cache.Put(url, string(response))
		if err != nil {
			logger.Error(err)
		}

		return response, nil
	default:
		return []byte{}, err
	}
}

//
//
//

func (c *HttpClient) Get(url string) (respBody []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return failure[[]byte](err)
	}

	// naive prevention for rejecting bot requests
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0")

	resp, err := c.client.Do(req)
	if err != nil {
		return failure[[]byte](err)
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return failure[[]byte](err)
	}

	return success(response)
}

func failure[V any](err error) (V, error) {
	var val V
	return val, ConnectionError{err}
}

func success[V any](val V) (V, error) {
	return val, nil
}
