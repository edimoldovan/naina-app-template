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

Font size, font weight, font family, and font color are all set by `elements.css` and `type.css` on semantic HTML elements (`h1`–`h6`, `p`, `a`, `span`, `label`, etc). Component CSS must not redeclare these properties unless overriding them for a specific, justified reason.

**Rules:**
- Never set `font-size` in a component style. The element's base size comes from `elements.css`.
- Never set `font-weight` in a component style unless you are intentionally deviating from the element's default weight (e.g. making a `p` bold inside a card).
- Never set `font-family` in a component style unless using `--font-family-display` for a display heading that differs from the body font. Never hardcode a font name.
- Never set `color` on text in a component style to establish a default. Only set `color` when overriding — e.g. muting a paragraph to `--color-neutral-5`, or applying an accent color to a label.
- Never use raw values for any of the above. Font sizes must use `--font-size-*` variables if an override is needed. Colors must use `--color-neutral-*` or `--color-color-*` variables.

**What "override" means:** if `elements.css` already sets the right value on that element in that context, do not repeat it. Only write a rule when you need a different value from the default.

## CSS structure

- Split CSS into one `<style>` block per component, labelled with a comment.
- Link `reset.css`, `colors.css`, `type.css`, `space.css`, `grids.css`, `elements.css`, and `components.css` via `<link>` tags in production templates. Only embed them as `<style>` blocks when delivering a standalone HTML file.
- Component `<style>` blocks follow after the base stylesheet links.
- Use CSS nesting (`&`) and child combinators (`>`) inside component blocks. Flat selector lists are not permitted inside a component block — nest instead.
- Prefer semantic HTML elements (`article`, `header`, `nav`, `ul`, `li`, etc.) over generic `div` wrappers so that nesting selectors follow the document structure.

## Component naming

- Assign a single class to the component root (e.g. `.site-nav`, `.listings-section`). Style everything inside it through nested `>` child selectors targeting semantic elements or minimal modifier classes (e.g. `.active`, `.hidden`, `.bg--1`).
- Do not use BEM double-underscore classes (`__`) when nesting eliminates the need for them. Only add a class to an element when a structural selector would be ambiguous or fragile.
- No utility classes beyond those defined in `grids.css` (`.full`, `.wide-1`, `.wide-2`, `.hol-wrapping-row`, `.hol-column`, `.spaced`).

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


- `select`, `input`, and `textarea` base styles come from `elements.css`. Do not override them unless strictly necessary.
- Use `--border-radius` for all border radii.

## What to always ask before building

- Which sections need full bleed (`.full`) vs standard column?
- Are there any colors in the design that do not map to the existing palette? If so, flag — never invent new variables.
- Are there any button variants needed beyond `.primary`, `.secondary`, `.tertiary`, `.delete`? If so, flag — never create them silently.