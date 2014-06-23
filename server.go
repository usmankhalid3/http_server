 package main

  import "net"
import "strings"
import "os"
import "io/ioutil"
import "strconv"

type T1 struct {
    f1 []byte
    f2 int32
}
  func processRequest(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }
        dataStr := string(buf[:nr])
        headers := strings.Fields(dataStr)
        var resource = headers[1]
        println("Requested resource: " + resource)

        if len(resource) > 0 {
            resource = "www" + resource
        }

        file, fileExists := fileExists(resource)

        var CRLF string = "\r\n"
        var statusCode string
        var contentType string
        var contentLength string
        var entityBody []byte

        if fileExists {
            defer file.Close()
            statusCode = "HTTP/1.0 200 OK" + CRLF
            contentType = "Content-Type: " + resContentType(resource) + CRLF
            entityBody, _ = ioutil.ReadFile(resource)
            contentLength = "Content-Length: " + strconv.Itoa(len(entityBody)) + CRLF
        } else {
            statusCode = "HTTP/1.0 404 Not Found" + CRLF
            contentType = "Content-type: " + "text/html" + CRLF
        }
        c.Write([]byte(statusCode))
        c.Write([]byte(contentType))
        c.Write([]byte(contentLength))
        c.Write([]byte(CRLF))
        if len(entityBody) > 0 {
            c.Write(entityBody)
        }
        c.Close()
    }
  }

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

  func fileExists(resource string) (*os.File, bool) {
    file, err := os.Open(resource)
    return file, (err == nil)
  }

  func response(c net.Conn) {
    println("This is the response")
  }

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