# galf

**Lite and easy deploy-able feature-flag service**

[![Go Version][go-badge]][go] [![License: MIT][license-badge]][license]

---

## 🚀 What is galf?

`galf` is a minimal, self-hostable feature-flag service written in Go.  
It provides a simple API for enabling/disabling flags across your applications without requiring heavy infrastructure.

Use it when you want:

- A lightweight feature-flag server you can deploy anywhere  
- A simple and predictable API  
- No external dependencies  
- Fast startup and low resource usage  
- Clean Go code that is easy to extend  

---

## 📁 Project Structure

```

galf/
├── cmd/            # Binaries / CLI entrypoints (coming soon)
├── api/            # HTTP handlers for feature-flag endpoints
├── internal/       # Core business logic (flag storage, evaluation, utils)
├── testutil/       # Shared test helpers
├── go.mod          # Go module definition
└── README.md       # Project documentation

```

As the project grows, expect additional folders:

```

docs/       # API docs, architecture notes
examples/   # Sample clients, example configs
deploy/     # Dockerfiles, Kubernetes manifests, Helm charts

````

---

## 🧪 Getting Started (Development)

### Requirements

- Go 1.20+  
- Git  

---

### Clone the project

```bash
git clone https://github.com/brettearle/galf.git
cd galf
````

---

### Build

```bash
go build ./cmd/galf
```

---

### Run locally

```bash
./galf --config path/to/config.yml
```

Example feature flag API call:

```http
GET /flags/my-feature
```

Example JSON response:

```json
{
  "name": "my-feature",
  "state": true
}
```

*(Exact API may change as the project evolves.)*

---

## 🔧 Contributing

Contributions are welcome — this project is in an early stage, and architectural decisions are still settling.

### Basic workflow

1. Fork the repo
2. Create a branch: `git checkout -b feature/your-change`
3. Add your code + tests
4. Run `go test ./...`
5. Commit: `git commit -m "Add your feature"`
6. Push and open a Pull Request

Please keep PRs focused and well-documented.

---

## 🛣️ Roadmap (Planned Features)

* [ ] In-memory flag store
* [ ] Persistent storage backend (JSON, SQLite, Postgres)
* [ ] REST + gRPC APIs
* [ ] Client SDKs (Go, JS/TS, Python)
* [ ] Web UI for managing flags
* [ ] CLI tool (`galfctl`)
* [ ] Docker + Kubernetes deployments
* [ ] Authentication / API keys

If you'd like to help build any of these, open an issue or PR!

---

## 🧩 Design Philosophy

* **Small, focused core** → easy to read, easy to fork
* **Minimise dependencies**
* **Predictable API** → ideal for microservices
* **Performant** → Go routines + lightweight HTTP stack
* **Testable** → modular internal packages

---

## 📄 License

This project is licensed under the **MIT License**.
See [LICENSE](LICENSE) for the full text.

---

## ✉️ Contact

Created by **Brett Earle**
GitHub: [@brettearle](https://github.com/brettearle)

---

[go]: https://golang.org
[go-badge]: https://img.shields.io/badge/go-1.20%2B-blue.svg
[license]: https://opensource.org/licenses/MIT
[license-badge]: https://img.shields.io/badge/license-MIT-green.svg

