
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
  nohup ./gopro > out.log &

Sample Request
==============

Please see ``dat/*.in`` for sample requests and ``dat/*.out`` for proxy logs.

::

  bash data/10.in

