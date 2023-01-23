// Package mix implements convenience methods to integrate the Laravel Mix
// asset bundler in your Go project.
//
// For more information on Laravel Mix, visit the [Laravel Mix] website.
//
// [Laravel Mix]: https://laravel-mix.com/
package mix

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type Config struct {
	// ManifestFile should contain the path to the manifest.mix.js file.
	ManifestFile string
	// PublicPath is the path under which asset files are served.
	PublicPath string
}

type Manifest map[string]string

func New(c Config) (Manifest, error) {
	data, err := os.ReadFile(c.ManifestFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't read mix manifest '%s': %w", c.ManifestFile, err)
	}

	manifest := make(map[string]string)
	if err = json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("couldn't parse mix manifest '%s': %w", c.ManifestFile, err)
	}

	if c.PublicPath != "" {
		for asset, p := range manifest {
			manifest[asset] = path.Join(c.PublicPath, p)
		}
	}

	return manifest, nil
}

// AssetPath returns the path to the asset given. An empty string and an error
// is returned if the path is not found in the manifest.
func (m Manifest) AssetPath(asset string) (string, error) {
	p, ok := m[asset]
	if !ok {
		return "", fmt.Errorf("asset '%s' not found in mix manifest", asset)
	}
	return p, nil
}
