# naina-app-template

Go web app template. Follow KISS — small, direct code, no speculative abstractions.

## Layout
- `cmd/` — entrypoint (`main.go` calls `handler.InitTemplates()` then starts the server).
- `internal/config/` — env/config helpers (e.g. `config.IsDev()`).
- `internal/router/` — HTTP routes.
- `internal/handler/` — HTTP handlers + template rendering (`render.go`).
- `internal/handler/html/` — embedded templates:
  - `pages/` — full HTML documents, one per route.
  - `partials/` — reusable `{{define "name"}}` fragments (head, nav, …).

## Rendering
- No layouts folder. Each page file is its own complete `<!DOCTYPE html>` document.
- Pages pull in partials via `{{template "name" .}}`.
- `InitTemplates()` parses every partial together with each page once at startup.
- In dev mode (`config.IsDev()`), templates are re-parsed from disk on every request for hot reload.
- Handlers call `render(w, "pagename", data)` — the key is the page filename without `.html`.

## Adding a page
1. Create `internal/handler/html/pages/foo.html` as a full HTML document.
2. Add a handler in `internal/handler/` that calls `render(w, "foo", data)`.
3. Register the route in `internal/router/`.

## Adding a partial
1. Create `internal/handler/html/partials/thing.html` with `{{define "thing"}}…{{end}}`.
2. Reference it from any page via `{{template "thing" .}}`. No other wiring needed.

## Build / run
- `./run.sh`

# Frontend styling rules

## Layout

- Every page has a single `<main class="grid">` at the top level. All content lives inside it.
- The grid is defined in `grids.css`. Default content width is the `standard` column. Children of `.grid` are placed there automatically.
- Use `.full`, `.wide-1`, `.wide-2` classes to escape the standard column for full-bleed or wider sections.
- Never use media queries. Use grid escape classes and `auto-fit`/`auto-fill` for responsive layout.

## Colors

- Only use variables defined in `colors.css`: `--color-neutral-1` through `--color-neutral-10` and `--color-color-0` through `--color-color-11`.
- Never use hex values, rgb, hsl, or any raw color anywhere in component CSS.
- The only exception is inside `colors.css` itself where oklch values are defined.
- Never add new color variables. Never rename existing ones. Never extend the color system.
- Light/dark mode is handled automatically by the existing `colors.css` media query. Do not add separate dark mode rules.

## Typography

- All font sizes use variables from `type.css`: `--font-size--2` through `--font-size-6`.
- All font families use `--font-family-1` (Manrope) for body/UI. Display headings may use a separate `--font-family-display` if defined.
- Never hardcode font sizes in px or rem directly.

## Spacing

- All spacing uses variables from `space.css`: `--space-3xs` through `--space-3xl`, including fluid step pairs like `--space-s-m`.
- `--border-radius` is always `1rem`, defined in the space block.
- Never hardcode spacing values in px or rem directly.

## Buttons

- Only four button classes are allowed: `.primary`, `.secondary`, `.tertiary`, `.delete`.
- These are defined in `elements.css`. Do not create custom button styles.
- No other button styling is permitted.

## No inline styles

- `style="..."` attributes are never allowed in HTML.
- The only exception is JavaScript setting properties dynamically at runtime (e.g. `transform` for animations, `outline` for programmatic highlights).

## CSS structure

- Split CSS into one `<style>` block per component, labelled with a comment.
- Embed `reset.css`, `colors.css`, `type.css`, `space.css`, `grids.css`, `elements.css`, and `components.css` as separate `<style>` blocks at the top of the file when delivering a standalone HTML file.
- Component `<style>` blocks follow after the base blocks.

## Component naming

- Use BEM-style class names: `.component-name__element`.
- No utility classes beyond those defined in `grids.css` (`.full`, `.wide-1`, `.wide-2`, `.hol-wrapping-row`, `.hol-column`, `.spaced`).

## Forms and inputs

- `select`, `input`, and `textarea` base styles come from `elements.css`. Do not override them unless strictly necessary.
- Use `--border-radius` for all border radii.

## What to always ask before building

- Which sections need full bleed (`.full`) vs standard column?
- Are there any colors in the design that do not map to the existing palette? If so, flag — never invent new variables.
- Are there any button variants needed beyond `.primary`, `.secondary`, `.tertiary`, `.delete`? If so, flag — never create them silently.


