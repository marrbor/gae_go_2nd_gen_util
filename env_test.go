/*
 * Copyright (c) 2017 Genetec corporation
 * -*- coding:utf-8 -*-
 *
 * ファイルの説明
 *
 */
package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppID(t *testing.T) {
	err := os.Setenv("GAE_APPLICATION", "test")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "test", GetAppID(), "error")
}

func TestGetDeploymentID(t *testing.T) {
	err := os.Setenv("GAE_DEPLOYMENT_ID", "id123")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "id123", GetDeploymentID(), "error")
}

func TestGetGAEEnv(t *testing.T) {
	err := os.Setenv("GAE_ENV", "product")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "product", GetGAEEnv())
}

func TestGetInstanceID(t *testing.T) {
	err := os.Setenv("GAE_INSTANCE", "instId123")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "instId123", GetInstanceID())
}

func TestGetMemoryMB(t *testing.T) {
	err := os.Setenv("GAE_MEMORY_MB", "100")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "100", GetMemoryMB())
}

func TestGetRuntime(t *testing.T) {
	err := os.Setenv("GAE_RUNTIME", "python")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "python", GetRuntime())
}

func TestGetService(t *testing.T) {
	err := os.Setenv("GAE_SERVICE", "background")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "background", GetService())
}

func TestGetGAEVersion(t *testing.T) {
	err := os.Setenv("GAE_VERSION", "go1.11")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "go1.11", GetGAEVersion())
}

func TestGetProjectID(t *testing.T) {
	err := os.Setenv("GOOGLE_CLOUD_PROJECT", "project123")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "project123", GetProjectID())
}

func TestGetNodeENV(t *testing.T) {
	err := os.Setenv("NODE_ENV", "project123")
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, "project123", GetNodeENV())
}

func TestGetPort(t *testing.T) {
	err := os.Setenv("PORT", "100")
	assert.NoError(t, err, "unexpected %+v", err)

	port, err := GetPort()
	assert.NoError(t, err, "unexpected %+v", err)
	assert.Equal(t, 100, port)
	err = os.Setenv("PORT", "abc")
	assert.NoError(t, err, "unexpected %+v", err)

	_, err = GetPort()
	assert.EqualError(t, err, "strconv.ParseInt: parsing \"abc\": invalid syntax")
}
