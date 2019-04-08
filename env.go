/*
 * Copyright (c) 2019 Masahide Matsumoto
 * -*- coding:utf-8 -*-
 *
 * GAE/Go 環境変数に関するユーティリティ
 * https://cloud.google.com/appengine/docs/standard/go111/runtime
 *
 */
package gae_go_2nd_gen_util

import (
	"os"
	"strconv"
)

// GetAppID は、アプリケーション ID を返します。
func GetAppID() string {
	return os.Getenv("GAE_APPLICATION")
}

// GetDeploymentID は、デプロイメント ID を返します。
func GetDeploymentID() string {
	return os.Getenv("GAE_DEPLOYMENT_ID")
}

// GetGAEEnv は、App Engine の環境を返します
func GetGAEEnv() string {
	return os.Getenv("GAE_ENV")
}

// GetInstanceID は、動作しているインスタンスIDを返します。
func GetInstanceID() string {
	return os.Getenv("GAE_INSTANCE")
}

// GetMemoryMB は、使用可能なメモリ量を返します。単位は MB です。
func GetMemoryMB() string {
	return os.Getenv("GAE_MEMORY_MB")
}

// GetRuntime は、app.yaml で指定されたランタイムを返します。
func GetRuntime() string {
	return os.Getenv("GAE_RUNTIME")
}

// GetService は、app.yaml で示されたサービス名を返します。
func GetService() string {
	return os.Getenv("GAE_SERVICE")
}

// GetGAEVersion は、app.yaml で示された現在のバージョンを返します。
func GetGAEVersion() string {
	return os.Getenv("GAE_VERSION")
}

// GetProjectID は、app.yaml で示された現在のバージョンを返します。
func GetProjectID() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

// GetNodeENV は
func GetNodeENV() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

// GetPort は待ち受けているポートを返します。
func GetPort() (int64, error) {
	return strconv.ParseInt(os.Getenv("PORT"), 10, 64)
}
