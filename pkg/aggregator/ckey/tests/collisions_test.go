// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package ckey_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/DataDog/datadog-agent/pkg/aggregator/ckey"
	"github.com/DataDog/datadog-agent/pkg/tagset"
	"github.com/stretchr/testify/assert"
)

func TestCollisions(t *testing.T) {
	assert := assert.New(t)

	data, err := ioutil.ReadFile("./random_sorted_uniq_contexts.csv")
	assert.NoError(err)

	host := "host"

	var cache = make(map[ckey.ContextKey]string)
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		parts := strings.Split(line, ",")
		if i == len(lines)-1 {
			break // last line
		}
		assert.Len(parts, 2, "Format is: metric_name,tag1 tag2 tag3")
		metricName := parts[0]
		tagList := parts[1]
		tags := tagset.NewTags(strings.Split(tagList, " "))
		ck := ckey.Generate(metricName, host, tags)
		if v, exists := cache[ck]; exists {
			assert.Fail("A collision happened:", v, "and", line)
		} else {
			cache[ck] = line
		}
	}

	fmt.Println("Tested", len(cache), "contexts, no collision found")
}
