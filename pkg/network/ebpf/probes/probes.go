// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux_bpf
// +build linux_bpf

package probes

import "fmt"

// ProbeName stores the name of the kernel probes setup for tracing
type ProbeName string

const (
	// InetCskListenStop traces the inet_csk_listen_stop system call (called for both ipv4 and ipv6)
	InetCskListenStop ProbeName = "kprobe/inet_csk_listen_stop"

	// TCPv6Connect traces the v6 connect() system call
	TCPv6Connect ProbeName = "kprobe/tcp_v6_connect"
	// TCPv6ConnectReturn traces the return value for the v6 connect() system call
	TCPv6ConnectReturn ProbeName = "kretprobe/tcp_v6_connect"

	// TCPSendMsg traces the tcp_sendmsg() system call
	TCPSendMsg ProbeName = "kprobe/tcp_sendmsg"

	// TCPSendMsgPre410 traces the tcp_sendmsg() system call on kernels prior to 4.1.0. This is created because
	// we need to load a different kprobe implementation
	TCPSendMsgPre410 ProbeName = "kprobe/tcp_sendmsg/pre_4_1_0"

	// TCPSendMsgReturn traces the return value for the tcp_sendmsg() system call
	// XXX: This is only used for telemetry for now to count the number of errors returned
	// by the tcp_sendmsg func (so we can have a # of tcp sent bytes we miscounted)
	TCPSendMsgReturn ProbeName = "kretprobe/tcp_sendmsg"

	// TCPGetSockOpt traces the tcp_getsockopt() kernel function
	// This probe is used for offset guessing only
	TCPGetSockOpt ProbeName = "kprobe/tcp_getsockopt"

	// SockGetSockOpt traces the sock_common_getsockopt() kernel function
	// This probe is used for offset guessing only
	SockGetSockOpt ProbeName = "kprobe/sock_common_getsockopt"

	// TCPSetState traces the tcp_set_state() kernel function
	TCPSetState ProbeName = "kprobe/tcp_set_state"

	// TCPCleanupRBuf traces the tcp_cleanup_rbuf() system call
	TCPCleanupRBuf ProbeName = "kprobe/tcp_cleanup_rbuf"
	// TCPClose traces the tcp_close() system call
	TCPClose ProbeName = "kprobe/tcp_close"
	// TCPCloseReturn traces the return of tcp_close() system call
	TCPCloseReturn ProbeName = "kretprobe/tcp_close"

	// We use the following two probes for UDP sends
	IPMakeSkb        ProbeName = "kprobe/ip_make_skb"
	IP6MakeSkb       ProbeName = "kprobe/ip6_make_skb"
	IP6MakeSkbPre470 ProbeName = "kprobe/ip6_make_skb/pre_4_7_0"

	// UDPRecvMsg traces the udp_recvmsg() system call
	UDPRecvMsg ProbeName = "kprobe/udp_recvmsg"
	// UDPRecvMsgPre410 traces the udp_recvmsg() system call on kernels prior to 4.1.0
	UDPRecvMsgPre410 ProbeName = "kprobe/udp_recvmsg/pre_4_1_0"
	// UDPRecvMsgReturn traces the return value for the udp_recvmsg() system call
	UDPRecvMsgReturn ProbeName = "kretprobe/udp_recvmsg"

	// UDPDestroySock traces the udp_destroy_sock() function
	UDPDestroySock ProbeName = "kprobe/udp_destroy_sock"
	// UDPDestroySockrReturn traces the return of the udp_destroy_sock() system call
	UDPDestroySockReturn ProbeName = "kretprobe/udp_destroy_sock"

	// TCPRetransmit traces the return value for the tcp_retransmit_skb() system call
	TCPRetransmit       ProbeName = "kprobe/tcp_retransmit_skb"
	TCPRetransmitPre470 ProbeName = "kprobe/tcp_retransmit_skb/pre_4_7_0"

	// InetCskAcceptReturn traces the return value for the inet_csk_accept syscall
	InetCskAcceptReturn ProbeName = "kretprobe/inet_csk_accept"

	// InetBind is the kprobe of the bind() syscall for IPv4
	InetBind ProbeName = "kprobe/inet_bind"
	// Inet6Bind is the kprobe of the bind() syscall for IPv6
	Inet6Bind ProbeName = "kprobe/inet6_bind"

	// InetBind is the kretprobe of the bind() syscall for IPv4
	InetBindRet ProbeName = "kretprobe/inet_bind"
	// Inet6Bind is the kretprobe of the bind() syscall for IPv6
	Inet6BindRet ProbeName = "kretprobe/inet6_bind"

	// SocketDnsFilter is the socket probe for dns
	SocketDnsFilter ProbeName = "socket/dns_filter"

	// SockMapFdReturn maps a file descriptor to a kernel sock
	SockMapFdReturn ProbeName = "kretprobe/sockfd_lookup_light"

	// ConntrackHashInsert is the probe for new conntrack entries
	ConntrackHashInsert ProbeName = "kprobe/__nf_conntrack_hash_insert"

	// SockFDInstall is the kprobe used for mapping socket FDs to kernel sock structs (kernel >= 5.5) for sys_connect() sys_accept()
	SockFDInstall ProbeName = "kprobe/fd_install"

	// SockFDLookup is the kprobe used for mapping socket FDs to kernel sock structs
	SockFDLookup ProbeName = "kprobe/sockfd_lookup_light"

	// SockFDLookupRet is the kretprobe used for mapping socket FDs to kernel sock structs
	SockFDLookupRet ProbeName = "kretprobe/sockfd_lookup_light"

	// DoSendfile is the kprobe used to trace traffic via SENDFILE(2) syscall
	DoSendfile ProbeName = "kprobe/do_sendfile"

	// DoSendfileRet is the kretprobe used to trace traffic via SENDFILE(2) syscall
	DoSendfileRet ProbeName = "kretprobe/do_sendfile"
)

// BPFMapName stores the name of the BPF maps storing statistics and other info
type BPFMapName string

const (
	ConnMap               BPFMapName = "conn_stats"
	TcpStatsMap           BPFMapName = "tcp_stats"
	ConnCloseEventMap     BPFMapName = "conn_close_event"
	TracerStatusMap       BPFMapName = "tracer_status"
	PortBindingsMap       BPFMapName = "port_bindings"
	UdpPortBindingsMap    BPFMapName = "udp_port_bindings"
	TelemetryMap          BPFMapName = "telemetry"
	ConnCloseBatchMap     BPFMapName = "conn_close_batch"
	ConntrackMap          BPFMapName = "conntrack"
	ConntrackTelemetryMap BPFMapName = "conntrack_telemetry"
	SockFDLookupArgsMap   BPFMapName = "sockfd_lookup_args"
	DoSendfileArgsMap     BPFMapName = "do_sendfile_args"
	SockByPidFDMap        BPFMapName = "sock_by_pid_fd"
	PidFDBySockMap        BPFMapName = "pid_fd_by_sock"
	TagsMap               BPFMapName = "conn_tags"
)

// SectionName returns the SectionName for the given BPF map
func (b BPFMapName) SectionName() string {
	return fmt.Sprintf("maps/%s", b)
}
