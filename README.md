# rockwellschrock.com - deconstructed

The site is built using YASS, a little standalone Go utility. See `yass/` directory README for more documentation.

## Dependencies

- Go 1.20 - install with `asdf install`
- fswatch (optional) - install with Homebrew
- direnv (optional)

## Local Site Development

```sh
# Install YASS
cd yass
go install
cd ..

# Start server on http://localhost:3000 and watch for file changes in site/
bin/server.sh
```

## Production Deployment

```sh
bin/build.sh
```
