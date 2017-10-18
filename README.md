Portoflio Website - [live](https://www.marwan.io)

**Running in development**

*backend*

`GO_MODE=development go run *.go`

or `GO_MODE=development gowatch` if you want autoreload: https://github.com/marwan-at-work/gowatch

*frontend*

`cd ./frontend && gopherjs build github.com/marwan-at-work/marwanio/frontend -o ../public/frontend.js -w`