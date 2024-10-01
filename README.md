![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/hurrah/coverage.svg)
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![MultiPlatformUnitTest](https://github.com/nao1215/hurrah/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/hurrah/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/hurrah/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/hurrah/actions/workflows/reviewdog.yml)

## hurrah - API Gateway 

> [!WARNING]
> This project is under development.

The hurrah command provides API Gateway and Proxy functionality. There are several motivations for using a custom API Gateway instead of a managed service like Amazon API Gateway. For example,

- To avoid cost increases due to pay-as-you-go pricing.
- Flexible customization such as setting the number of retries and timeouts.

The hurrah project greatly reflects my personal goals, which are as follows:

- Getting over 1k stars
- Selling paid features via plugins (e.g., selling an OIDC plugin)
- Acquiring knowledge related to API Gateways

However, I love the culture of Open Source, so I want to keep as much of the code free as possible. If the features mature and I receive support from everyone, such as through GitHub Sponsors, I may develop everything as free features.

## How to install

### Use "go install"

```shell
go install github.com/nao1215/hurrah/cmd/hurrah@latest
```

### Use homebrew

```shell
brew install nao1215/tap/hurrah
```

## How to use
### flag
```shell
 ./hurrah -h
Usage:
  -config string
        a path to the configuration file (default "config.toml")
  -debug
        whether to run in debug mode. By default, only output warning/error logs
  -port string
        a port number to listen on (default "8080")
```

### config.toml
The hurrah command reads the configuration from the `config.toml` file. You can specify the configuration file with the `-config` flag.

```toml
[[routes]]
path = "/service1"
backend = "http://localhost:8081"
timeout = 10

[[routes]]
path = "/service2"
backend = "http://localhost:8082"
```

| Key | Description |
| --- | ----------- |
| path | The path to match the incoming request. |
| backend | The URL to forward the request to. |
| timeout | The timeout for the request. By default, it is 30 seconds. |


## Roadmap

- [ ] **Routing**
  - [ ] Path-based routing
    - [ ] Define URL paths and corresponding backend services
    - [ ] Implement dynamic routing logic (e.g., based on `/service1` → backend A, `/service2` → backend B)
    - [ ] Create a TOML file for configuration and support hot-reloading
    - [ ] Ensure proper request forwarding to target backend services
    - [ ] Add tests for various routing paths

  - [ ] **Host-based routing**
    - [ ] Implement routing based on the host header (e.g., `service1.example.com` routes to backend A, `service2.example.com` routes to backend B)
    - [ ] Support wildcard and subdomain matching (e.g., `*.example.com`)
    - [ ] Add validation to ensure hostnames are correctly mapped
  
  - [ ] **Method-based routing**
    - [ ] Route requests based on HTTP methods (GET, POST, etc.)
    - [ ] Create configurations allowing specific backends for different methods on the same path
    - [ ] Add rate limiting or method-specific processing (e.g., CORS handling for OPTIONS requests)

  - [ ] **Header-based routing**
    - [ ] Add routing logic that handles requests based on specific HTTP headers
    - [ ] Implement rules to route based on content-type, custom headers, or cookies
    - [ ] Configure load balancing based on header values (e.g., `User-Agent` routing)

  - [ ] **Query Parameter-based routing**
    - [ ] Route traffic based on query parameters (e.g., `/search?service=backendA` routes to backend A)
    - [ ] Add validation and security checks on query parameters
    - [ ] Create fallback behavior for missing or invalid parameters

- [ ] **Security**
  - [ ] **Authentication and Authorization**
    - [ ] Integrate OAuth2 or OpenID Connect (OIDC) for authentication
    - [ ] Implement JWT validation for API requests
    - [ ] Support API keys and HMAC signing
    - [ ] Add role-based access control (RBAC) for different routes
  
  - [ ] **Rate Limiting**
    - [ ] Implement rate limiting per route, per user, or globally
    - [ ] Add the ability to configure rate limits dynamically
    - [ ] Integrate IP-based blocking for misuse

  - [ ] **Traffic Encryption**
    - [ ] Enable SSL/TLS for secure communication between clients and the gateway
    - [ ] Add support for mutual TLS (mTLS) between the gateway and backend services
    - [ ] Ensure automatic certificate renewal (e.g., with Let's Encrypt)

- [ ] **Load Balancing**
  - [ ] Implement round-robin or least connection-based load balancing
  - [ ] Add health checks for backend services to ensure uptime
  - [ ] Configure failover behavior for backend service downtime
  - [ ] Support sticky sessions or affinity

- [ ] **Monitoring and Logging**
  - [ ] Integrate request and error logging for all routed requests
  - [ ] Set up access logs with details such as client IP, user agent, response status, etc.
  - [ ] Implement request tracing (e.g., via OpenTelemetry or Jaeger)
  - [ ] Add support for monitoring metrics like response times, error rates, and request counts

- [ ] **Caching**
  - [ ] Implement caching for static resources
  - [ ] Add support for request-level caching based on cache-control headers
  - [ ] Create invalidation rules for cache expiry

- [ ] **Middleware**
  - [ ] Implement middleware for common tasks (e.g., request validation, authentication, logging)
  - [ ] Add support for custom middleware (user-defined logic)

- [ ] **Plugins**
  - [ ] Create a plugin architecture for extending functionality (e.g., OIDC plugin)
  - [ ] Define a plugin interface and lifecycle management (loading, unloading, configuration)
  - [ ] Allow external developers to build and share custom plugins



## Contributing
First off, thanks for taking the time to contribute! ❤️  See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.
Contributions are not only related to development. For example, GitHub Star motivates me to develop!

### Star History
[![Star History Chart](https://api.star-history.com/svg?repos=nao1215/hurrah&type=Date)](https://star-history.com/#nao1215/hurrah&Date)


## Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/hurrah/issues)


## LICENSE
[Apache License Version 2.0](./LICENSE).

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://debimate.jp/"><img src="https://avatars.githubusercontent.com/u/22737008?v=4?s=70" width="70px;" alt="CHIKAMATSU Naohiro"/><br /><sub><b>CHIKAMATSU Naohiro</b></sub></a><br /><a href="https://github.com/nao1215/hurrah/commits?author=nao1215" title="Code">💻</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!