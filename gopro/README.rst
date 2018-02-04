
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

Install dependencies::

  go get github.com/sirupsen/logrus

Start server::

  go build
  ./gopro

Sample Request
==============

HTTP
""""

Request::

  $ curl --silent --include --header 'X-purl: http://www.google.com/' 'http://localhost:8080/' --output /dev/null

Proxy Logs::

  $ ./gopro
  {"httpver":"HTTP/1.1","level":"debug","method":"GET","msg":"Received request from User Agent","peer":"ua","pkg":"main","reqid":1517716295,"scheme":"http","time":"2018-02-04T09:21:35+05:30","url":"http://www.google.com/"}
  {"httpver":"HTTP/1.1","level":"debug","method":"GET","msg":"Sending request to Origin Server","peer":"origin","pkg":"main","reqid":1517716295,"scheme":"http","time":"2018-02-04T09:21:35+05:30","url":"http://www.google.com/"}
  {"httpver":"HTTP/1.1","level":"debug","msg":"Received response from Origin Server","peer":"origin","pkg":"main","reqid":1517716295,"scheme":"http","statuscode":200,"time":"2018-02-04T09:21:35+05:30","url":"http://www.google.co.in/?gfe_rd=cr\u0026dcr=0\u0026ei=R4N2WufbHJLm8wfrtaDgBQ"}

HTTPS
"""""

Request::

  $ curl --silent --include --header 'X-purl: https://www.google.com/' 'http://localhost:8080/' --output /dev/null

Proxy Logs::

  $ ./gopro
  {"httpver":"HTTP/1.1","level":"debug","method":"GET","msg":"Received request from User Agent","peer":"ua","pkg":"main","reqid":1517716333,"scheme":"https","time":"2018-02-04T09:22:13+05:30","url":"https://www.google.com/"}
  {"httpver":"HTTP/1.1","level":"debug","method":"GET","msg":"Sending request to Origin Server","peer":"origin","pkg":"main","reqid":1517716333,"scheme":"https","time":"2018-02-04T09:22:13+05:30","url":"https://www.google.com/"}
  {"httpver":"HTTP/2.0","level":"debug","msg":"Received response from Origin Server","peer":"origin","pkg":"main","reqid":1517716333,"scheme":"https","statuscode":200,"time":"2018-02-04T09:22:14+05:30","url":"https://www.google.co.in/?gfe_rd=cr\u0026dcr=0\u0026ei=boN2WqCqEovm8weS4ZqgDw"}

