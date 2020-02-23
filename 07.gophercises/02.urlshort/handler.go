package main

import (
	yaml "gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths (keys in the map) to their corresponding URL
// (values that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// TODO: 1.Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok { // if we can match a path...
			http.Redirect(w, r, dest, http.StatusFound) // 		redirect to it
			return
		} // else ...
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return an http.HandlerFunc
// (Which also implements http.Handler) that will attempt to map any paths to their corresponding URL.
// If the path is not provided in the YAML, then the fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
// - path:  /some-path
//   url:   http://www.some-url.com/demo
//
//  The only errors that can be returned all related to having invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via a mapping of paths to url.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: 1.Implement this...
	// 1. Parse the yaml somehow
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	// 2. Convert YAML array into map
	pathsToUrls := buildMap(pathUrls)
	// 3. return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
