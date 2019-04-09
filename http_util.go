/*
 * Copyright (c) 2019 Masahide Matsumoto
 * -*- coding:utf-8 -*-
 *
 * Web API を構築するためのユーティリティ
 *
 */
package util

import (
	"encoding/json"
	"net/http"
)

// Request

// RequestToParams は、request にあるボディー（JSON）を params でポインタ渡しされた go struct に変換します。
func RequestToParams(r *http.Request, params interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(&params)
}

// Response

// OK は、200 OK のヘッダだけを返します。
func OK(w http.ResponseWriter) {
	w.WriteHeader(200)
}

// JSONResponse は、200 OKで JSON オブジェクトを返します。エンコードに失敗したら 500 エラーを返します。
func JSONResponse(w http.ResponseWriter, data interface{}) {
	if data == nil {
		OK(w)
		return
	}
	j, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	// レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// エラーを返す
func errResponse(w http.ResponseWriter, code int, err error) {
	msg := http.StatusText(code)
	if err != nil {
		msg = msg + " " + err.Error()
	}
	http.Error(w, msg, code)
}

// BadRequest は、400エラーを返します。
func BadRequest(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusBadRequest, err)
}

// Unauthorized は、401エラーを返します。
func Unauthorized(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusUnauthorized, err)
}

// Forbidden は、403エラーを返します。
func Forbidden(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusForbidden, err)
}

// NotFound は、404エラーを返します。
func NotFound(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusNotFound, err)
}

// MethodNotAllowed は、405エラーを返します。
func MethodNotAllowed(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusMethodNotAllowed, err)
}

// InternalServerError は、500エラーを返します。
func InternalServerError(w http.ResponseWriter, err error) {
	errResponse(w, http.StatusInternalServerError, err)
}
