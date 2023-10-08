load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

# rules_proto defines abstract rules for building Protocol Buffers.
http_archive(
    name = "rules_proto",
    sha256 = "2490dca4f249b8a9a3ab07bd1ba6eca085aaf8e45a734af92aad0c42d9dc7aaf",
    strip_prefix = "rules_proto-218ffa7dfa5408492dc86c01ee637614f8695c45",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/218ffa7dfa5408492dc86c01ee637614f8695c45.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/218ffa7dfa5408492dc86c01ee637614f8695c45.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)



load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(version = "1.20.7")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

load("@bazel_gazelle//:deps.bzl", "go_repository")

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

go_repository(
    name = "com_github_jackc_pgpassfile",
    importpath = "github.com/jackc/pgpassfile",
    sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jackc_pgservicefile",
    importpath = "github.com/jackc/pgservicefile",
    sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
    version = "v0.0.0-20221227161230-091c0ba34f0a",
)

go_repository(
    name = "com_github_jackc_pgx_v5",
    importpath = "github.com/jackc/pgx/v5",
    sum = "h1:cxFyXhxlvAifxnkKKdlxv8XqUf59tDlYjnV5YYfsJJY=",
    version = "v5.4.3",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//container:pull.bzl", "container_pull")

container_pull(
    name = "go_base",
    registry = "index.docker.io",
    repository = "golang",
    tag = "1.20",
)





#
#go_repository(
#    name = "io_temporal_go_api",
#    importpath = "go.temporal.io/api",
#    sum = "h1:l2HrMI/gE5JwFu9wgmZdofBIQ5MzziOEBs8mnbJUcJs=",
#    version = "v1.21.0",
#)
#go_repository(
#    name = "io_temporal_go_sdk",
#    importpath = "go.temporal.io/sdk",
#    sum = "h1:mAk5VFR+z4s8QVzRx3iIpRnHcEO3m10CYNjnRXrhVq4=",
#    version = "v1.24.0",
#)
#
#go_repository(
#    name = "com_github_pborman_uuid",
#    importpath = "github.com/pborman/uuid",
#    sum = "h1:+ZZIw58t/ozdjRaXh/3awHfmWRbzYxJoAdNJxe/3pvw=",
#    version = "v1.2.1",
#)
#
#go_repository(
#    name = "com_github_gogo_googleapis",
#    importpath = "github.com/gogo/googleapis",
#    sum = "h1:1Yx4Myt7BxzvUr5ldGSbwYiZG6t9wGBZ+8/fX3Wvtq0=",
#    version = "v1.4.1",
#)
#
#go_repository(
#    name = "com_github_gogo_protobuf",
#    importpath = "github.com/gogo/protobuf",
#    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
#    version = "v1.3.2",
#)
#
#go_repository(
#    name = "com_github_gogo_status",
#    importpath = "github.com/gogo/status",
#    sum = "h1:DuHXlSFHNKqTQ+/ACf5Vs6r4X/dH2EgIzR9Vr+H65kg=",
#    version = "v1.1.1",
#)
#go_repository(
#    name = "com_github_golang_freetype",
#    importpath = "github.com/golang/freetype",
#    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
#    version = "v0.0.0-20170609003504-e2365dfdc4a0",
#)
#
#go_repository(
#    name = "com_github_robfig_cron",
#    importpath = "github.com/robfig/cron",
#    sum = "h1:ZjScXvvxeQ63Dbyxy76Fj3AT3Ut0aKsyd2/tl3DTMuQ=",
#    version = "v1.2.0",
#)
#go_repository(
#    name = "com_github_facebookgo_clock",
#    importpath = "github.com/facebookgo/clock",
#    sum = "h1:yDWHCSQ40h88yih2JAcL6Ls/kVkSE8GFACTGVnMPruw=",
#    version = "v0.0.0-20150410010913-600d898af40a",
#)
#go_repository(
#    name = "org_golang_x_time",
#    importpath = "golang.org/x/time",
#    sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
#    version = "v0.3.0",
#)
#go_repository(
#    name = "com_github_stretchr_testify",
#    importpath = "github.com/stretchr/testify",
#    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
#    version = "v1.8.4",
#)
#go_repository(
#    name = "com_github_grpc_ecosystem_go_grpc_middleware",
#    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
#    sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
#    version = "v1.3.0",
#)
#go_repository(
#    name = "org_uber_go_atomic",
#    importpath = "go.uber.org/atomic",
#    sum = "h1:ECmE8Bn/WFTYwEW/bpKD3M8VtR/zQVbavAoalC1PYyE=",
#    version = "v1.9.0",
#)
#go_repository(
#    name = "com_github_davecgh_go_spew",
#    importpath = "github.com/davecgh/go-spew",
#    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
#    version = "v1.1.1",
#)
#go_repository(
#    name = "com_github_stretchr_objx",
#    importpath = "github.com/stretchr/objx",
#    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
#    version = "v0.5.0",
#)
#go_repository(
#    name = "com_github_ghodss_yaml",
#    importpath = "github.com/ghodss/yaml",
#    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
#    version = "v1.0.0",
#)
#go_repository(
#    name = "in_gopkg_yaml_v2",
#    importpath = "gopkg.in/yaml.v2",
#    sum = "h1:fvjTMHxHEw/mxHbtzPi3JCcKXQRAnQTBRo6YCJSVHKI=",
#    version = "v2.2.3",
#)
#go_repository(
#    name = "in_gopkg_yaml_v3",
#    importpath = "gopkg.in/yaml.v3",
#    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
#    version = "v3.0.1",
#)