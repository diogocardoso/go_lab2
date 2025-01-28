# GO LAB 2
Projeto do Laboratório "Tracing distribuído e span" do treinamento GoExpert(FullCycle).

## O desafio
Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin) juntamente com a cidade. Esse sistema deverá implementar OTEL(Open Telemetry) e Zipkin.

## Como rodar o projeto: manual

``` shell
## 1. Clone o repo

## 2. Configure o .env

## 3. Coloque sua api-key como valor na variável OPEN_WEATHERMAP_API_KEY no .env

## 4. Baixe compose, se estiver up
docker-compose down

## 5. Suba o compose 
docker-compose up -d 
    ou 
docker-compose up --build

## 6. Faça uma chamada no input-api para gerar os traces
curl curl -s -X POST "http://localhost:8080/cep" -H "Content-Type: application/json" -d '{"cep": "13330300"}'

## 7. Veja os traces via Zipkin: 
http://localhost:9411