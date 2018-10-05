# GO WebApp

### Why Go?
    - Focus on simplicity and clarity
    - Full-features HTTP processing and templating libraries
    - Easy to deploy

#### Handling Basic Requests
    - Need to use Connection Handler for concurrent requests. Otherwise can just handle 1 request on the port it's listening on.
        - Connection Handler creates a goroutine which will listen on another port to handle communication back and forth to honor that request.

    - Built in Handlers
        - NotFoundHandler
            - Returns a 404 to the requester.
        - RedirectHandler(url string, code int)
        - StripPrefix(prefix string, h handler)
        - TimeoutHandler(h handler, dt time.Duration, msg string)
            - dt -> the amount of time the handler, h, is allowed to process
        - FileServer(root FileSystem)
            - FileSystem is an interface that defines one method, Open
            - Can use type Dir to delegate to machine OS

### Pipelines
{{command1, command2, command3}}

Commands
    - Literal String or Number
    - Function Name
    - Data Field (i.e. {{.Title}})
    - Method
    - Argument

Output
    - Must have one or two values.
        - Second must be an error type.

{{ command1 command2 | command3 }}
    - The pipe passes the reuslt of previous command as last argument of next command.

export GOPATH=~/projects/go-webapp/
go install src/example