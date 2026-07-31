# Godot UI and accessibility

Inspect the installed Godot version, Control hierarchy, Containers, theme resources, Input Map, focus neighbors, localization, and stretch settings.

## Patterns

- Build layout with Control nodes and Containers rather than manual per-resolution coordinates.
- Use Themes and theme variations for consistent visual states.
- Define focus mode, initial focus, directional neighbors, and focus restoration.
- Keep gameplay actions separate from built-in UI navigation actions.
- Use CanvasLayer according to HUD and world-space needs.
- Test translated text, font fallback, layout direction, and text expansion.
- Store reduced-motion, captions, contrast, text scale, and remapping settings in a durable settings profile.

## Testing

Navigate entire flows with keyboard or controller, inspect focus after hiding nodes, and test all supported stretch modes and safe areas on exported targets.

## Official references

- [User interface](https://docs.godotengine.org/en/stable/tutorials/ui/index.html)
- [Keyboard/controller navigation](https://docs.godotengine.org/en/stable/tutorials/ui/gui_navigation.html)
- [Internationalization](https://docs.godotengine.org/en/stable/tutorials/i18n/index.html)
- [Multiple resolutions](https://docs.godotengine.org/en/stable/tutorials/rendering/multiple_resolutions.html)
