Distributed Node JS Applications
* Build two applications
    - _recipe-api_ : internal api that isn't accessed from outside world. It is only accessed by other internal applications. So you can use any protocol to communicate it with becuase you own the service and the client that need it.
    - _web-api_    : this is an api that is accessed by third partied over the internet. It exposes HTTP server so that web browsers can easily communicate with it.
* Service Relationships :
    - _web-api_    : This is the downstream client.
    - _recipt-api_ : This is the upstream server.
* Both of these applications acan be refereed to as server becuase they are both actively listening for incoming network requests.
* However when describing the specific realtionship between the two API's :
    - the *web-api* can be refreed to as the **client/consumer**,
    - the *recipe-api* can be referred to as **producer/server**
* When referring to the realtionship between web-browser and web-api :
    - the web browser is called the client/consumer,
    - the web-api server is called the server/producer.
