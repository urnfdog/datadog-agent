#ifndef __HTTP_MAPS_H
#define __HTTP_MAPS_H

#include "tracer.h"
#include "bpf_helpers.h"
#include "http-types.h"

#include "http-shared-maps.h"

/* This map used for notifying userspace that a HTTP batch is ready to be consumed */
struct bpf_map_def SEC("maps/http_notifications") http_notifications = {
    .type = BPF_MAP_TYPE_PERF_EVENT_ARRAY,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u32),
    .max_entries = 0, // This will get overridden at runtime
    .pinning = 0,
    .namespace = "",
};

struct bpf_map_def SEC("maps/ssl_sock_by_ctx") ssl_sock_by_ctx = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(void *),
    .value_size = sizeof(ssl_sock_t),
    .max_entries = 1, // This will get overridden at runtime using max_tracked_connections
    .pinning = 0,
    .namespace = "",
};

struct bpf_map_def SEC("maps/ssl_read_args") ssl_read_args = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(__u64),
    .value_size = sizeof(ssl_read_args_t),
    .max_entries = 1024,
    .pinning = 0,
    .namespace = "",
};

struct bpf_map_def SEC("maps/bio_new_socket_args") bio_new_socket_args = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(__u64), // pid_tgid
    .value_size = sizeof(__u32), // socket_fd
    .max_entries = 1024,
    .pinning = 0,
    .namespace = "",
};

struct bpf_map_def SEC("maps/fd_by_ssl_bio") fd_by_ssl_bio = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(__u32),
    .value_size = sizeof(void *),
    .max_entries = 1024,
    .pinning = 0,
    .namespace = "",
};

struct bpf_map_def SEC("maps/open_at_args") open_at_args = {
    .type = BPF_MAP_TYPE_HASH,
    .key_size = sizeof(__u64), // pid_tgid
    .value_size = sizeof(lib_path_t),
    .max_entries = 1024,
    .pinning = 0,
    .namespace = "",
};

/* This map used for notifying userspace of a shared library being loaded */
struct bpf_map_def SEC("maps/shared_libraries") shared_libraries = {
    .type = BPF_MAP_TYPE_PERF_EVENT_ARRAY,
    .key_size = sizeof(__u32),
    .value_size = sizeof(__u32),
    .max_entries = 0, // This will get overridden at runtime
    .pinning = 0,
    .namespace = "",
};

#endif
