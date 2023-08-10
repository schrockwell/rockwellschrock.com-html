# rockwellschrock.com - deconstructed

The site is built using YASS, a little standalone Go utility. See `yass/` directory README for more documentation.

## Dependencies

- Go 1.20
- fswatch (optional)
- direnv (optional)

## Local Site Development

```sh
# Install dependencies (first time) and 'yass' binary
bin/setup.sh

# Start server on http://localhost:3000 and watch for file changes in site/*
bin/server.sh

# Or just build the site
yass build
```

## Production Deployment

```sh
bin/deploy.sh
```
