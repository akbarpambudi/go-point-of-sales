## Background

DDD (Domain-Driven Design) has recently become a hot topic in software development, especially in the era of microservices like today.
This is due to the need to be able to divide our business problems into services that are in accordance with their designation,
so, that we benefit from the microservice itself. However, the implementation of DDD, especially in the tactical sector, is still a challenge in the Golang programming language.

A few days ago, I read an article that I thought might be the answer for implementing DDD in GO. The article was written by a Three Dots Labs
the author of the famous watermill library. Therefore, in this repository I tried to adopt the concept of DDD from the article as well as several other sources to the point of sales case.

## Inspiration Sources
This repository is highly inspired by some sources bellow:
- https://threedots.tech/post/ddd-lite-in-go-introduction/
- https://threedots.tech/post/ddd-cqrs-clean-architecture-combined/

** note: Feel free to give a feedback about how I implement the concept from those articles. 

## How to use this project
### Setup and Preparation
1. clone this repository
2. execute `make setup` to install all required tools.
3. sync the project dependencies using `go mod vendor`
### Run Test
1. run `make unit_test` for executing all unit test.
2. run `make component_test` for executing all component test.
3. run `make test` to executing unit test and component test.
