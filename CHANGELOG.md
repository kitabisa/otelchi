# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.3.0](https://github.com/kitabisa/otelchi/compare/v1.2.0...v1.3.0) (2024-07-09)


### Features

* change module name ([ec40f5d](https://github.com/kitabisa/otelchi/commit/ec40f5df156d9500ad78c6a049a182a15ec3b258))

## [1.2.0](https://github.com/kitabisa/otelchi/compare/v1.1.0...v1.2.0) (2024-07-09)


### Features

* updated middleware.go ([fd00708](https://github.com/kitabisa/otelchi/commit/fd007089ca8c275828955134d30b21f4113cace0))

## [1.1.0](https://github.com/kitabisa/otelchi/compare/v1.0.0...v1.1.0) (2024-07-07)


### Features

* add filter function to skip tracing based on the request. ([a064d40](https://github.com/kitabisa/otelchi/commit/a064d4089f463d2fb89d1383190ba0249f816b2d))
* add multiple filters in WithFilter option ([3292e4d](https://github.com/kitabisa/otelchi/commit/3292e4da065d4ab9f83ea388c135c28bfc82eac9))
* add trace id to response header ([c65058a](https://github.com/kitabisa/otelchi/commit/c65058a69a0248e7ec92251c94341fb7bb2f6a6a))
* add WithPublicEndpoint & WithPublicEndpointFn option to config.go; ([64f2d2d](https://github.com/kitabisa/otelchi/commit/64f2d2da298f53e46df859183ce5a2c0c12a742e))
* added traceID on context ([59aee53](https://github.com/kitabisa/otelchi/commit/59aee539b17c9f49435ecb8f11d7c4f1fec0535b))
* adjust example for version 0.7 ([be5af93](https://github.com/kitabisa/otelchi/commit/be5af93d274c18ebb5788c161f1ec7b7bb612753))
* finish implement publicEndpointFn in middleware.go; ([b70e4bf](https://github.com/kitabisa/otelchi/commit/b70e4bf4aa67893aa664b67cf7dab4b429435222))
* remove http.target attribute based on https://github.com/open-telemetry/opentelemetry-go/blob/v1.14.0/semconv/internal/v2/http.go#L160-L165; adjust tests accordingly; ([e88a9d7](https://github.com/kitabisa/otelchi/commit/e88a9d7fffb06b71b66be83f5f97ad64ef61d527))
* renamed to trace-id ([7b7af22](https://github.com/kitabisa/otelchi/commit/7b7af2204b69f31592fbc630f350263683af993a))
* return 0.9.0 in Version(); ([cbbec48](https://github.com/kitabisa/otelchi/commit/cbbec48c5365de259e4db2dbc7e6bea1d60cdee3))
* set library version to `0.7.0`; ([10d8643](https://github.com/kitabisa/otelchi/commit/10d864338bdf1b95c0808a0ee2ad3c5ec80d8713))
* update semconv to 1.17.0 ([9172525](https://github.com/kitabisa/otelchi/commit/91725256c56af0e72fe75c3a1c4764b565c52648))
* update semconv to v1.20, update example basic to go1.19,otel1.24,semconv1.20 ([09acd4b](https://github.com/kitabisa/otelchi/commit/09acd4b56d9ca3291378242b0a7e03a56aac760d))
* upgrade to go1.21 and otel 1.26 ([fc85136](https://github.com/kitabisa/otelchi/commit/fc85136c59ed7172e1c4aab5b16e913917b7ff73))


### Bug Fixes

* broken empty routes in chi/v5.0.8 ([fc4a7b5](https://github.com/kitabisa/otelchi/commit/fc4a7b559ac0cebbe3b95730f1445eec2a8d2dcc))
* implement default trace-id in response header ([38d0a62](https://github.com/kitabisa/otelchi/commit/38d0a62003766a75920d3cd939a6372162b9d8f3))
* implement the change request ([b908ef0](https://github.com/kitabisa/otelchi/commit/b908ef01a24524c20ed1925d0c18e4300b8c4ccd))
* issue on multiservice example ([f5b4663](https://github.com/kitabisa/otelchi/commit/f5b466395c17b98a80586604a30e7027b45c32bf))
* set the default of status code to 200 in case of no op in http handler; ([cf20db3](https://github.com/kitabisa/otelchi/commit/cf20db350e1c241373737f699ef2233c317e42e7))
* update `X-Trace-ID` into `X-Trace-Id` since this is the common format (e.g X-AMZN-Trace-Id); update incorrect comment for `WithTraceIDResponseHeader` option; ([97501d1](https://github.com/kitabisa/otelchi/commit/97501d182238de47222cae0d0fc3eb27c996318c))
* use correct instrumentation version value; ([a2aab44](https://github.com/kitabisa/otelchi/commit/a2aab443cd74afdc731daa4e79a480546c8e84de))

## 1.0.0 (2024-07-07)


### Features

* add filter function to skip tracing based on the request. ([a064d40](https://github.com/kitabisa/otelchi/commit/a064d4089f463d2fb89d1383190ba0249f816b2d))
* add multiple filters in WithFilter option ([3292e4d](https://github.com/kitabisa/otelchi/commit/3292e4da065d4ab9f83ea388c135c28bfc82eac9))
* add trace id to response header ([c65058a](https://github.com/kitabisa/otelchi/commit/c65058a69a0248e7ec92251c94341fb7bb2f6a6a))
* add WithPublicEndpoint & WithPublicEndpointFn option to config.go; ([64f2d2d](https://github.com/kitabisa/otelchi/commit/64f2d2da298f53e46df859183ce5a2c0c12a742e))
* added traceID on context ([59aee53](https://github.com/kitabisa/otelchi/commit/59aee539b17c9f49435ecb8f11d7c4f1fec0535b))
* adjust example for version 0.7 ([be5af93](https://github.com/kitabisa/otelchi/commit/be5af93d274c18ebb5788c161f1ec7b7bb612753))
* finish implement publicEndpointFn in middleware.go; ([b70e4bf](https://github.com/kitabisa/otelchi/commit/b70e4bf4aa67893aa664b67cf7dab4b429435222))
* remove http.target attribute based on https://github.com/open-telemetry/opentelemetry-go/blob/v1.14.0/semconv/internal/v2/http.go#L160-L165; adjust tests accordingly; ([e88a9d7](https://github.com/kitabisa/otelchi/commit/e88a9d7fffb06b71b66be83f5f97ad64ef61d527))
* renamed to trace-id ([7b7af22](https://github.com/kitabisa/otelchi/commit/7b7af2204b69f31592fbc630f350263683af993a))
* return 0.9.0 in Version(); ([cbbec48](https://github.com/kitabisa/otelchi/commit/cbbec48c5365de259e4db2dbc7e6bea1d60cdee3))
* set library version to `0.7.0`; ([10d8643](https://github.com/kitabisa/otelchi/commit/10d864338bdf1b95c0808a0ee2ad3c5ec80d8713))
* update semconv to 1.17.0 ([9172525](https://github.com/kitabisa/otelchi/commit/91725256c56af0e72fe75c3a1c4764b565c52648))
* update semconv to v1.20, update example basic to go1.19,otel1.24,semconv1.20 ([09acd4b](https://github.com/kitabisa/otelchi/commit/09acd4b56d9ca3291378242b0a7e03a56aac760d))
* upgrade to go1.21 and otel 1.26 ([fc85136](https://github.com/kitabisa/otelchi/commit/fc85136c59ed7172e1c4aab5b16e913917b7ff73))


### Bug Fixes

* broken empty routes in chi/v5.0.8 ([fc4a7b5](https://github.com/kitabisa/otelchi/commit/fc4a7b559ac0cebbe3b95730f1445eec2a8d2dcc))
* implement default trace-id in response header ([38d0a62](https://github.com/kitabisa/otelchi/commit/38d0a62003766a75920d3cd939a6372162b9d8f3))
* implement the change request ([b908ef0](https://github.com/kitabisa/otelchi/commit/b908ef01a24524c20ed1925d0c18e4300b8c4ccd))
* issue on multiservice example ([f5b4663](https://github.com/kitabisa/otelchi/commit/f5b466395c17b98a80586604a30e7027b45c32bf))
* set the default of status code to 200 in case of no op in http handler; ([cf20db3](https://github.com/kitabisa/otelchi/commit/cf20db350e1c241373737f699ef2233c317e42e7))
* update `X-Trace-ID` into `X-Trace-Id` since this is the common format (e.g X-AMZN-Trace-Id); update incorrect comment for `WithTraceIDResponseHeader` option; ([97501d1](https://github.com/kitabisa/otelchi/commit/97501d182238de47222cae0d0fc3eb27c996318c))
* use correct instrumentation version value; ([a2aab44](https://github.com/kitabisa/otelchi/commit/a2aab443cd74afdc731daa4e79a480546c8e84de))

## [1.1.0](https://github.com/kitabisa/otelchi/compare/v1.0.0...v1.1.0) (2024-07-07)


### Features

* renamed to trace-id ([edd5dd9](https://github.com/kitabisa/otelchi/commit/edd5dd97f9db6840811618810e7ded3973be23ed))

## [Unreleased]

## [0.9.0] - 2024-07-06

### Changed

- `WithFilter` option now support multiple filter functions, just like in [otelmux](https://github.com/open-telemetry/opentelemetry-go-contrib/blob/v1.24.0/instrumentation/github.com/gorilla/mux/otelmux/config.go#L106-L110). ([#47])
- Upgrade `go.opentelemetry.io/otel`, `go.opentelemetry.io/otel/sdk`, & `go.opentelemetry.io/otel/trace` to `v1.28.0`. ([#49])
- Upgrade `github.com/go-chi/chi/v5` to `v5.1.0`. ([#49])
- Set the go versions for testing in both `Makefile` & `compatibility-test.yml` to `1.21` & `1.22`. ([#49])

### Removed

- Drop support for Go `<1.21`. ([#49])

## [0.8.0] - 2024-04-29

### ⚠️ Notice ⚠️

This release is the last to support Go `1.19`. The next release will require at least Go `1.21`.

### Added

- Add `WithPublicEndpoint` & `WithPublicEndpointFn` options. ([#43])

### Changed

- Upgrade to `v1.24.0` of `go.opentelemetry.io/otel`. ([#41])
- Upgrade to `v1.20.0` of `go.opentelemetry.io/otel/semconv`. ([#41])
- Adjust Go version for both `examples/basic` & `examples/multi-services` to `1.19` & `go.opentelemetry.io/otel` to `v1.24.0`. ([#41])
- Update otelhttp version to `0.49.0` since it is the version that uses otel `1.24.0` internally, check [here](https://github.com/open-telemetry/opentelemetry-go-contrib/blob/v1.24.0/instrumentation/net/http/otelhttp/go.mod#L8) for details. ([#42])
- Set the go versions in compatibility-test.yml to 1.19, 1.20, & 1.21. ([#42])
- Set the sampling strategy to always sample in test cases to avoid random error. ([#42])
- Use `otlptrace` exporter instead of `jaeger` exporter in `examples/multi-services`. ([#42])

### Removed

- Remove the deprecated `jaeger` exporter from `examples/multi-services` & use `otlptrace` exporter instead. ([#42])
- Drop support for Go `<1.19`. ([#41])

## [0.7.0] - 2024-04-22

### ⚠️ Notice ⚠️

This release is the last to support Go `1.18`. The next release will require at least Go `1.19`.

### Changed

- Upgrade to `v1.14.0` of `go.opentelemetry.io/otel`. ([#38])
- Upgrade to `v1.17.0` of `go.opentelemetry.io/otel/semconv`. ([#38])
- Adjust Go version for both `examples/basic` & `examples/multi-services` to `1.18` & `go.opentelemetry.io/otel` to `v1.14.0`. ([#38])
- Change `http.server_name` attributes to `net.host.name`, this is because semconv is removing this attribute for http. ([#38])

### Removed

- Remove `http.target` attribute on implementation & tests based on [this comment](https://github.com/open-telemetry/opentelemetry-go/blob/v1.17.0/semconv/internal/v2/http.go#L160-L165). ([#39])
- Drop support for Go `<1.18`. ([#38])

## [0.6.0] - 2024-04-02

### ⚠️ Notice ⚠️

This release is the last to support Go `1.15`. The next release will require at least Go `1.18`.

### Added

- Add `WithTraceIDResponseHeader` option to enable adding trace id into response header. ([#36])
- Add multiple go versions test scripts for local and CI pipeline. ([#29])
- Add compatibility testing for `ubuntu`, `macos` and `windows`. ([#32])
- Add repo essentials docs. ([#33])

### Changed

- Upgrade to `v5.0.12` of `go-chi/chi`. ([#29])
- Upgrade to `v1.10.0` of `go.opentelemetry.io/otel`. ([#29])
- Upgrade to `v1.12.0` of `go.opentelemetry.io/otel/semconv`. ([#29])
- Set the required go version for both `examples/basic` & `examples/multi-services` to `1.15`, `go-chi/chi` to `v5.0.12`, & `go.opentelemetry.io/otel` to `v1.10.0` ([#35])

## [0.5.2] - 2024-03-25

### Fixed

- Fix empty status code. ([#30])

### Changed

- Return `http.StatusOK` (200) as a default `http.status_code` span attribute. ([#30])

## [0.5.1] - 2023-02-18

### Fixed

- Fix broken empty routes. ([#18])

### Changed

- Upgrade to `v5.0.8` of `go-chi/chi`.

## [0.5.0] - 2022-10-02

### Added

- Add multi services example. ([#9])
- Add `WithFilter()` option to ignore tracing in certain endpoints. ([#11])

## [0.4.0] - 2022-02-22

### Added

- Add Option `WithRequestMethodInSpanName()` to handle vendor that do not include HTTP request method as mentioned in [#6]. ([#7])
- Refine description for `WithChiRoutes()` option to announce it is possible to override the span name in underlying handler with this option.

### Changed

## [0.3.0] - 2022-01-18

### Fixed

- Fix both `docker-compose.yml` & `Dockerfile` in the example. ([#5])

### Added

- Add `WithChiRoutes()` option to make the middleware able to determine full route pattern on span creation. ([#5])
- Set all known span attributes on span creation rather than set them after request is being executed. ([#5])

## [0.2.1] - 2022-01-08

### Added

- Add build example to CI pipeline. ([#2])

### Changed

- Use `ctx.RoutePattern()` to get span name, this is to strip out noisy wildcard pattern. ([#1])

## [0.2.0] - 2021-10-18

### Added

- Set service name on tracer provider from code example.

### Changed

- Update dependencies in go.mod
- Upgrade to `v1.0.1` of `go.opentelemetry.io/otel`.
- Upgrade to `v5.0.4` of `go-chi/chi`.
- Update latest test to use `otelmux` format.

### Removed

- Remove `HTTPResponseContentLengthKey`
- Remove `HTTPTargetKey`, since automatically set in `HTTPServerAttributesFromHTTPRequest`

## [0.1.0] - 2021-08-11

This is the first release of otelchi.
It contains instrumentation for trace and depends on:

- otel => `v1.0.0-RC2`
- go-chi/chi => `v5.0.3`

### Added

- Instrumentation for trace.
- CI files.
- Example code for a basic usage.
- Apache-2.0 license.

[#49]: https://github.com/riandyrn/otelchi/pull/49
[#47]: https://github.com/riandyrn/otelchi/pull/47
[#43]: https://github.com/riandyrn/otelchi/pull/43
[#42]: https://github.com/riandyrn/otelchi/pull/42
[#41]: https://github.com/riandyrn/otelchi/pull/41
[#39]: https://github.com/riandyrn/otelchi/pull/39
[#38]: https://github.com/riandyrn/otelchi/pull/38
[#36]: https://github.com/riandyrn/otelchi/pull/36
[#35]: https://github.com/riandyrn/otelchi/pull/35
[#33]: https://github.com/riandyrn/otelchi/pull/33
[#32]: https://github.com/riandyrn/otelchi/pull/32
[#30]: https://github.com/riandyrn/otelchi/pull/30
[#29]: https://github.com/riandyrn/otelchi/pull/29
[#18]: https://github.com/riandyrn/otelchi/pull/18
[#11]: https://github.com/riandyrn/otelchi/pull/11
[#9]: https://github.com/riandyrn/otelchi/pull/9
[#7]: https://github.com/riandyrn/otelchi/pull/7
[#6]: https://github.com/riandyrn/otelchi/pull/6
[#5]: https://github.com/riandyrn/otelchi/pull/5
[#2]: https://github.com/riandyrn/otelchi/pull/2
[#1]: https://github.com/riandyrn/otelchi/pull/1

[Unreleased]: https://github.com/riandyrn/otelchi/compare/v0.9.0...HEAD
[0.9.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.9.0
[0.8.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.8.0
[0.7.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.7.0
[0.6.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.6.0
[0.5.2]: https://github.com/riandyrn/otelchi/releases/tag/v0.5.2
[0.5.1]: https://github.com/riandyrn/otelchi/releases/tag/v0.5.1
[0.5.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.5.0
[0.4.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.4.0
[0.3.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.3.0
[0.2.1]: https://github.com/riandyrn/otelchi/releases/tag/v0.2.1
[0.2.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.2.0
[0.1.0]: https://github.com/riandyrn/otelchi/releases/tag/v0.1.0
