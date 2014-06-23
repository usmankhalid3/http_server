 package main

import "net"
import "strings"
import "os"
import "io/ioutil"
import "strconv"

//processes incoming requests
func processRequest(c net.Conn) {

    for {
        //read the headers
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }
        dataStr := string(buf[:nr])

        //tokenize the headers
        headers := strings.Fields(dataStr)

        //second piece is the resource name
        var resource = headers[1]
        println("Requested resource: " + resource)

        //all resources are served from the www directory
        if len(resource) > 0 {
            resource = "www" + resource
        }

        //check if the file exists
        file, fileExists := fileExists(resource)

        var CRLF string = "\r\n"
        var statusCode string
        var contentType string
        var contentLength string
        var entityBody []byte

        // if file exists, defer the closing to the end
        if fileExists {
            defer file.Close()
        } else {
            //if file doesn't exist, prepare redirect to default image
            resource = "www/404.jpg"
        }

        //prepare response
        statusCode = "HTTP/1.0 200 OK" + CRLF
        contentType = "Content-Type: " + resContentType(resource) + CRLF
        entityBody, _ = ioutil.ReadFile(resource)
        contentLength = "Content-Length: " + strconv.Itoa(len(entityBody)) + CRLF

        //write back the response to client
        c.Write([]byte(statusCode))
        c.Write([]byte(contentType))
        c.Write([]byte(contentLength))
        c.Write([]byte(CRLF))
        c.Write(entityBody)

        //close the incoming connection
        c.Close()
    }
}

//returns the content-type based on the file extension
func resContentType(resource string) string {
    if strings.HasSuffix(resource, "html") || strings.HasSuffix(resource, "htm") {
        return "text/html"
    }
    if strings.HasSuffix(resource, "jpeg") || strings.HasSuffix(resource, "jpg") {
        return "image/jpeg"
    }
    if strings.HasSuffix(resource, "png") {
        return "image/png"
    }
    if strings.HasSuffix(resource, "css") {
     return "text/css"
    }
    return "text/html"
}

//checks if a file exists
func fileExists(resource string) (*os.File, bool) {
    file, err := os.Open(resource)
    return file, (err == nil)
}

//server boots up here
func main() {
    l, err := net.Listen("tcp", ":8082")
    if err != nil {
        println("listen error")//, err.String())
        return
    }
    println("Listening on port: 8082...")
    for {
        fd, err := l.Accept()
        println("Incoming connection...")
        if err != nil {
            println("accept error")//, err.String())
            return
        }
        go processRequest(fd)
    }
  }