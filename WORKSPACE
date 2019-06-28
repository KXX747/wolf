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

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies","go_repository")

go_repository(
    name = "com_github_KXX747_wolf",
    commit = "39ee1c966ad66d75df32c248bb5513d649113528",
    importpath = "github.com/KXX747/wolf",
    # urls = ["https://github.com/cespare/xxhash/archive/master.zip"],
     #type = "zip",
)

go_repository(
    name = "org_golang_x_tools",
    #commit = "4adf7a708c2de4c9ea24a1f351c2e1c9b82fbde8",
    importpath = "golang.org/x/tools",
    urls = ["https://github.com/golang/tools/archive/master.zip"],
    type = "zip",
)

go_repository(
    name = "org_golang_x_sys",
    #commit = "4adf7a708c2de4c9ea24a1f351c2e1c9b82fbde8",
    importpath = "golang.org/x/sys",
    urls = ["https://github.com/golang/sys/archive/master.zip"],
    type = "zip",
)

go_repository(
    name = "org_golang_x_text",
    #commit = "4adf7a708c2de4c9ea24a1f351c2e1c9b82fbde8",
    importpath = "golang.org/x/text",
    urls = ["https://github.com/golang/text/archive/master.zip"],
    type = "zip",
)


go_repository(
    name = "org_golang_x_net",
    #commit = "4adf7a708c2de4c9ea24a1f351c2e1c9b82fbde8",
    importpath = "golang.org/x/net",
    urls = ["https://github.com/golang/net/archive/master.zip"],
    type = "zip",
)




gazelle_dependencies()



load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

#load("//build:workspace.bzl", "bili_workspace")
#bili_workspace()

