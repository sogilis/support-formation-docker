# support-formation-docker

Small golang app to host image on the disk.

* Default port is `7777`

* Requires at least go `1.16`
* All dependencies are vendored (Less risks of breaking over time)
* The templates files (used to generate the page) are embedded in the binary to makes it portable

# How to use:

* To launch the app: `make run` or `go run .`
* To build : `go build . -o app`
* CLI is documented, you can use `<binary> -h`


## CLI
``` 
Usage of <binary>:
  -port uint
    	HTTP server port used to listen (default 7777)
  -storage-path string
    	Image storage folder (default "/tmp/adc60fb2-31fa-4cc3-ae6e-1cf8727143aa")
```