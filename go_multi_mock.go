
import (
	"net/http"
	"net/http/httptest"
)

type MockResponse struct {
	Status int
	Body   []byte
}

type Responses []*MockResponse

func (r *Responses) removeFirstElement() {
	newResponses := *r
	newResponses = append(newResponses[:0], newResponses[1:]...)
	*r = newResponses
}

func StartMockBarometer(responses Responses) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		response := responses[0].Status
		body := responses[0].Body
		responses.removeFirstElement()
		res.WriteHeader(response)
		res.Write(body)
	}))
}
