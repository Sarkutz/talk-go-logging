
##################
 Go Proxy (GoPro)
##################

**************
 Introduction
**************


.. figure:: ../doc/gopro-overview.svg    
   :alt: gopro-overview

   GoPro Overview


*******
 Usage
*******

Start server::

  go build
  ./gopro

Sample Request
==============

Request::

  $ curl --silent --include --header 'X-purl: http://www.google.com/' 'http://localhost:8080/' --output /dev/null

Proxy Logs::

  $ go run gopro.go 
  gopro:: 2018/02/02 00:16:51 gopro.go:27: req = &{Method:GET URL:/ Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[User-Agent:[curl/7.47.0] Accept:[*/*] X-Purl:[http://www.google.com/]] Body:0x928bb0 ContentLength:0 TransferEncoding:[] Close:false Host:localhost:8080 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:43220 RequestURI:/ TLS:<nil> Cancel:<nil>}
  gopro:: 2018/02/02 00:16:51 gopro.go:40: req = &{Method:GET URL:http://www.google.com/ Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[User-Agent:[curl/7.47.0] Accept:[*/*]] Body:0x928bb0 ContentLength:0 TransferEncoding:[] Close:false Host:www.google.com Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:43220 RequestURI: TLS:<nil> Cancel:<nil>}
  gopro:: 2018/02/02 00:16:51 gopro.go:49: resp = &{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Content-Type:[text/html; charset=ISO-8859-1] X-Frame-Options:[SAMEORIGIN] Expires:[-1] Cache-Control:[private, max-age=0] Server:[gws] X-Xss-Protection:[1; mode=block] Set-Cookie:[NID=122=W6iXylzubVcUoRJGJllzcXMGTfEpEEnhqRptmIGEVKkIYelA3P-UJslz4k_97nClPUgTvxdG26qeB568OkZpnPOUBqToqvYmJHlnI52GHMB2zg6hF1_Do4xpOT3FVRws; expires=Fri, 03-Aug-2018 18:46:51 GMT; path=/; domain=.google.co.in; HttpOnly] Date:[Thu, 01 Feb 2018 18:46:51 GMT] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."]] Body:0xc820012d40 ContentLength:-1 TransferEncoding:[] Close:false Trailer:map[] Request:0xc8200ca2a0 TLS:<nil>}

