package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string
	URL string
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}


func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

		var pathUrls []pathUrl

		err := yaml.Unmarshal(yml, &pathUrls)

		if err != nil {
			return nil, err
		}

	pathsToURLS := make(map[string]string)

	for _, pu := range pathUrls {
		pathsToURLS[pu.Path] = pu.URL
	}

	return MapHandler(pathsToURLS, fallback), nil

}