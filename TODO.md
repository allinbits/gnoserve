WIP
---

Gnoserve initial setup and features

BACKLOG
-------
- [ ] build a template mechanism that depends on functions deployed to gno.land
    - try out gno functions (call out to realm to render template) - as MD extensions
     
- [ ] remove all hardcoded configs like starting realm path - support a config file or env var
 
- [ ] Build image/object index on chain - add widget to reference on-chain data
    - can/should we allow a realm to reference other realms' data via widget? 
    - this would be a tag like: { "gnoMark": "render" realm: "r/otherrealm", path: "/path/to/data" }
    - path would be used for args
 
- [ ] add whitelist of html tags we allow inside <gno-mark> blocks
    - can we use the existing goldmark sanitizer?

- [ ] try 250*250 bmp grid - png rendering collaborative drawing app
 
  DONE
----
- [x] support image hosting use case from realm /r/stackdump/www /r/stackdump/bmp:filename.jpg
    - Used the img64 markdown plugin to host images embedded in page
- [x] add GnoFrame tag support - roughly equivalent to farcaster frames
- [x] consider implementing a frontend-only solution to template rendering
  we could still lint the json body on the server side and then render it on the client side
  https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_custom_elements#implementing_a_custom_element

ICEBOX
------
- [ ] Consider Refactor could we depend on gnoweb and/or gnodev in a better way?
- [ ] support realm index/search
- [ ] add a means to let users add gnoweb endpoints to an onchain registry
- [ ] refine dependencies on gnodev - make a first-class api to host 3rd party plugins
- [ ] deploy a compatible interface for gno-mark widget system to gno.and and add a registry