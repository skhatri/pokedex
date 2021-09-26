load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_skhatri_api_router_go",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/skhatri/api-router-go",
        sum = "h1:9+JPcjL9l9Thnfs9Uw7FQ+u9V2Mt2o5L3rBXcLaBLQo=",
        version = "v0.9.5",
    )

def go_custom_dependencies():
    go_repository(
        name = "com_google_uuid",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/google/uuid",
        sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
        version = "v1.3.0",
    )
