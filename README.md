# micro-bastion

Micro-bastion is a minimal HTTP forwarding service originally built for accessing microservices from outside their cluster for development purposes.

## Installation

Can be built with standard Go installation:

```bash
go build
```

Or use the Dockerfile
```bash
docker build -t micro-bastion .
docker run -p 8888:8888 -it micro-bastion # starting the server
```

## Usage

```
./micro-bastion [--port PORT]
```

Then you can make requests to it on the port it is listening on (default 8888) as follows:

```
http://localhost:8888/[hostname]/[port]/[path]
```

and micro-bastion will forward the request to the specified host and port (note that the port is mandatory). Micro-bastion currently only supports http.

Example:
```
http://localhost:8888/example.org/80/
http://localhost:8888/xn--kdaa.eu/80/vant4.png
```

(it will also serve back any redirects as-is, so if opened in a browser, the second example will have you redirected to the actual website with https)


If put into e.g. AWS ECS, it can resolve service discovery domains, like
```
http://localhost:8888/myservice.local/3000/myendpoint
```

## Security

Running this opened to the internet would make essentially an open proxy. It is strongly recommended to limit access to it by firewall. Micro-bastion does not currently support any access control methods on its own.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[ISC](https://choosealicense.com/licenses/isc/), see [LICENSE](LICENSE)
