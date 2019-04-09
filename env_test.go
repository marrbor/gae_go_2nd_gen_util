/*
 * Copyright (c) 2017 Genetec corporation
 * -*- coding:utf-8 -*-
 *
 * ファイルの説明
 *
 */
package gae_go_2nd_gen_util

import (
	"os"
	"testing"
)

func TestGetAppID(t *testing.T) {
	if err := os.Setenv("GAE_APPLICATION", "test"); err != nil {
		t.Fatalf("error")
	}

	if GetAppID() != "test" {
		t.Fatalf("error")
	}
}

func TestGetDeploymentID(t *testing.T) {
	if err := os.Setenv("GAE_DEPLOYMENT_ID", "id123"); err != nil {
		t.Fatalf("error")
	}
	if GetDeploymentID() != "id123" {
		t.Fatalf("error")
	}
}

func TestGetGAEEnv(t *testing.T) {
	if err := os.Setenv("GAE_ENV", "product"); err != nil {
		t.Fatalf("error")
	}
	if GetGAEEnv() != "product" {
		t.Fatalf("error")
	}
}

func TestGetInstanceID(t *testing.T) {
	if err := os.Setenv("GAE_INSTANCE", "instId123"); err != nil {
		t.Fatalf("error")
	}
	if GetInstanceID() != "instId123" {
		t.Fatalf("error")
	}
}

func TestGetMemoryMB(t *testing.T) {
	if err := os.Setenv("GAE_MEMORY_MB", "100"); err != nil {
		t.Fatalf("error")
	}
	if GetMemoryMB() != "100" {
		t.Fatalf("error")
	}
}

func TestGetRuntime(t *testing.T) {
	if err := os.Setenv("GAE_RUNTIME", "python"); err != nil {
		t.Fatalf("error")
	}
	if GetRuntime() != "python" {
		t.Fatalf("error")
	}
}

func TestGetService(t *testing.T) {
	if err := os.Setenv("GAE_SERVICE", "background"); err != nil {
		t.Fatalf("error")
	}
	if GetService() != "background" {
		t.Fatalf("error")
	}
}

func TestGetGAEVersion(t *testing.T) {
	if err := os.Setenv("GAE_VERSION", "go1.11"); err != nil {
		t.Fatalf("error")
	}
	if GetGAEVersion() != "go1.11" {
		t.Fatalf("error")
	}
}

func TestGetProjectID(t *testing.T) {
	if err := os.Setenv("GOOGLE_CLOUD_PROJECT", "project123"); err != nil {
		t.Fatalf("error")
	}
	if GetProjectID() != "project123" {
		t.Fatalf("error")
	}
}

func TestGetNodeENV(t *testing.T) {
	if err := os.Setenv("NODE_ENV", "project123"); err != nil {
		t.Fatalf("error")
	}
	if GetNodeENV() != "project123" {
		t.Fatalf("error")
	}
}

func TestGetPort(t *testing.T) {
	if err := os.Setenv("PORT", "100"); err != nil {
		t.Fatalf("error")
	}
	port, err := GetPort()
	if err != nil {
		t.Fatalf("error")
	}

	if port != 100 {
		t.Fatalf("error")
	}
	if err := os.Setenv("PORT", "abc"); err != nil {
		t.Fatalf("error")
	}
	_, err2 := GetPort()
	if err2 == nil {
		t.Fatalf("error")
	}
	x := err2.Error()
	if x != "strconv.ParseInt: parsing \"abc\": invalid syntax" {
		t.Fatalf("error: %s", x)
	}
}
