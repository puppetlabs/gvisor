load("//tools:defs.bzl", "go_library")
load("//tools/go_generics:defs.bzl", "go_template_instance")

package(licenses = ["notice"])

go_template_instance(
    name = "atomicptr_credentials",
    out = "atomicptr_credentials_unsafe.go",
    package = "auth",
    suffix = "Credentials",
    template = "//pkg/sync/atomicptr:generic_atomicptr",
    types = {
        "Value": "Credentials",
    },
)

go_template_instance(
    name = "id_map_range",
    out = "id_map_range.go",
    package = "auth",
    prefix = "idMap",
    template = "//pkg/segment:generic_range",
    types = {
        "T": "uint32",
    },
)

go_template_instance(
    name = "id_map_set",
    out = "id_map_set.go",
    consts = {
        "minDegree": "3",
    },
    package = "auth",
    prefix = "idMap",
    template = "//pkg/segment:generic_set",
    types = {
        "Key": "uint32",
        "Range": "idMapRange",
        "Value": "uint32",
        "Functions": "idMapFunctions",
    },
)

go_library(
    name = "auth",
    srcs = [
        "atomicptr_credentials_unsafe.go",
        "auth.go",
        "capability_set.go",
        "context.go",
        "credentials.go",
        "id.go",
        "id_map.go",
        "id_map_functions.go",
        "id_map_range.go",
        "id_map_set.go",
        "user_namespace.go",
    ],
    marshal = True,
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/bits",
        "//pkg/context",
        "//pkg/errors/linuxerr",
        "//pkg/log",
        "//pkg/sync",
    ],
)
