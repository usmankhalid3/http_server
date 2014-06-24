http_server
===========

A Go based very simple HTTP server

Deployment
==========
1. Download or clone the source code.
2. cd to the source code directory.
3. go run server.go
4. The server will run on port 8082.
5. Hit http://localhost:8082/start.html to see one sample of hosted package.

Steps to host files
===================
All files are served from within the "www" directory located in the source code. Please add all content there.

Notes/Known issues
==================
1. The server port is hardcoded to 8082 for now.
2. Only static content can be served at the moment.
3. Only supports html, htm, jpg, png & css files.
