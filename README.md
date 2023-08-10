# rockwellschrock.com - The Old Web Era

The site is generated using YASS, a little standalone Go utility. It utilizes the [text/template](https://pkg.go.dev/text/template) package from the standard library for rendering.

## Directory Structure

    _site/            <-- generated site ends up here (.gitignore'd)
    web/              <-- site content root
      _templates/     <-- flat dir of templates
        root.html     <-- the root page layout (required)
      index.html      <-- home page

YASS recursively walks the `web` directory, copying files over to the `_site` directory. The `.html` files are evaluated as individual templates, using all templates from `_templates` as a baseline. All other file extensions are copied directly.

## Dependencies

- Go 1.20
- fswatch (optional, for dev only)
- direnv (optional, for dev only)

## Local Development

```sh
# Install dependencies and 'yass' binary (first time)
scripts/setup.sh

# Start dev server on http://localhost:3000 and watch for file changes in site/*
scripts/dev.sh

# Or just build the site and serve it up
yass build && yass server
```

## Production Deployment

```sh
# Install 'yass' and build the site
scripts/deploy.sh
```
