## This is a simple Go project that gets a string and generates QRcode


### Dependencies
 Don't forget to import dependencies:
 ```
 Go get "github.com/skip2/go-qrcode"
 ```
### use these commands to create Linux or Windows executable files

  - Windows

    ```
    GOOS=windows  GOARCH=amd64 go build -o qrcodeGenera
    tor.exe
    ```
  - Linux
  
    ```
    GOOS=linux GOARCH=amd64 go build -o qrcodeGenerator
    ```
