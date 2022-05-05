package chapter13

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTransport struct {
	req **http.Request
	res *http.Response
	err error
}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	*(t.req) = req

	return t.res, t.err
}

var _ http.RoundTripper = &testTransport{}

func newTransport(req **http.Request, res *http.Response, err error) http.RoundTripper {
	return &testTransport{
		req: req,
		res: res,
		err: err,
	}
}

func testHTTPRequest(t *testing.T) {
	var req *http.Request
	res := httptest.NewRecorder()
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.WriteString(`{"ranking": ["Back to the Future", "Rambo"]}`)

	c := http.Client{
		Transport: newTransport(&req, res.Result(), nil),
	}

	r, err := c.Get("http://example.com/movies/1985")

	assert.Nil(t, err)
	assert.Equal(t, 200, r.StatusCode)
}
