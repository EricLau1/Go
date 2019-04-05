// +build unit

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SumResult struct {
	x        int
	y        int
	expected int
}

type Results []SumResult

var sumResults = Results{
	SumResult{1, 1, 2},
	SumResult{1, -1, 0},
	SumResult{-1, -1, -2},
}

func TestSum(t *testing.T) {
	for _, test := range sumResults {
		if r := Sum(test.x, test.y); r != test.expected {
			t.Errorf("Test Failed: inputs sum(%d, %d), expected %d, received %d", test.x, test.y, test.expected, r)
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "Hello, Developers!" {
		t.Fatal("String contents do not match expected")
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"OK\"}")
	}

	req := httptest.NewRequest("GET", "https://ericlau1.github.io", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.StatusCode != 200 {
		t.Fatal("Status Code not OK")
	}
}
