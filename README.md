# Keleustēs

Resolution interfaces for GitOps on Kubernetes. Defines the contracts for
answering: **what resources should exist?**

## What This Is

A Go module containing two interface definitions and their associated types. Not
a controller, not a server, not a CLI. These are the contracts that resolution
implementations satisfy.

### Retriever

Takes an Application source and optional credentials, fetches the source, and
returns a readable filesystem (`fs.FS`). Implementations are responsible for
authentication and transport — git clone, OCI pull, or whatever the source
requires.

### Renderer

Takes a filesystem and produces a resource inventory
(`[]unstructured.Unstructured`). Implementations are responsible for manifest
rendering — helm template, kustomize build, or raw YAML.

## Why It Exists

Every GitOps system resolves sources into resources, but they all do it
internally, tightly coupled to their own reconcile and provisioning logic.
Keleustēs extracts the question into standalone interfaces so that:

- Implementations are independently testable
- Adapters can compose retrieval and rendering strategies freely
- The ecosystem can converge on shared contracts instead of each project
  reinventing the same abstractions

## Ecosystem

Keleustēs is one of two GitOps primitive interfaces defined by
[katastroma](https://github.com/katastroma). The other is
[katartismos](https://github.com/katastroma/katartismos) (the provisioner
interface).

[Orpheus](https://github.com/katastroma/orpheus) is katastroma's reference
implementation.
