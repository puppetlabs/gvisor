load("//tools:defs.bzl", "go_library", "go_test")

package(licenses = ["notice"])

go_library(
    name = "transport",
    srcs = [
        "datagram.go",
        "transport.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//pkg/tcpip"],
)

go_test(
    name = "datagram_test",
    size = "small",
    srcs = ["datagram_test.go"],
    deps = [
        ":transport",
        "//pkg/tcpip",
        "//pkg/tcpip/header",
        "//pkg/tcpip/link/loopback",
        "//pkg/tcpip/network/ipv4",
        "//pkg/tcpip/stack",
        "//pkg/tcpip/testutil",
        "//pkg/tcpip/transport/icmp",
        "//pkg/tcpip/transport/raw",
        "//pkg/tcpip/transport/udp",
        "//pkg/waiter",
    ],
)
