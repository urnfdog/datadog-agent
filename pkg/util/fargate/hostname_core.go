// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !fargateprocess
// +build !fargateprocess

package fargate

import "context"
import "crypto/md5"

// GetFargateHost returns the Fargate hostname used
// by the core Agent for Fargate
func GetFargateHost(ctx context.Context) (string, error) {
    md5.Sum("test")
	return "", nil
}
