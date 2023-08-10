# rockwellschrock.com - deconstructed

The site is built using YASS, a little standalone Go utility. See `yass/` directory README for more documentation.

## Dependencies

- Go 1.20
- fswatch (optional)
- direnv (optional)

## Local Site Development

```sh
# Install dependencies (first time) and 'yass' binary
scripts/setup.sh

# Start dev server on http://localhost:3000 and watch for file changes in site/*
scripts/dev.sh

# Or just build the site
yass build
```

## Production Deployment

```sh
scripts/deploy.sh
```

## Directory Structure

    _site/            <-- generated site ends up here (.gitignore'd)
    web/              <-- site content root
      _templates/     <-- flat dir of templates
        root.html     <-- the root page layout
      index.html      <-- home page

It recursively walks the `web` directory, copying files over to the `_site` directory. `.html` files are evaluated as templates, making available all templates from `_templates`. All other file types are just copied.
