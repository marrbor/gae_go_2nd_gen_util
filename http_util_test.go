/*
 * Copyright (c) 2017 Genetec corporation
 * -*- coding:utf-8 -*-
 *
 * ファイルの説明
 *
 */
package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type (
	TestApiBody struct {
		Int  int64    `json:"i_64"`
		Str  string   `json:"str"`
		AInt []int64  `json:"a_int"`
		AStr []string `json:"a_str"`
	}
)

func testErrors(t *testing.T, code int, msg string, f func(w http.ResponseWriter, err2 error)) {

	ErrorStr := "error"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/err" {
			f(w, fmt.Errorf(ErrorStr))
			return
		}
		f(w, nil)
	}))
	defer ts.Close()

	// リクエストする
	req, err := http.NewRequest(
		"GET",
		ts.URL,
		nil,
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != code {
		t.Fatalf("error.")
	}

	// リクエストする
	req2, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/err", ts.URL),
		nil,
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp2.Body.Close()

	if resp.StatusCode != code {
		t.Fatalf("error.")
	}

	bb, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	bs := string(bb)
	if bs != fmt.Sprintf("%s %s\n", msg, ErrorStr) {
		t.Fatalf("error.(%s)", bs)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type errReader int

var ErrorStr = "test error"

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New(ErrorStr)
}

func TestRequestToParams(t *testing.T) {

	var x TestApiBody
	r := httptest.NewRequest(http.MethodPost, "/test", errReader(0))
	e := RequestToParams(r, &x)
	if e == nil {
		t.Fatalf("error")
	}

	if e.Error() != ErrorStr {
		t.Fatalf("error")
	}

	td := TestApiBody{
		Int:  -1,
		Str:  "test",
		AInt: []int64{1, 2, 3},
		AStr: []string{"one", "two", "three"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req TestApiBody
		if err := RequestToParams(r, &req); err != nil {
			t.Fatalf("%+v", err)
		}

		if req.Int != td.Int {
			t.Fatalf("error. %d - %d ", td.Int, req.Int)
		}

		if req.Str != td.Str {
			t.Fatalf("error. %s - %s ", td.Str, req.Str)
		}

		for i := range req.AInt {
			if req.AInt[i] != td.AInt[i] {
				t.Fatalf("error. %d - %d", req.AInt[i], td.AInt[i])
			}
		}

		for i := range req.AStr {
			if req.AStr[i] != td.AStr[i] {
				t.Fatalf("error. %s - %s", req.AStr[i], td.AStr[i])
			}
		}

	}))
	defer ts.Close()

	// リクエストする
	body, err := json.Marshal(td)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	req, err := http.NewRequest(
		"POST",
		ts.URL,
		bytes.NewReader(body),
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp.Body.Close()

}

func TestOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		OK(w)
	}))
	defer ts.Close()

	// リクエストする
	req, err := http.NewRequest(
		"GET",
		ts.URL,
		nil,
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("error.")
	}

	if resp.Status != "200 OK" {
		t.Fatalf("error.(%s)", resp.Status)
	}
}

func TestJSONResponse(t *testing.T) {

	td := TestApiBody{
		Int:  -1,
		Str:  "test",
		AInt: []int64{1, 2, 3},
		AStr: []string{"one", "two", "three"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/err" {
			// nil
			JSONResponse(w, nil)
			return
		}

		if r.URL.Path == "/err2" {
			// JSON化できないデータ
			JSONResponse(w, map[string]interface{}{"foo": make(chan int)})
			return
		}

		var req TestApiBody
		if err := RequestToParams(r, &req); err != nil {
			t.Fatalf("%+v", err)
		}

		req.Int += 1
		req.Str += "a"

		for i := range req.AInt {
			req.AInt[i] += 1
		}

		for i := range req.AStr {
			req.AStr[i] += "a"
		}
		JSONResponse(w, req)
	}))
	defer ts.Close()

	// リクエストする
	body, err := json.Marshal(td)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	req, err := http.NewRequest(
		"POST",
		ts.URL,
		bytes.NewReader(body),
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp.Body.Close()

	var receive TestApiBody
	if err := json.NewDecoder(resp.Body).Decode(&receive); err != nil {
		t.Fatalf("%+v", err)
	}

	if receive.Int != td.Int+1 {
		t.Fatalf("error. %d - %d ", td.Int, receive.Int)
	}

	if receive.Str != td.Str+"a" {
		t.Fatalf("error. %s - %s ", td.Str, receive.Str)
	}

	for i := range receive.AInt {
		if receive.AInt[i] != td.AInt[i]+1 {
			t.Fatalf("error. %d - %d", receive.AInt[i], td.AInt[i])
		}
	}

	for i := range receive.AStr {
		if receive.AStr[i] != td.AStr[i]+"a" {
			t.Fatalf("error. %s - %s", receive.AStr[i], td.AStr[i])
		}
	}

	req2, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/err", ts.URL),
		nil,
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if resp2.StatusCode != 200 {
		t.Fatalf("error")
	}

	rb, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer resp2.Body.Close()

	rs := string(rb)
	if 0 < len(rs) {
		t.Fatalf("%d: %s %+v", len(rs), rs, rb)
	}

	req3, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/err2", ts.URL),
		nil,
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	resp3, err := client.Do(req3)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if resp3.StatusCode != 500 {
		t.Fatalf("error")
	}

	if resp3.Status != "500 Internal Server Error" {
		t.Fatalf("error(%s)", resp3.Status)
	}
}

func TestBadRequest(t *testing.T) {
	testErrors(t, 400, "Bad Request", BadRequest)
}

func TestUnauthorized(t *testing.T) {
	testErrors(t, 401, "Unauthorized", Unauthorized)
}

func TestForbidden(t *testing.T) {
	testErrors(t, 403, "Forbidden", Forbidden)
}

func TestNotFound(t *testing.T) {
	testErrors(t, 404, "Not Found", NotFound)
}

func TestMethodNotAllowed(t *testing.T) {
	testErrors(t, 405, "Method Not Allowed", MethodNotAllowed)
}

func TestInternalServerError(t *testing.T) {
	testErrors(t, 500, "Internal Server Error", InternalServerError)
}
