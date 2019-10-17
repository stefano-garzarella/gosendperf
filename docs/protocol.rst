=================
sendPerf protocol
=================
:Author: Stefano Garzarella <stefano.garzarella@gmail.com>
:Date: $Date: 2019-10-17 22:22:00 +0200 $
:status: This is a "work in progress"
:Revision: $Revision: 0001 $
:Description: Protocol definition for the sendPerf application

:abstract:

    This document describes the protocol used in the sendPerf application.

.. contents:: Table of Contents
.. section-numbering::

Introduction
============
The *test* is started when the client_ sends a request to the server_.
The request_ consists of a JSON containing all the information for the test.
The server_ must send a JSON response_, after that the transmitter can start the
test.

Actors
------

.. _server:

**server**
  Listen on a well-defined port (can be changed). It can handle multiple
  clients (default 1)


.. _client:

**client**
  Connect to the server, sending the test parameters_ with the json.

Request
=======

Parameters
----------

* type

  * request

* test

  * throughput
  * latency

* direction

  * c2s: client -> server [default]
  * s2c
  * bidirectional


JSON
----

.. code-block:: JSON

  {
      "type": "request",
      "test": "value",
      "direction": "value"
  }

Response
========

Fields
------

* type

  * response

* status

  * ack
  * nack

JSON
----

.. code-block:: JSON

  {
      "type": "response",
      "status": "value"
  }
