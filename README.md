# Gnoserve

A forkable repository for third-party rendering of gno.land / gnoweb.

## Motivation

gno.land appropriately restricts many web page features, such as JavaScript and CSS configurations.  
Gnoserve offers a foundation for building custom "alternate" routes and extending rendering functionality.

**Examples:**
- Serve SVG images at a new route: `/svg/r/your-realm-svg`
- Serve RSS feeds at a new route: `/rss/r/your-realm-svg`
- Enable additional Goldmark extensions like [goldmark-mermaid](https://github.com/abhinav/goldmark-mermaid)
- Create your own Goldmark extension (see [goldmark/codefence](https://github.com/allinbits/gnoserve/blob/task/gnodev_prototype_v1/gnomark/codefence.go))

## Status

gno.land and gnoweb are currently undergoing refactoring.  
This is expected to make hosting a fork of gnoweb much easier.

## Workarounds & Alternative Extension Methods

For now, the most effective way to extend functionality is to run your own web server that interacts with gno.land via RPC calls.

Look for an upcoming rewrite of this repository, to better integrate with gnodev/gnoweb using a more extensible upstream library.

## Contributing

Feel free to open an issue to ask questions or discuss possible integrations.
