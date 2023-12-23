# Server

* ~~Run on port `:8080`~~
* ~~Fetch USD-BRL data from [here](https://economia.awesomeapi.com.br/json/last/USD-BR)~~ 
* ~~Timeouts~~
  * ~~DB -> `10ms`~~
  * ~~Currency API -> `200ms`~~
* ~~Should return only `bid` as *json*~~


# Client

* Should request `/cotacao` from `server`
* Should receive only `bid` from `server`
* Timeout
  * Server -> `300ms` 
* Save response into `cotacao.txt`
  ``Dólar: {value}`` 