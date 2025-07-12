# Gnoserve

A forkable repo for 3rd party gno.land / gnoweb rendering.

## Motiviation

Gno.land rightly restrices many features of a web-page like javascript and css settings.
Gnoserve provides a foundation to provide custom 'alternate' routes and to extend existing rendering functionality

Examples:
- serving an SVG image at a new route /svg/r/your-realm-svg
- serving an RSS Feed at a new route /rss/r/your-realm-svg
- Enabling other Goldmark extentions like https://github.com/abhinav/goldmark-mermaid
- Writing your own Goldmark exention see [goldmark/codefence](https://github.com/allinbits/gnoserve/blob/task/gnodev_prototype_v1/gnomark/codefence.go)

## Status

Currently gno.land / gnoweb is undergoing a refactor.
We expect this will make it much easer to host a fork of gnoweb.

## Workarounds - other ways to extend

For now the best way to extend is by running your own webserver that uses the gno.land RPC calls to interact.

## Contributing

Open an issue to ask a question or discuss possible integrations.

Will likely see a full rewrite of this repo in the near future to more cleanly integerate with gnodev / gnoweb
