import(
  "path/filepath"
  "io/ioutil"
  "os"
  "fmt"
  "net/http"
  "strings"
)

// The AssetHeaders struct is used to set Cache-Control headers to all GET and HEAD
// requests to /assets in production. Because these assets have digested names, we
// can set the cache time really high, and use this app as origin for a CDN.
type AssetHeaders struct {
}

func (s *AssetHeaders) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  
  // Ignore if not production
  if os.Getenv("GO_ENV") != "production" {
    next(rw, r)
    return
  }

  // Ignore all but GET and HEAD
  if r.Method != "GET" && r.Method != "HEAD" {
    next(rw, r)
    return
  }

  // Ignore everything not in /assets
  path := r.URL.Path
  if !strings.HasPrefix(path, "/assets") {
    next(rw, r)
    return
  }

  // Set asset caching to one year 
  rw.Header().Set("Cache-Control", "public, max-age=31536000")
  next(rw, r)
  
}