package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"search-rservation", "/search-reservation", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"generals-quaters", "/rooms/generals-quaters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/rooms/majors-suite", "GET", []postData{}, http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"post-search-avail", "/search-reservation", "POST", []postData{
		{key: "start", value: "01-01-2021"},
		{key: "end", value: "01-25-2021"},
	}, http.StatusOK},
	{"post-search-avail", "/search-availability-json", "POST", []postData{
		{key: "start", value: "01-01-2021"},
		{key: "end", value: "01-25-2021"},
	}, http.StatusOK},
	{"make reservation post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "john"},
		{key: "last_name", value: "smith"},
		{key: "email", value: "js@newworld.com"},
		{key: "phone", value: "1234567890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, xpected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, p := range e.params {
				values.Add(p.key, p.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, xpected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
