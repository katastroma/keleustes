# Keleustēs

Resolver interface for GitOps on Kubernetes. Defines the contract for answering:
**what resources should exist?**

## What This Is

A Go module containing an interface definition and its associated types. Not a
controller, not a server, not a CLI. It is the contract that resolvers
implement.

A resolver takes a source — a git repository, a branch, a path, and credentials
— and produces a resource inventory: the set of Kubernetes resources that should
exist according to that source.

## Why It Exists

Every GitOps system resolves sources into resources, but they all do it
internally, tightly coupled to their own reconcile and provisioning logic.
Keleustēs extracts the question into a standalone interface so that:

- Resolver implementations are independently testable
- Orchestrators can swap resolvers without changing their reconcile loop
- The ecosystem can converge on a shared contract instead of each project
  reinventing the same abstraction

## Ecosystem

Keleustēs is one of two GitOps primitive interfaces defined by
[katastroma](https://github.com/katastroma). The other is
[katartismos](https://github.com/katastroma/katartismos) (the provisioner
interface).

[Orpheus](https://github.com/katastroma/orpheus) is katastroma's reference
resolver implementation.
