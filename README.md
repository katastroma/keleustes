# Keleustēs

Renderer interface for GitOps on Kubernetes. Defines the gRPC service contract
for producing Kubernetes manifests from source content.

A renderer takes source content and produces manifests — the set of Kubernetes
resources that should be applied to the cluster. Implementations are responsible
for rendering — helm template, kustomize build, or raw YAML.

## Why It Exists

Rendering is one of three fundamental GitOps operations (fetch, render,
provision). Keleustēs extracts the rendering contract so that:

- Renderer implementations are independently testable and deployable
- The orchestrator ([pharos](https://github.com/katastroma/pharos)) can call any
  renderer that satisfies the contract
- The ecosystem can converge on a shared contract instead of each project
  coupling rendering into a monolith

## Ecosystem

Keleustēs is one of three GitOps service interfaces defined by
[katastroma](https://github.com/katastroma):

- [naukleros](https://github.com/katastroma/naukleros) — retriever interface
- **keleustēs** (this) — renderer interface
- [katartismos](https://github.com/katastroma/katartismos) — provisioner
  interface

[Orpheus](https://github.com/katastroma/orpheus) is katastroma's renderer
implementation.
