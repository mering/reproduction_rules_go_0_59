load("@rules_cc//cc:cc_library.bzl", "cc_library")
load("@rules_go//go:def.bzl", "go_library", "go_test")

cc_library(
    name = "world",
    deps = [
        "@abseil-cpp//absl/log:flags",
    ],
)

go_library(
    name = "world_go",
    srcs = ["world.go"],
    cdeps = [":world"],
    cgo = 1,
)

go_test(
    name = "world_go_test",
    srcs = ["world_go_test.go"],
    cdeps = [":world"],
    cgo = 1,
    deps = [
        ":world_go",
    ],
)
