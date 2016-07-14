# dfalt
Simple library to use environment vars in golang flag statements or a default value when not set or invalid

## Example
```
func main() {
     var (
     	 httpAddr = flag.String("http", dfalt.EnvString("HTTP_ADDR", ":5050"), "HTTP service address")
     )
     flag.Parse()
     serve(*httpAddr)
}

func serve(addr string) {
    http.ListenAndServe(addr, nil)
}
```
