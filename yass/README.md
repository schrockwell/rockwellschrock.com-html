# yass - Yet Another Static Site builder

YASS requires a specific directory structure:

    _out/             <-- generated site ends up here
    site/             <-- site root
      _templates/     <-- flat dir of templates
        root.html     <-- the root page layout

It recursively walks the `site` directory, copying files over to the `_out` directory. `.html` files are evaluated as templates, making available all templates from `_templates`. All other file types are just copied.
