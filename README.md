# Recipes Website

## Introduction

I have a background in Python, but I'm trying to learn Go and Javascript (via 
TypeScript). To that end, I am making this little "website", which has a Go
backend API, a database, and the JS frontend using Vue.

I've picked these because someone considerably smarter than me suggested I 
should!

## Go API

This is built using the chi router. I've not taken necessarily the easiest
route here either, since I'm not using straightforward JSON (un)marshalling,
but chi's `render` package. I decided to take this route for a couple of
reasons:

1. Sometimes it's just better to take the more complicated route and learn the
hard way. I've found this has served as a good introduction to interfaces.
2. I can easily extend the structs that handle the request and response
formatting for structure (de)composition and validation.

## Infrastructure

For the sake of this project, I'll be handling this all via Docker Compose.
There will be three containers it needs to manage, one for the Go API, one for
the remote database, and one for the Vite dev server for the frontend.

The database will be a Postgresql database, where the `compose.yaml` file is
configured for persistant storage.

## TypeScript Frontend

Not started yet!

## Resources

I took most of my inspiration for the Go API from the following:

* [go-chi rest example](http://github.com/go-chi/chi/blob/master/_examples/rest/main.go)
* [This Earthly tutorial](https://earthly.dev/blog/golang-chi/)
