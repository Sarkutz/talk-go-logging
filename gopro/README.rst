
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

  $ ./gopro
  req = &{Method:GET URL:/ Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[User-Agent:[curl/7.47.0] Accept:[*/*] X-Purl:[http://www.google.com/]] Body:0x927bb0 ContentLength:0 TransferEncoding:[] Close:false Host:localhost:8080 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:43050 RequestURI:/ TLS:<nil> Cancel:<nil>}
  req = &{Method:GET URL:http://www.google.com/ Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[User-Agent:[curl/7.47.0] Accept:[*/*]] Body:0x927bb0 ContentLength:0 TransferEncoding:[] Close:false Host:www.google.com Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:127.0.0.1:43050 RequestURI: TLS:<nil> Cancel:<nil>}
  resp = &{Status:200 OK StatusCode:200 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Content-Type:[text/html; charset=ISO-8859-1] Set-Cookie:[1P_JAR=2018-02-01-18; expires=Sat, 03-Mar-2018 18:19:59 GMT; path=/; domain=.google.co.in NID=122=Rij52bFQqDnJVGLGal5g5hxxofHgBqJo4drbPdDiO7Vs6jD64PWCQIEhuMYBjlSlw4BkRn31JugDTIW9SwpRYhwzvzB3Q_kn2PWuJDYtFKBj05SumPz6qBE75ahv_-Jr; expires=Fri, 03-Aug-2018 18:19:59 GMT; path=/; domain=.google.co.in; HttpOnly] Server:[gws] X-Xss-Protection:[1; mode=block] X-Frame-Options:[SAMEORIGIN] Date:[Thu, 01 Feb 2018 18:19:59 GMT] Expires:[-1] Cache-Control:[private, max-age=0] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."]] Body:0xc820012c40 ContentLength:-1 TransferEncoding:[] Close:false Trailer:map[] Request:0xc8200ca2a0 TLS:<nil>}

