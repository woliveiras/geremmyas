---
description: "Use when writing or reviewing ChromaDB collections, clients, embeddings, metadata filters, persistence, and RAG retrieval code."
applyTo: "**/chroma/**/*, **/chromadb/**/*, **/rag/**/*, **/retrievers/**/*, **/embeddings/**/*, **/*chroma*.*, **/*retriever*.*, **/*embedding*.*"
---

# ChromaDB Conventions

- Choose the client mode deliberately: `PersistentClient` for embedded local
  persistence, `HttpClient` for a separate Chroma service.
- Use stable collection names and stable document IDs so updates are
  deterministic.
- Keep the embedding function and embedding model consistent for the lifetime
  of a collection unless you plan a re-embedding migration.
- Store metadata needed for filtering, ownership, freshness, and debugging at
  insert time.
- Use metadata filters to narrow retrieval before ranking when the user,
  tenant, document type, or time range is known.
- Do not recreate or delete collections as a cleanup shortcut in persistent or
  shared environments.
- Add retrieval tests that check inserted IDs, metadata filters, expected
  matches, and no-match behavior.
- Back up persistent collection data and document how to rebuild it from source
  documents and embeddings.
- Keep chunking, embedding, storage, and retrieval evaluation as separate steps
  when debugging RAG quality.
