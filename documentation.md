# 📝 Design Document: Distributed Commit Log in Go

| Field            | Value                                |
| ---------------- | ------------------------------------ |
| **Author**       | Kimba SABI N'GOYE                    |
| **Reviewer(s)**  | TBD                                  |
| **Created**      | 2025-05-31                           |
| **Last Updated** | 2025-05-31                           |
| **Status**       | In Progress                          |
| **Version**      | 1.0.2                                |
| **Project**      | Distributed Commit Log (Go + TDD)    |
| **Milestone**    | Milestone 2 – Store + Record Framing |

---

## 1. **Overview**

This document outlines the design of a **Distributed Commit Log** built in **Go** using **test-driven development (TDD)**. The system serves as a foundational backend infrastructure project and a guided learning path in systems programming, covering persistence, networking, and distributed coordination.

---

## 2. **Goals**

* Build an append-only, immutable commit log
* Support durable persistence and offset-based reads
* Expose gRPC and HTTP APIs
* Enable clustering with replication
* Provide TLS and authentication mechanisms
* Use idiomatic, modular, and testable Go code

---

## 3. **Non-Goals**

* Log compaction or deletion policies
* Advanced consensus (full Raft, Paxos)
* Multi-tenant ACL beyond simple role mapping
* Observability and metrics (initially)

---

## 4. **Use Cases**

* Event sourcing / stream processing pipelines
* Audit logging and traceability
* Message brokering in microservice architectures

---

## 5. **System Architecture and Components**

### Core Components

| Component | Responsibility                                              |
| --------- | ----------------------------------------------------------- |
| `Record`  | The data stored in the log (value and offset).              |
| `Store`   | Manages the binary file where records are written and read. |
| `Index`   | Maps logical offsets to physical file positions.            |
| `Segment` | Groups a `Store` and an `Index`, bounded by size/offset.    |
| `Log`     | Coordinates all segments and handles client operations.     |

---

## 6. **Milestone Breakdown**

### ✅ Milestone 0: In-Memory Log

* `[]Record` slice protected by mutex
* Supports `Append()` and `Read()`
* Used for early tests and API stubbing

### ✅ Milestone 1: Store Layer (File-Backed, Length-Prefixed)

* Implement `Store.Append([]byte) (n, pos uint64, err)`
* Implement `Store.Read(pos uint64) ([]byte, error)`
* Use `bufio.Writer`, `binary.BigEndian`, and file sync
* Introduce framing format: \[8-byte length]\[payload]
* Write TDD tests for Append, Read, ReadAt, and Close
* ✅ Completed and committed

### 🔜 Milestone 2: Index Layer

* Fixed-size entries: (offset, position)
* Index maps logical record offset → file position
* Binary encoding for efficient lookups

### 🔜 Milestone 3: Segment

* Combines Store + Index
* Tracks `baseOffset`, `nextOffset`
* Handles appending and segment size boundaries

### 🔜 Milestone 4: Log Abstraction

* Orchestrates multiple segments
* Append dispatches to active segment
* Read locates correct segment by offset

### 🔜 Milestone 5+: API Layer, TLS, Replication

* Add gRPC and HTTP APIs
* Secure communication and basic access control
* Clustered replication (manual leader or Raft)

---

## 7. **Serialization Format**

We use **Protocol Buffers** to define and encode the `Record`. This ensures compatibility with gRPC APIs and efficient binary representation on disk.

### record.proto

```proto
syntax = "proto3";
package api;
message Record {
  bytes value = 1;
  uint64 offset = 2;
}
```

### Go Usage

```go
proto.Marshal(&record)     // to []byte
proto.Unmarshal(data, &r)  // from []byte
```

---

## 8. **Testing Strategy**

| Layer   | Tests                                                     |
| ------- | --------------------------------------------------------- |
| Record  | Proto serialization/deserialization                       |
| Store   | Append, Read, ReadAt, length framing, offset correctness  |
| Index   | Append, Read, offset mapping, boundary checks             |
| Segment | End-to-end segment lifecycle with store and index         |
| Log     | Multi-segment append and read, segment rotation, failures |

---

## 9. **Go Features Used**

| Area          | Go Feature                                          |
| ------------- | --------------------------------------------------- |
| File I/O      | `os.File`, `bufio.Writer`, `ReadAt`, `Seek`, `Sync` |
| Serialization | `encoding/binary`, `protobuf`                       |
| Concurrency   | `sync.Mutex`, `sync.RWMutex`                        |
| Struct Design | Composition, encapsulation, method receivers        |
| Testing       | `testing` package, `require.*` from `testify`       |

---

## 10. **Next Actions**

* Integrate `proto.Marshal` in `Store.Append` for writing `Record`
* Implement matching deserialization in `Store.Read`
* Write tests for protobuf-encoded records
* Proceed to Index abstraction
