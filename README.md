# Distributed Commit Log (Go + TDD)

This project is a **distributed, append-only commit log** built from scratch in **Go**, designed as both a learning exercise and a production-grade infrastructure component. It mirrors the architecture of systems like Kafka, NATS Streaming, and WAL implementations in modern databases.

---

## Features (in progress)

* [x] In-memory log for fast prototyping and tests
* [x] File-backed `Store` with length-prefixed framing
* [ ] Protobuf-encoded `Record` persistence
* [ ] Indexed segments with offset-to-position mapping
* [ ] Segment rotation and compaction (planned)
* [ ] Multi-segment `Log` orchestration
* [ ] gRPC and HTTP API layers
* [ ] TLS-secured transport with authentication
* [ ] Manual leader-based clustering (Raft optional)

---

## Concepts Covered

* Append-only storage and data immutability
* File I/O with `os.File`, buffering, and flushing
* Binary encoding with `encoding/binary`
* Structured data with Protocol Buffers
* TDD methodology and `testify` assertions
* Mutexes and thread-safe concurrency

---

## Project Structure

```
internal/
  log/
    store.go         # Append-only file abstraction
    store_test.go    # Full coverage of append, read, flush, close
    ...              # (Coming soon: segment, index, log layers)
api/
  record.proto       # Protobuf schema for `Record`
  record.pb.go       # Generated Go types for gRPC and disk encoding
```

---

## Running Tests

Ensure you have Go installed (Go 1.22+).

```bash
go test ./internal/log -v
```

---

## Protobuf Code Generation

```bash
# Install plugins if not already installed
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Generate Go code from .proto definition
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       api/record.proto
```

---

## Roadmap

* [x] Milestone 0 – In-Memory Log
* [x] Milestone 1 – File-backed Store (framed append + read)
* [ ] Milestone 2 – Index (offset → position mapping)
* [ ] Milestone 3 – Segment (store + index per chunk)
* [ ] Milestone 4 – Log abstraction (multi-segment controller)
* [ ] Milestone 5 – gRPC, TLS, and clustering

---

## Author

Kimba SABI N'GOYE

---

## License

MIT License
