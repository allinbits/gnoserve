## `gnoserve`: Markdown Extensions for Gno.land

Interactive Extensions for gnodev Markdown Rendering

Status: Pre-Alpha

KNOWN ISSUES
------------
- lots of copy/paste from gnodev - need to refactor upstream interfaces
- support loading in staging or localdev - currently only READONLY ux
- does not expose RPC for the local test chain
- lacking a standard for *extra* metadata (i.e. structured data)
- lacking *extra* API routes to support RSS or other standards
  
https://gnoserve.fly.dev/ <- try it out on fly.io

Read More [on notion](https://allinbits.notion.site/GnoFrame-1e0302c277ee80c3b9e0e1c41eebbb65).

### Contribute

Feel free to open an PR to ask for potential features / support

### Motivation

We want to push Markdown beyond static text — without losing its simplicity.
By extending gnodev's rendering engine, we enable interactive models, visual editors,
and programmable UIs to live directly inside realm Markdown.

This lets Gno developers:

- Keep content readable and versionable in plain text
- Embed interactive logic models (e.g., Petri nets, grids, sprites)
- Preview and test rich simulations without leaving the realm

⚙️ Code-as-content meets content-as-interface.

### What's New in `gnoserve` vs `gnoweb`

`gnoserve` introduces several enhancements over `gnoweb`, making it more powerful and flexible for Markdown rendering:

1. **Gno-Mark Support**:  
   A new `<gno-mark>` tag wraps all custom Markdown elements.

2. **Additional Plugins**:  
   We have enabled several new plugins to extend Markdown functionality:
   - **`img64` Plugin**:  
     This plugin allows images to be embedded as base64-encoded data, with caching tied to the page's lifetime.
     This solves common CDN-related issues and ensures reliable image delivery.
   - **`mermaid` Plugin**:  
     We are excited to use Mermaid for creating diagrams and visualizations directly in Markdown.
     This makes it easier to include rich, interactive diagrams without external tools.

3. **Improved Caching and Rendering**:  
   The `img64` plugin ensures that images are cached efficiently, reducing dependency on external resources and improving page load times.

These features make `gnosrv` a more robust and developer-friendly tool for creating interactive and visually rich Markdown content.
