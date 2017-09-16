
# go-starter

Starter project for golang to help you get started with web development.

# Install dependencies & Build

Install dev dependencies && run locally:
    
    go get github.com/codegangsta/gin
    go get github.com/GeertJohan/go.rice/rice
    go get

    gin

Build assets to binary and deploy 'go-starter' binary

    rice embed-go
    go build
