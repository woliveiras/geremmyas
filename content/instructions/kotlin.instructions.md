---
description: "Kotlin Android code quality based on official Android architecture guides, Kotlin conventions, and Jetpack Compose best practices."
applyTo: "**/*.kt"
---

# Kotlin Android Conventions

Write idiomatic Kotlin that reduces mutable state, encapsulates finite states
with sealed classes, models data with data classes, exposes only immutable types,
and lets the UI react to observable state.

## Architecture

- Use **layered architecture**: UI layer, domain layer (optional), data layer
- Every screen receives state and emits events (UDF)
- Each screen has **one primary `uiState`** exposed as `StateFlow`
- `ViewModel` coordinates use cases and updates `uiState`; never holds
  `Context`, `Activity`, or `Resources`
- Data layer (network, database, preferences) lives behind **repositories**
- UI never references Retrofit, Room, DataStore, or services directly
- Single-activity with Navigation is the default for new apps
- Collect state in UI with `collectAsStateWithLifecycle()`

## State & Concurrency

- Expose `StateFlow<UiState>`, never `MutableStateFlow` or `LiveData`
- Use `sealed interface` for UI states (Loading, Content, Error)
- Use `viewModelScope.launch` for ViewModel coroutines
- Use `Flow` for streams; `suspend` for single-shot operations
- Apply `flowOn(ioDispatcher)` in repositories, not in ViewModels
- Use `catch` operator for Flow error handling
- Never use `GlobalScope`; scope all work to lifecycle owners
- Use `StateFlow` for state, `SharedFlow` for one-shot events

## Kotlin Idioms

- Prefer `val` and read-only collections; mutate only when necessary
- Use `data class` for models, DTOs, UI state, and parameter objects
- Use `sealed interface`/`sealed class` for exhaustive `when` without `else`
- Use extension functions for domain semantics; avoid hiding effects
- Avoid `!!`; prefer `requireNotNull`, `?:`, `?.let`, `as?`
- Use `require` / `check` for preconditions and invariants
- Use `copy()` for immutable state updates
- Prefer functional collection operations (`map`, `filter`, `first`) over
  imperative loops with mutation
- Use scope functions with clear intent: `apply` for configuration, `let` for
  null chains, `also` for side effects; never nest deeply

## Dependency Injection (Hilt)

- `@HiltViewModel` with `@Inject constructor` for ViewModels
- Repositories: `@Singleton` scope, bound via `@Binds` in `@Module`
- Use **constructor injection** exclusively; never field injection
- Inject `CoroutineDispatcher` for testability (not hardcoded `Dispatchers.IO`)
- Use `@Provides` only for third-party or complex object creation

## Jetpack Compose

- Screen composables receive `viewModel` parameter (default `hiltViewModel()`)
- Extract private helper composables for complex UI sections
- Use `Modifier` as first optional parameter in reusable composables
- Follow Material 3 design patterns and theme tokens
- Keep composables stateless; hoist state to caller or ViewModel
- Use `remember` / `rememberSaveable` for local ephemeral state only
- Preview composables with `@Preview` using hardcoded sample data

## Room & Persistence

- DAOs return `Flow<List<Entity>>` for observable queries
- Use `suspend` for write operations (insert, update, delete)
- Define entities with `@Entity`; use `@PrimaryKey(autoGenerate = true)` or
  explicit IDs
- Soft-delete with `deletedAt` timestamp; filter with
  `WHERE deletedAt IS NULL`
- Write migrations for schema changes; test migrations with
  `MigrationTestHelper`
- Use DataStore for preferences; avoid SharedPreferences in new code

## Networking

- Retrofit for REST in Android-only apps; Ktor Client for KMP
- Define API interfaces with `suspend` return types
- Map DTOs to domain models at the repository boundary
- Keep network models separate from UI/domain models
- Use sealed results for network responses (Success, Error, Loading)

## Naming

- `PascalCase` for classes, interfaces, sealed types, composables
- `camelCase` for functions, properties, local variables
- `UPPER_SNAKE_CASE` for compile-time constants (`const val`)
- Packages: lowercase, no underscores
- Test names: backtick-delimited descriptive phrases

## Error Handling

- Fail early at boundaries with `require` / `check`
- Use sealed result types instead of exceptions for expected failures
- Catch specific exceptions; never catch `Exception` or `Throwable` broadly
- Log errors with context (not PII); use Timber for structured logging
- Use `runCatching` sparingly; prefer explicit error modeling

## Testing

- Backtick-delimited test names: `` fun `returns error when name is blank`() ``
- Fake repositories and data sources for ViewModel tests
- Use `runTest` for coroutine tests
- Test Flow emissions with Turbine or `first()` on deterministic fakes
- Use Truth for assertions
- Robolectric for Android-dependent unit tests without emulator
- Espresso for critical UI integration tests
- Never use `Thread.sleep` or `delay` in tests

## Build & Tooling

- Gradle Kotlin DSL with version catalogs (`libs.versions.toml`)
- Use KSP instead of kapt when supported (Hilt, Room)
- Android Lint + detekt + ktlint for static analysis
- `.editorconfig` committed and enforced
- R8 enabled in release builds with `isMinifyEnabled = true`

## Anti-patterns

- `LiveData` in new code (use `StateFlow`)
- `!!` without prior proof of non-null
- `var` exposed in public APIs
- `GlobalScope.launch` or unscoped coroutines
- Field injection (`@Inject lateinit var`)
- Nested scope functions reducing readability
- Fat ViewModels with business logic (extract use cases)
- Mutable collections in public API surface
- `SharedPreferences` in new code (use DataStore)
- Catching generic `Exception` without rethrowing
