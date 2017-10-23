# Portoflio Website - [live](https://www.marwan.io)

## **Running in development**

### *backend*

`GO_MODE=development go run *.go`

or `GO_MODE=development gowatch` if you want autoreload: https://github.com/marwan-at-work/gowatch

###  *frontend*

**One time conifg**

`cd ./frontend && npm i && webpack` 

If you don't have npm, install node.js. If you don't have webpack, `npm i -g webpack`. 

**Build & watch**

`cd ./frontend && gopherjs build github.com/marwan-at-work/marwanio/frontend -o ../public/frontend.js -w`