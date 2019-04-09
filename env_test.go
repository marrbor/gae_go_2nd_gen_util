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
	os.Setenv("GAE_APPLICATION", "test")
	if GetAppID() != "test" {
		t.Fatalf("error")
	}
}

func TestGetDeploymentID(t *testing.T) {
	os.Setenv("GAE_DEPLOYMENT_ID", "id123")
	if GetDeploymentID() != "id123" {
		t.Fatalf("error")
	}
}

func TestGetGAEEnv(t *testing.T) {
	os.Setenv("GAE_ENV", "product")
	if GetGAEEnv() != "product" {
		t.Fatalf("error")
	}
}

func TestGetInstanceID(t *testing.T) {
	os.Setenv("GAE_INSTANCE", "instId123")
	if GetInstanceID() != "instId123" {
		t.Fatalf("error")
	}
}

func TestGetMemoryMB(t *testing.T) {
	os.Setenv("GAE_MEMORY_MB", "100")
	if GetMemoryMB() != "100" {
		t.Fatalf("error")
	}
}

func TestGetRuntime(t *testing.T) {
	os.Setenv("GAE_RUNTIME", "python")
	if GetRuntime() != "python" {
		t.Fatalf("error")
	}
}

func TestGetService(t *testing.T) {
	os.Setenv("GAE_SERVICE", "background")
	if GetService() != "background" {
		t.Fatalf("error")
	}
}

func TestGetGAEVersion(t *testing.T) {
	os.Setenv("GAE_VERSION", "go1.11")
	if GetGAEVersion() != "go1.11" {
		t.Fatalf("error")
	}
}

func TestGetProjectID(t *testing.T) {
	os.Setenv("GOOGLE_CLOUD_PROJECT", "project123")
	if GetProjectID() != "project123" {
		t.Fatalf("error")
	}
}

func TestGetNodeENV(t *testing.T) {
	os.Setenv("NODE_ENV", "project123")
	if GetNodeENV() != "project123" {
		t.Fatalf("error")
	}
}

func TestGetPort(t *testing.T) {
	os.Setenv("PORT", "100")
	port, err := GetPort()
	if err != nil {
		t.Fatalf("error")
	}

	if port != 100 {
		t.Fatalf("error")
	}
	os.Setenv("PORT", "abc")
	_, err2 := GetPort()
	if err2 == nil {
		t.Fatalf("error")
	}
	x := err2.Error()
	if x != "strconv.ParseInt: parsing \"abc\": invalid syntax" {
		t.Fatalf("error: %s", x)
	}
}
