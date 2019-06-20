load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.18.6/rules_go-0.18.6.tar.gz",
    ],
    sha256 = "f04d2373bcaf8aa09bccb08a98a57e721306c8f6043a2a0ee610fd6853dcde3d",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

http_archive(
    name="org_golang_x_tools",
    urls=["https://github.com/golang/tools/archive/master.zip"],
    sha256="258ad1e138ae70a990f5612bc04662856721e347d68b84a1e6d370265aa0d7e8",

)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

#go_repository(
#    name = "org_golang_x_tools",
#    remote = "https://github.com/golang/tools"
#)

#go_repository(
#    name = "com_google_protobuf",
#    build_file_proto_mode = "disable_global",
#    importpath = "github.com/google/protobuf",
#    urls = ["http://bazel-cabin.bilibili.co/google/protobuf/48cb18e5c419ddd23d9badcfe4e9df7bde1979b2"],
#    strip_prefix = "protobuf-48cb18e5c419ddd23d9badcfe4e9df7bde1979b2",
#    goype = "zip",
#)
#
#go_repository(
#    name = "org_golang_x_tools ",
#    build_file_proto_mode = "disable_global",
#    importpath = "github.com/golang/tools",
#    urls = ["https://github.com/golang/tools/archive/master.zip"],
#    type = "zip",
#)


go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()