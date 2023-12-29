package httpclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/matDobek/gov--attendance-check/internal/cache"
	"github.com/matDobek/gov--attendance-check/internal/logger"
	"github.com/matDobek/gov--attendance-check/internal/predicates"
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

var (
	ErrStatusInfo        = errors.New("received unwated status: info")
	ErrStatusRedirect    = errors.New("received unwated status: redirect")
	ErrStatusClientError = errors.New("received unwated status: client error")
	ErrStatusServerError = errors.New("received unwated status: server error")
)

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
		return nil, ConnectionError{err}
	}

	// naive prevention for rejecting bot requests
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0")

	time.Sleep(2 * time.Second)
	logger.Info("GET: %s", url)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, ConnectionError{err}
	}
	defer resp.Body.Close()

	if predicates.Between(resp.StatusCode, 100, 199) {
		return []byte{}, ErrStatusInfo
	} else if predicates.Between(resp.StatusCode, 200, 299) {
		// much success
	} else if predicates.Between(resp.StatusCode, 300, 399) {
		return []byte{}, ErrStatusRedirect
	} else if predicates.Between(resp.StatusCode, 400, 499) {
		return []byte{}, ErrStatusClientError
	} else {
		return []byte{}, ErrStatusServerError
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
