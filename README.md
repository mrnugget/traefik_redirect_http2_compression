## Setup
```
docker-compose build
docker-compose up
```

## The Bug

```
$ /usr/local/opt/curl/bin/curl -v -k --http2 -H 'Accept-Encoding: gzip' -H 'Host: redirectwebapp.docker.localhost' https://localhost/redirect
*   Trying ::1...
* TCP_NODELAY set
* Connection failed
* connect to ::1 port 443 failed: Connection refused
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH
* successfully set certificate verify locations:
*   CAfile: /usr/local/etc/openssl/cert.pem
  CApath: /usr/local/etc/openssl/certs
* TLSv1.2 (OUT), TLS header, Certificate Status (22):
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Client hello (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS change cipher, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd
*  start date: Nov  1 13:52:21 2017 GMT
*  expire date: Nov  1 13:52:21 2018 GMT
*  issuer: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x7fbe92808400)
> GET /redirect HTTP/2
> Host: redirectwebapp.docker.localhost
> User-Agent: curl/7.56.1
> Accept: */*
> Accept-Encoding: gzip
>
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
< HTTP/2 200
< content-type: text/html; charset=utf-8
< date: Wed, 01 Nov 2017 14:46:45 GMT
< location: https://redirectwebapp.docker.localhost/end
< vary: Accept-Encoding
< content-length: 109
<
* Connection #0 to host localhost left intact
<html><body>You are being <a href="https://redirectwebapp.docker.localhost/end">redirected</a>.</body></html>% 
```

