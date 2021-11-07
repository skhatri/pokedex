workspace(name = "pokedex-clients")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("//:deps.bzl", "go_custom_dependencies", "go_dependencies")

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_custom_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "1f4e59843b61981a96835dc4ac377ad4da9f8c334ebe5e0bb3f58f80c09735f4",
    strip_prefix = "rules_docker-0.19.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.19.0/rules_docker-v0.19.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()
load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()




load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

RULES_JVM_EXTERNAL_TAG = "4.1"
RULES_JVM_EXTERNAL_SHA = "f36441aa876c4f6427bfb2d1f2d723b48e9d930b62662bf723ddfb8fc80f0140"
http_archive(
    name = "rules_jvm_external",
    strip_prefix = "rules_jvm_external-%s" % RULES_JVM_EXTERNAL_TAG,
    sha256 = RULES_JVM_EXTERNAL_SHA,
    url = "https://github.com/bazelbuild/rules_jvm_external/archive/%s.zip" % RULES_JVM_EXTERNAL_TAG,
)


load("@rules_jvm_external//:defs.bzl", "maven_install")

maven_install(
    name = "java_client",
    artifacts = [
        "org.junit.jupiter:junit-jupiter-api:5.8.1",
        "org.junit.jupiter:junit-jupiter-engine:5.8.1",
        "org.junit.jupiter:junit-jupiter-params:5.8.1",
        "org.junit.vintage:junit-vintage-engine:5.8.1",
        "org.junit.platform:junit-platform-commons:1.8.1",
        "org.junit.platform:junit-platform-console:1.8.1",
        "org.junit.platform:junit-platform-runner:1.8.1",
        "org.junit.platform:junit-platform-launcher:1.8.1",
        "org.junit.platform:junit-platform-engine:1.8.1",

        "com.fasterxml.jackson.core:jackson-core:2.12.5",
        "com.fasterxml.jackson.core:jackson-databind:2.12.5",

        "com.squareup.okhttp3:okhttp:4.9.2",
    ],
    repositories = [
        "https://repo1.maven.org/maven2",
    ],
)


http_archive(
    name = "rules_rust",
    sha256 = "531bdd470728b61ce41cf7604dc4f9a115983e455d46ac1d0c1632f613ab9fc3",
    strip_prefix = "rules_rust-d8238877c0e552639d3e057aadd6bfcf37592408",
    urls = [
        "https://github.com/bazelbuild/rules_rust/archive/d8238877c0e552639d3e057aadd6bfcf37592408.tar.gz",
    ],
)

load("@rules_rust//rust:repositories.bzl", "rust_repositories")

rust_repositories(version = "1.53.0", edition="2018")

http_archive(
    name = "cargo_raze",
    sha256 = "c664e258ea79e7e4ec2f2b57bca8b1c37f11c8d5748e02b8224810da969eb681",
    strip_prefix = "cargo-raze-0.11.0",
    url = "https://github.com/google/cargo-raze/archive/v0.11.0.tar.gz",
)

load("@cargo_raze//:repositories.bzl", "cargo_raze_repositories")
cargo_raze_repositories()

load("@cargo_raze//:transitive_deps.bzl", "cargo_raze_transitive_deps")
cargo_raze_transitive_deps()

#cargo install cargo-raze
#cargo raze
load("//cargo:crates.bzl", "raze_fetch_remote_crates")
raze_fetch_remote_crates()
