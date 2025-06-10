WIP
---

Gnoserve - prototyping gno-frame design


BACKLOG
-------
- [ ] extend gno-cap to be usable
  - [ ] use gno js client to test rendering larger SVGs
  - [ ] add API for users to change pixel colors

- [ ] test/fix json+ld codefence support - does it work properly?
- [ ] remove all hardcoded configs in this repo - like starting realm path - support a config file or env var

 
  DONE
----
- [x] (try to use jsonld) Build image/object index on chain - add widget to reference on-chain data
- [x] 'reddit place' demo - try 250*250 bmp grid - png rendering collaborative drawing app
- [x] implement RSS feed using r/gnoserve posts as a feed

ICEBOX
------
- [ ] consider extracting 'gnomark' into a separate repo 
- [ ] build a template mechanism that depends on functions deployed to gno.land
   - try out gno functions (call out to realm to render template) - as MD extensions
- [ ] Consider Refactor could we depend on gnoweb and/or gnodev in a better way?
- [ ] support realm index/search
- [ ] add a means to let users add gnoweb endpoints to an onchain registry
- [ ] refine dependencies on gnodev - make a first-class api to host 3rd party plugins
- [ ] deploy a compatible interface for gno-mark widget system to gno.and and add a registry
