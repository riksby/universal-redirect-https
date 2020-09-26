# Universal HTTP to HTTPS redirect microservice

Listens on port 80 and redirects any requests to their equivalent https url.
This can be used as an alternative standalone service to Traefik's builtin redirect functionality.
Removes port if present:

| request                          | redirect                          |
| -------------------------------- | --------------------------------- |
| http://mydomain.com/path?foo=bar | https://mydomain.com/path?foo=bar |
| http://foo.mydomain.com:1234/bar | https://foo.mydomain.com/bar      |

## Image

[ghcr.io/riksby/universal-redirect-https](https://github.com/orgs/riksby/packages/container/package/universal-redirect-https)

## Usage

```sh
docker run --rm -it -p 8080:80 ghcr.io/riksby/universal-redirect-https
```
