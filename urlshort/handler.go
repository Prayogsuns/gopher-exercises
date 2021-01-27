package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
	// "fmt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		path_url := pathsToUrls[r.URL.Path] 
    if path_url != "" {
      http.RedirectHandler(path_url, 301).ServeHTTP(w, r) 
		} else {
      fallback.ServeHTTP(w, r)
		}
	}
	
	return myHandler
}

var yaml_data = map[string]string{}

type YamlD struct {
  Path string `yaml:"path"`
	Url string  `yaml:"url"`
}
var yd = []YamlD{}
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
  err := yaml.Unmarshal(yml, &yd)
  if err != nil {
		return nil, err
	}
	for _, data := range yd {
		// fmt.Println(data)
    yaml_data[data.Path] = data.Url
	}
	mapHandler := MapHandler(yaml_data, fallback)
	return mapHandler, nil
}