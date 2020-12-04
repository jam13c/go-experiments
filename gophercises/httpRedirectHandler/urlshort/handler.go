package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if p, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, p, http.StatusMovedPermanently)
		}
		fallback.ServeHTTP(w, r)
	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(paths)
	return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse provided json and then return blah blah
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(paths)
	return MapHandler(pathMap, fallback), nil
}

func parseYaml(ymlBytes []byte) ([]pathToURL, error) {
	var paths []pathToURL
	err := yaml.Unmarshal(ymlBytes, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func parseJSON(jsonBytes []byte) ([]pathToURL, error) {
	var paths []pathToURL
	err := json.Unmarshal(jsonBytes, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func buildMap(paths []pathToURL) map[string]string {
	var result = make(map[string]string)
	for _, pu := range paths {
		result[pu.Path] = pu.URL
	}
	return result
}

type pathToURL struct {
	Path string `yaml:"path",json:"path`
	URL  string `yaml:"url",json:"url"`
}
