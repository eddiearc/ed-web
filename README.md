# ed-web
A simple http framework in go.

## feature

### http
Use golang's `http` package's method `http.ListenAndServe` to listening server port.

### router
Use Trie to reduce redundant k-v(path-handler) design.

### context
Design request context to hold http writer and many request message, which will exposure to users. 

### group and middleware
Design requests are grouped and users can customize their middleware and apply it.

### template
Use golang's `html/template` package to render html page. 

### panic dispose
Design a middleware which use golang `recover` to avoid panic throw to user continue. 
