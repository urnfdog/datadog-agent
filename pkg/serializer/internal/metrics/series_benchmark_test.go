// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package metrics

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/metrics"
	"github.com/DataDog/datadog-agent/pkg/tagset"
)

func benchmarkPopulateDeviceField(numberOfTags int, b *testing.B) {
	tags := make([]string, 0, numberOfTags+1)
	for i := 0; i < numberOfTags; i++ {
		tags = append(tags, "some:tag")
	}
	t := tagset.CompositeTagsFromSlice(append(tags, "device:test"))

	serie := &metrics.Serie{
		Tags: t,
	}
	series := []*metrics.Serie{serie}

	for n := 0; n < b.N; n++ {
		serie.Tags = t
		for _, serie := range series {
			populateDeviceField(serie)
		}
	}
}

func BenchmarkPopulateDeviceField1(b *testing.B)  { benchmarkPopulateDeviceField(1, b) }
func BenchmarkPopulateDeviceField2(b *testing.B)  { benchmarkPopulateDeviceField(2, b) }
func BenchmarkPopulateDeviceField3(b *testing.B)  { benchmarkPopulateDeviceField(3, b) }
func BenchmarkPopulateDeviceField10(b *testing.B) { benchmarkPopulateDeviceField(10, b) }
func BenchmarkPopulateDeviceField20(b *testing.B) { benchmarkPopulateDeviceField(20, b) }
func BenchmarkPopulateDeviceField40(b *testing.B) { benchmarkPopulateDeviceField(40, b) }
