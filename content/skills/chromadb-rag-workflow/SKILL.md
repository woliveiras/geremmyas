---
name: chromadb-rag-workflow
description: "Design or review a ChromaDB-backed RAG workflow. Use when: creating collections, choosing embedding functions, chunking documents, writing retrievers. Do not use: for general LLM integration, non-RAG patterns."
---


# ChromaDB RAG Workflow

Build Chroma-backed retrieval with stable collections, metadata, and evaluation.

## Process

1. Identify source documents, update cadence, tenancy/security boundaries, and
   retrieval success criteria.
2. Choose `PersistentClient` or `HttpClient` based on deployment shape.
3. Define collection naming, embedding function/model, document IDs, and
   metadata schema.
4. Choose chunking and normalization before ingestion.
5. Ingest with stable IDs and metadata needed for filters and debugging.
6. Add retrieval tests for expected matches, metadata filters, no-match cases,
   and stale/deleted document behavior.
7. Document backup, rebuild, and re-embedding strategy.
8. Keep retrieval metrics or qualitative eval examples close to the feature.

## Rules

- Do not change embedding model for an existing collection without a
  re-embedding plan.
- Do not delete/recreate persistent collections as normal cleanup.
- Do not rely on vector similarity alone when tenant, document type, or
  freshness filters are known.
- Keep source documents recoverable so vector data can be rebuilt.

## Output

- Collection and metadata design
- Ingestion/retrieval flow
- Persistence and backup plan
- Retrieval evaluation checklist
