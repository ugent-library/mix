[![Go Reference](https://pkg.go.dev/badge/github.com/ugent-library/mix.svg)](https://pkg.go.dev/github.com/ugent-library/mix)

# ugent-library/mix

Package mix implements convenience methods to integrate the Laravel Mix
asset bundler in your Go project.

For more information on Laravel Mix, visit the [Laravel Mix](https://laravel-mix.com/) website.

## Install

```sh
go get -u ugent-library/mix
```

## Example

Setup your `webpack.mix.js`:

```js
const mix = require('laravel-mix')
mix.sass('./assets/css/app.scss', 'css')
mix.setPublicPath('./static')
// ...
```

And in your Go app:

```go
    // setup file server
    mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    // setup assets
    assets, err := mix.New(mix.Config{
        ManifestFile: "static/mix-manifest.json",
        PublicPath:   "/static/",
    })
    if err != nil {
        log.Fatal(err)
    }

    // make asset path helper available to templates
    funcs := template.FuncMap{
        "assetPath": assets.AssetPath,
    }

    // ...
```