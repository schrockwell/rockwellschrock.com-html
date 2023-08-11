# rockwellschrock.com - The Old Web Era

The site is generated using YASS (Yet Another Static Site), a little standalone Go utility. It utilizes the [text/template](https://pkg.go.dev/text/template) package from the standard library for rendering.

## Directory Structure

    _site/            <-- generated site ends up here (.gitignore'd)
    web/              <-- site content root
      _templates/     <-- flat dir of templates
        root.html     <-- the root page layout (required)
      index.html      <-- home page

YASS recursively walks the `web` directory, copying files over to the `_site` directory. The `.html` files are evaluated as individual templates, using all templates from `_templates` as a baseline. All other file extensions are copied directly.

## Commands

- `yass convert` - process and rename all JPEG images in `web/`, in-place
- `yass queen` - generate the site from `web/` and output to `_site/`
- `yass server` - run a local dev server that serves `_site/` on port 3000

## Dependencies

- Required
  - Go 1.20
  - Make
- Optional, for development only
  - direnv
  - fswatch

## Local Development

```sh
# Install dependencies and 'yass' binary (first time)
make install

# Start dev server on http://localhost:3000 and watch for file changes in site/*
make dev

# Or just build the site
make
```

## Production Deployment

```sh
# Install 'yass' and build the site
make deploy
```
