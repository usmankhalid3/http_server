package main

  import (
    "net"
    "time"
    "io"
  )

  func reader(r io.Reader) {
    buf := make([]byte, 1024)
    for {
      n, err := r.Read(buf[:])
      if err != nil {
        return
      }
      println("Client got:", string(buf[0:n]))
    }
  }

  func main() {
    c,err := net.Dial("tcp", "localhost:8081")
    if err != nil {
        //panic(err.String())
    }
    defer c.Close()

    go reader(c)
    for {
        _,err := c.Write([]byte("hi"))
        if err != nil {
            //println(err.String())
            break
        }
        time.Sleep(1e9)
    }
  }