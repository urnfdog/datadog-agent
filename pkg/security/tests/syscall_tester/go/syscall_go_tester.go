// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// +build syscalltesters

package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"

	manager "github.com/DataDog/ebpf-manager"
)

var (
	bpfLoad  bool
	bpfClone bool
)

// probe.c
//
// struct bpf_map_def SEC("maps/cache") cache = {
//      .type = BPF_MAP_TYPE_HASH,
//      .key_size = sizeof(u32),
//      .value_size = sizeof(u32),
//      .max_entries = 10,
//  };
//
//  SEC("kprobe/vfs_open")
//  int kprobe_vfs_open(void *ctx)
//  {
//      u32 key = 1;
//      u32 *value = bpf_map_lookup_elem(&cache, &key);
//      if (value == 0) {
//          bpf_printk("map entry 1 is empty!\n");
//      }
//      bpf_printk("hello world!\n");
//      return 0;
//  }
//
//  char _license[] SEC("license") = "GPL";
//  __u32 _version SEC("version") = 0xFFFFFFFE;

var ebpfProbe = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x4f\x6b\xd4\x40\x14\x7f\x93\x5d\xed\x36\xd6\x52\xbd\x54\xc2\x1e\xb2\xe0\xc1\x53\xda\x2d\xea\x4d\xa8\x05\xf5\xe0\x1e\x8a\xa0\x37\x09\x69\x3a\x25\x4b\xf3\x67\x49\x42\x75\xad\xe2\x45\xef\x82\xf8\x01\xf4\xe2\xd5\x8b\xec\x45\x70\x3f\x82\x1f\xa1\x47\x0f\x1e\x2a\x88\x5e\x8a\x23\x6f\xf2\x26\x1b\xa6\x49\xeb\x07\xe8\x0f\x92\x99\xf7\xcb\xbc\x7f\xf3\x9b\x24\x2f\xef\x0c\xee\x1a\x8c\x81\x02\x83\x3f\x30\xb3\x66\x08\xda\xb3\xf9\x3a\xdd\x2f\x02\x83\x09\x03\xb9\xde\xb7\x8e\x04\xb2\x5f\x3f\x14\x6b\xe6\x0c\x80\x23\x21\xc4\x15\x2d\xd8\x6b\x28\xd6\x3f\x84\x25\x69\x4f\xe8\x79\x66\xfd\x16\xca\xee\x99\x00\xbb\xd6\xaf\xd2\x8e\x46\xf9\xd8\xb7\x0e\xa5\x8d\xf1\xc6\x76\xdf\xc6\xf9\x30\xb3\xf9\xbe\xf5\xbd\xe4\x23\x6f\x24\x79\x1e\xe7\xe9\xbe\x75\x50\xd4\xf3\x9e\xea\x61\x00\x07\x42\x88\x89\x01\xb0\x4c\x75\x9c\xa7\xf8\x26\x60\xbe\x1f\x65\xbe\x34\xdc\xee\xf9\x95\xb8\x01\x0f\x43\x9c\x27\xf6\x93\xe4\xa4\xb8\x8b\xd5\xb8\xd4\xef\xbb\x72\x5f\x01\xda\x74\x99\x35\xfb\x7b\x86\x02\xa8\x21\xea\x87\x1a\xa3\xbe\xa8\x7d\xcf\x94\x12\xe0\xf6\xa3\x34\x26\xdc\xdb\x1c\xe0\xda\xbf\x42\x88\x25\xf2\x63\xcf\x1e\x40\xe7\xf9\x05\xb6\x80\x9a\xd1\xa5\x70\xad\xee\x40\x6b\x78\x24\xb5\xf9\x29\x74\xfe\x85\xbc\xb7\x60\x5a\xe3\xd3\x82\xd6\x31\xee\x3a\x00\x5c\x82\x4e\x69\xab\xd7\x66\x4e\xf2\xf3\xc7\xf8\x4d\xc9\x9f\x2b\xf9\x2e\xd5\x8a\x7d\x5d\xae\xc4\x57\x3d\xd8\xaa\x5f\x3a\x67\xdd\x8a\x2d\x57\x3b\x39\x7f\x9a\x83\xbb\xc7\xd3\x6c\x98\xc4\xb0\x3b\x4a\x93\x2d\xee\xee\xed\x64\x6e\x32\xe2\x31\x38\x29\x0f\x0b\x6e\xa5\xe4\xdc\x70\xe8\xf3\x38\xe3\xf2\xa1\xc3\x03\x77\x27\xf5\x22\x8e\x52\x64\x2b\xbe\xe7\x07\x38\x1d\xc6\x8e\x0f\x4e\x96\xa7\xb9\xb7\x05\x4e\x36\x8e\xe4\x98\x26\xdb\x5e\xee\x21\xdd\x77\xfa\x37\x61\xb0\xb1\xb1\xea\xae\x69\xf4\xe9\x7b\xff\x3f\x78\xac\xfa\xd3\xf0\x8d\xc8\xb7\x1a\xaf\x4b\xce\x2a\x7b\x56\xc5\x7a\x43\xbe\xb6\x66\x5f\x3d\xc5\x5f\x3f\x63\x1d\x6d\x1d\xea\x36\x5f\x93\xe7\x13\xd5\xaf\xce\xf1\x02\xf5\xa9\xfc\x15\x7f\xbf\xaa\x71\x05\x01\xe5\xed\xb2\x93\xeb\x8f\xc8\x7f\x4d\xe3\xdf\x18\xc5\xb8\xac\xf1\xe5\x7b\x45\xe3\xab\x06\xff\x8f\xe4\xbf\xa8\xf1\x4c\x1b\x6f\x34\xd4\xff\xd9\xa8\xaf\x57\xd7\xaf\xd3\xe0\xff\xa5\xc1\x5f\xb7\x6f\x93\xbf\xa1\xf1\x53\x22\x56\x6b\xf2\x55\x71\xab\x41\xbf\x69\x8d\x7e\x66\x8d\x7e\xbc\x26\x37\xe2\x90\x48\xf5\x7d\x51\xff\x09\xe5\xaf\xbe\x63\xff\x02\x00\x00\xff\xff\x44\xea\xbc\x91\xa8\x07\x00\x00")

func bindataRead(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to read ebpfProbe: %w", err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("failed to read ebpfProbe: %w", err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BPFClone(m *manager.Manager) error {
	if _, err := m.CloneMap("cache", "cache_clone", manager.MapOptions{}); err != nil {
		return fmt.Errorf("couldn't clone 'cache' map: %w", err)
	}
	return nil
}

func BPFLoad() error {
	m := &manager.Manager{
		Probes: []*manager.Probe{
			{
				ProbeIdentificationPair: manager.ProbeIdentificationPair{
					UID:          "MyVFSOpen",
					EBPFSection:  "kprobe/vfs_open",
					EBPFFuncName: "kprobe_vfs_open",
				},
			},
		},
		Maps: []*manager.Map{
			{
				Name: "cache",
			},
		},
	}
	defer func() {
		_ = m.Stop(manager.CleanAll)
	}()

	rawProbe, err := bindataRead(ebpfProbe)
	if err != nil {
		return err
	}

	if err = m.Init(bytes.NewReader(rawProbe)); err != nil {
		return fmt.Errorf("failed to initialize manager: %w", err)
	}

	if err = BPFClone(m); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.BoolVar(&bpfLoad, "load-bpf", false, "load the eBPF progams")
	flag.BoolVar(&bpfClone, "clone-bpf", false, "clone maps")

	flag.Parse()

	if bpfLoad {
		if err := BPFLoad(); err != nil {
			panic(err)
		}
	}
}
