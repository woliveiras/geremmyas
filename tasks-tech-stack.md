# Tasks: skills e instructions por stack

Este backlog planeja a proxima fase dos arquivos de `skills` e
`instructions` do `copilot-configs`.

O `tasks.md` anterior deve permanecer como historico concluido. Este arquivo
organiza novas tarefas para cobrir Node.js, Python, Go, infra, cloud, bancos,
RAG, LLMs e Supabase.

## Decisoes base

- Use `instructions` para regras recorrentes, curtas e especificas de uma
  tecnologia ou tipo de arquivo.
- Use `skills` para workflows acionaveis com começo, meio, fim, decisoes,
  artefatos e verificacao.
- Mantenha cada `instructions/*.instructions.md` curto e opinativo. Ele deve
  registrar o que nao e obvio para o agente, nao reescrever a documentacao da
  ferramenta.
- Use fontes oficiais ou primarias antes de escrever regras novas.
- Nao copie texto longo de documentacao externa; sintetize regras aplicaveis ao
  seu fluxo.

## Matriz instruction vs skill

| Area | Instruction | Skill | Decisao |
|------|-------------|-------|---------|
| NestJS | `nestjs.instructions.md` | - | Instruction apenas: regras de estrutura, DI, controllers, providers, pipes e validacao. |
| Fastify | `fastify.instructions.md` | - | Instruction apenas: regras de plugins, schemas, hooks, logger e erros. |
| Python | `python.instructions.md` | - | Instruction apenas: baseline curto de linguagem. |
| FastAPI | `fastapi.instructions.md` | - | Instruction apenas: regras de routers, dependencies, models e erros HTTP. |
| Pydantic | `pydantic.instructions.md` | - | Instruction apenas: contratos, validacao, serializacao, settings e boundaries em Python. |
| LangChain | `langchain.instructions.md` | - | Instruction apenas: regras recorrentes para chains, retrievers e tools. |
| LangGraph | `langgraph.instructions.md` | `langgraph-agent-design` | Ambos: instruction para codigo de graph, skill para desenho de agente/workflow. |
| LLM services | `llm-service.instructions.md` | `llm-integration-review` | Ambos: instruction para integracao recorrente, skill para revisar ou desenhar uma integracao completa. |
| Go | `go.instructions.md` | - | Instruction apenas: baseline curto de linguagem. |
| Echo | `echo.instructions.md` | - | Instruction apenas: regras de handlers, middleware, contexto e erros. |
| SQLite com Go | `sqlite.instructions.md` | - | Instruction apenas: regras de driver, migrations, pragmas e transacoes. |
| embed.FS | `go-embed.instructions.md` | - | Instruction apenas: regra especifica de empacotamento e acesso a assets. |
| Air | `air.instructions.md` | - | Instruction apenas: hot reload local, sem impacto em producao. |
| Docker | `docker.instructions.md` | - | Instruction apenas: regras de imagem e Dockerfile. |
| Docker Compose | `docker-compose.instructions.md` | - | Instruction apenas: regras de orquestracao local, redes, volumes e healthchecks. |
| GitHub Actions | `github-actions.instructions.md` | `ci-workflow` | Ambos: instruction para YAML, skill para criar ou depurar pipelines. |
| Google Cloud | `gcp.instructions.md` | `gcloud-operation` | Ambos: instruction para uso seguro de GCP, skill para operacoes com `gcloud`. |
| Terraform | `terraform.instructions.md` | `terraform-change` | Ambos: instruction para arquivos `.tf`, skill para fluxo `plan/apply` com aprovacao. |
| PostgreSQL | `postgres.instructions.md` | `postgres-query-review` | Ambos: instruction para schema/SQL, skill para revisar query, migration ou plano. |
| ChromaDB | `chromadb.instructions.md` | `chromadb-rag-workflow` | Ambos: instruction para uso recorrente, skill para desenhar ou revisar RAG. |
| Supabase | `supabase.instructions.md` | `supabase-workflow` | Ambos: instruction para cliente/RLS/policies, skill para schema, migrations e hardening. |

## Fontes base

- NestJS: https://docs.nestjs.com/
- Fastify: https://fastify.dev/docs/latest/
- LangChain: https://python.langchain.com/docs/
- LangGraph: https://langchain-ai.github.io/langgraph/
- Echo: https://echo.labstack.com/docs
- Go embed: https://pkg.go.dev/embed
- Air: https://github.com/air-verse/air
- Docker: https://docs.docker.com/
- GitHub Actions: https://docs.github.com/en/actions
- Google Cloud CLI: https://cloud.google.com/sdk/docs
- Terraform: https://developer.hashicorp.com/terraform/docs
- PostgreSQL: https://www.postgresql.org/docs/current/
- ChromaDB: https://docs.trychroma.com/
- Supabase: https://supabase.com/docs
- OpenAI API: https://platform.openai.com/docs/

## Task 1: Definir matriz instruction vs skill

- [x] Revisar a lista de tecnologias e classificar cada item como
  `instruction`, `skill` ou ambos.
- [x] Registrar que Supabase tera `supabase.instructions.md` e o skill
  `supabase-workflow`.
- [x] Registrar que integracoes com LLMs terao um instruction
  provider-agnostic chamado `llm-service.instructions.md`.
- [x] Registrar que skills serao criados apenas para workflows, nao para toda
  biblioteca.

Acceptance criteria:

- Cada tecnologia listada neste arquivo tem destino claro.
- A classificacao evita duplicar regras entre instructions e skills.
- O arquivo preserva a decisao de manter o backlog separado do `tasks.md`
  anterior.

## Task 2: Node.js APIs

- [x] Criar `project/.github/instructions/nestjs.instructions.md`.
- [x] Criar `project/.github/instructions/fastify.instructions.md`.
- [x] Incluir regras especificas de NestJS: controllers finos, providers,
  dependency injection, pipes, guards, interceptors, filters, config validation
  e DTO validation.
- [x] Incluir regras especificas de Fastify: plugins encapsulados, schemas de
  validation/serialization, hooks, decorators, logger e tratamento consistente
  de erros.
- [x] Atualizar `install.sh` para detectar `@nestjs/core` e `fastify`.
- [x] Atualizar `README.md` com os novos instruction files.

Acceptance criteria:

- As regras ficam focadas no que muda a implementacao em projetos reais.
- NestJS e Fastify nao duplicam regras genericas de TypeScript ou API security.
- A deteccao automatica instala os instructions corretos quando as
  dependencias aparecem no `package.json`.

## Task 3: Python AI/API

- [x] Manter `python.instructions.md` como baseline curto de linguagem.
- [x] Revisar `fastapi.instructions.md` para confirmar que ele contem apenas
  regras especificas de FastAPI.
- [x] Criar `project/.github/instructions/pydantic.instructions.md`.
- [x] Criar `project/.github/instructions/langchain.instructions.md`.
- [x] Criar `project/.github/instructions/langgraph.instructions.md`.
- [x] Criar `project/.github/instructions/llm-service.instructions.md`.
- [x] Atualizar `install.sh` para detectar `fastapi`, `pydantic`,
  `pydantic-settings`, `langchain`, `langgraph`, `openai`, `anthropic`,
  `google-genai`, `litellm` e dependencias equivalentes quando presentes em
  manifestos Python.

Acceptance criteria:

- LangChain cobre composicao, entradas/saidas tipadas, retrievers, tools,
  tracing/callbacks e ausencia de estado global escondido.
- Pydantic cobre `BaseModel`, `Field`, `ConfigDict`, `model_validate`,
  `model_dump`, settings e separacao entre API, persistencia, settings e
  dominio.
- LangGraph cobre state schema, nodes, edges, checkpoints, human-in-the-loop,
  resume/interrupt e idempotencia.
- `llm-service.instructions.md` cobre timeouts, retries, rate limits, structured
  outputs, secrets, redaction, observabilidade, custos e testes de contrato.

## Task 4: Go

- [x] Manter `go.instructions.md` como baseline curto da linguagem.
- [x] Criar `project/.github/instructions/echo.instructions.md`.
- [x] Revisar `project/.github/instructions/sqlite.instructions.md` com foco em
  `modernc.org/sqlite`.
- [x] Criar `project/.github/instructions/go-embed.instructions.md`.
- [x] Criar `project/.github/instructions/air.instructions.md`.
- [x] Atualizar `install.sh` para detectar `github.com/labstack/echo`,
  `modernc.org/sqlite`, uso de `//go:embed` e arquivos `.air.toml`.

Acceptance criteria:

- Echo cobre middleware, contexto, binding/validation, erros HTTP centralizados
  e graceful shutdown.
- SQLite cobre cuidados praticos com driver pure Go, migrations, pragmas,
  transacoes e testes.
- `go-embed.instructions.md` cobre caminhos, `//go:embed`, `fs.Sub` e ausencia
  de dependencias em caminhos de runtime.
- `air.instructions.md` fica limitado a hot reload de desenvolvimento.

## Task 5: Infra, CI/CD e Cloud

- [ ] Revisar `project/.github/instructions/docker.instructions.md`.
- [ ] Criar `project/.github/instructions/docker-compose.instructions.md`.
- [ ] Criar `project/.github/instructions/github-actions.instructions.md`.
- [ ] Criar `project/.github/instructions/gcp.instructions.md`.
- [ ] Criar `project/.github/instructions/terraform.instructions.md`.
- [ ] Atualizar `install.sh` para detectar Dockerfile, Compose files,
  `.github/workflows/*`, arquivos Terraform e configuracoes GCP relevantes.
- [ ] Atualizar `README.md` com a nova cobertura de infra.

Acceptance criteria:

- Docker cobre imagem, build, usuario nao-root, secrets e camadas.
- Docker Compose cobre healthchecks, redes, volumes, env files e uso local.
- GitHub Actions cobre permissions minimas, pinning, cache, matrix,
  concurrency, artifacts, secrets e OIDC.
- GCP cobre `gcloud` CLI, projeto ativo, account/auth, ADC, impersonation e
  scripts nao interativos.
- Terraform cobre `fmt`, `validate`, `plan`, state remoto, lockfile, modules,
  secrets, imports e `moved` blocks.

## Task 6: Data, RAG e BaaS

- [ ] Criar `project/.github/instructions/postgres.instructions.md`.
- [ ] Criar `project/.github/instructions/chromadb.instructions.md`.
- [ ] Criar `project/.github/instructions/supabase.instructions.md`.
- [ ] Atualizar `install.sh` para detectar dependencias ou diretorios
  relacionados a PostgreSQL, ChromaDB e Supabase.
- [ ] Atualizar `README.md` com os novos arquivos.

Acceptance criteria:

- PostgreSQL cobre migrations, constraints, indices, transacoes, query plans,
  connection pooling e cuidados com concorrencia.
- ChromaDB cobre persistent vs HTTP client, collections, embedding function,
  IDs estaveis, metadata filters, backups e avaliacao de retrieval.
- Supabase cobre anon key vs service role, RLS, policies, migrations, generated
  types, auth, storage e edge functions quando aplicavel.

## Task 7: Skills de workflow

- [ ] Criar `terraform-change`.
- [ ] Criar `gcloud-operation`.
- [ ] Criar `ci-workflow`.
- [ ] Criar `llm-integration-review`.
- [ ] Criar `langgraph-agent-design`.
- [ ] Criar `supabase-workflow`.
- [ ] Criar `postgres-query-review`.
- [ ] Criar `chromadb-rag-workflow`.
- [ ] Adicionar templates ou assets apenas quando o skill produzir artefatos
  duraveis.
- [ ] Atualizar `README.md` com os novos skills.

Acceptance criteria:

- Cada skill tem `name`, `description`, gatilhos claros, processo, regras e
  output esperado.
- Nenhum skill existe apenas para repetir boas praticas de uma biblioteca.
- Skills de infra deixam comandos destrutivos ou de apply explicitos e sujeitos
  a confirmacao humana.

## Task 8: Roteamento em AGENTS.md

- [ ] Atualizar `project/AGENTS.md` para explicar quando usar os novos
  instructions e skills.
- [ ] Evitar duplicar procedimentos que ja estejam nos skills.
- [ ] Garantir que workflows de LLM, Supabase, Terraform, GCP e CI tenham
  roteamento claro.

Acceptance criteria:

- `AGENTS.md` continua sendo contrato de operacao, nao documentacao longa.
- O agente consegue escolher entre instruction e skill sem depender de memoria
  externa.

## Task 9: Instalacao e desinstalacao

- [ ] Atualizar `install.sh` com os novos arquivos e regras de deteccao.
- [ ] Atualizar `uninstall.sh` se os novos arquivos precisarem ser removidos em
  instalacoes globais ou project-level.
- [ ] Garantir que arquivos opcionais so sejam instalados quando fizer sentido
  para o projeto detectado.
- [ ] Manter skills instalaveis quando forem de workflow geral, mesmo que a
  stack nao esteja presente no repo atual.

Acceptance criteria:

- `bash -n install.sh` passa.
- `bash -n uninstall.sh` passa.
- Smoke tests de instalacao cobrem pelo menos: NestJS, Fastify, FastAPI,
  LangGraph, Echo, Terraform, GitHub Actions, Supabase e Docker Compose.

## Task 10: Revisao final

- [ ] Conferir que todos os novos instructions sao curtos, especificos e nao
  obvios.
- [ ] Conferir que todos os novos skills tem workflow completo.
- [ ] Conferir que `README.md` lista todos os arquivos novos.
- [ ] Conferir que `install.sh` e `uninstall.sh` continuam validos.
- [ ] Conferir que nenhuma regra nova copia nomes, textos ou identidade de
  projetos externos.

Acceptance criteria:

- O repositorio fica pronto para suportar as stacks listadas sem inflar o
  contexto padrao dos agentes.
- As regras novas sao baseadas em fontes oficiais ou primarias.
- O backlog pode ser executado em iteracoes independentes por tecnologia.
