# Godot persistence

Inspect the installed Godot version, target platforms, current `user://` layout, resource usage, and serialization format.

## Patterns

- Store user-writable data under `user://`, not `res://`.
- Serialize plain dictionaries, arrays, primitives, and stable IDs through a versioned adapter.
- Validate parsed data before creating or mutating nodes.
- Keep `Resource` files for authored data unless a deliberate runtime-save format is designed.
- Use temporary and backup paths when the platform supports a safe replacement flow.
- Treat mobile lifecycle, web storage, and console platform behavior as target-specific acceptance.

## Security

Do not execute saved text as code or instantiate arbitrary classes from untrusted fields. Encryption keys shipped with a client cannot provide strong secrecy against the device owner.

## Official references

- [Saving games](https://docs.godotengine.org/en/stable/tutorials/io/saving_games.html)
- [FileAccess](https://docs.godotengine.org/en/stable/classes/class_fileaccess.html)
- [JSON](https://docs.godotengine.org/en/stable/classes/class_json.html)
- [Data paths](https://docs.godotengine.org/en/stable/tutorials/io/data_paths.html)
