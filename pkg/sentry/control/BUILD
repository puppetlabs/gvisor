load("//tools:defs.bzl", "go_library", "go_test", "proto_library")

package(licenses = ["notice"])

proto_library(
    name = "control",
    srcs = ["control.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "control",
    srcs = [
        "control.go",
        "events.go",
        "fs.go",
        "lifecycle.go",
        "logging.go",
        "pprof.go",
        "proc.go",
        "state.go",
        "usage.go",
    ],
    visibility = [
        "//:sandbox",
    ],
    deps = [
        "//pkg/abi/linux",
        "//pkg/context",
        "//pkg/eventchannel",
        "//pkg/fd",
        "//pkg/log",
        "//pkg/sentry/fdimport",
        "//pkg/sentry/fs",
        "//pkg/sentry/fs/host",
        "//pkg/sentry/fs/user",
        "//pkg/sentry/fsimpl/host",
        "//pkg/sentry/fsmetric",
        "//pkg/sentry/kernel",
        "//pkg/sentry/kernel/auth",
        "//pkg/sentry/kernel/time",
        "//pkg/sentry/limits",
        "//pkg/sentry/state",
        "//pkg/sentry/strace",
        "//pkg/sentry/usage",
        "//pkg/sentry/vfs",
        "//pkg/sentry/watchdog",
        "//pkg/sync",
        "//pkg/tcpip/link/sniffer",
        "//pkg/urpc",
        "//pkg/usermem",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)

go_test(
    name = "control_test",
    size = "small",
    srcs = ["proc_test.go"],
    library = ":control",
    deps = [
        "//pkg/log",
        "//pkg/sentry/kernel/time",
        "//pkg/sentry/usage",
    ],
)
