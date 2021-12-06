**Example of Microservices using go-kit**

Go kit is a collection of Go (golang) packages (libraries) that help you build robust, reliable, maintainable microservices

**Architecture and design**

Go kit services are laid out in three layers:

1. Transport layer
2. Endpoint layer
3. Service layer

Requests enter the service at layer 1, flow down to layer 3, and responses take the reverse course.

1. **Transport layer**
The transport domain is bound to concrete transports like HTTP or gRPC. You can support a legacy HTTP API and a newer RPC service, all in a single microservice.

2. **Endpoint layer**
An endpoint is like an action/handler on a controller; it’s where safety and antifragile logic lives. If you implement two transports (HTTP and gRPC), you might have two methods of sending requests to the same endpoint.

3. **Service layer**
Services are where all of the business logic is implemented. A service usually glues together multiple endpoints. In Go kit, services are typically modeled as interfaces, and implementations of those interfaces contain the business logic.


**Example Project**

1 - Initialize the project using ```go mod```. Execute the command inside the project empty folder

```go
go mod init github.com/githubaccount/gokit-example-string-service
```

2 - The project will have the following folder structure

```
cmd  
internal      
    endpoint
    entity
    handler
    service
```

The ```cmd``` folder contains the main application entry point files for the project

The ```internal``` folder will store packages that are meant to be scoped to this project. The internal directory is a Go convention and doesn’t allow the Go compiler to accidentally use its packages in an external project.

***Service*** 

We start defining the service. Here we implement our business logic. In this example, a function to convert string to uppercase and another to count the length. Normally in go kit, is a good practice to define an interface for our service

```
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
```

***Enpoint***

In Go kit, the primary messaging pattern is RPC. So, every method in our interface will be modeled as a remote procedure call. For each method, we define request and response structs, capturing all of the input and output parameters respectively.

In the entity folder, we create the request and response structs
```
package entity

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}
```

And then we define an endpoint for every method in the service

***Handler***

In the Handler, we define the transport layer. In this case, we use an httptransport 

***Run the example***  
Execute the application

```
go run  .\cmd\main.go
```

Request example

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase

curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count


***References***
 Most of the texts are from go kit documentation 
<https://gokit.io/faq/>
<https://gokit.io/examples/stringsvc.html>
