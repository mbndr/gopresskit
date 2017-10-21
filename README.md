# Gopresskit

[![Build Status](https://travis-ci.org/mbndr/logo.svg?branch=master)](https://travis-ci.org/mbndr/logo)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/gopresskit)](https://goreportcard.com/report/github.com/mbndr/gopresskit)
[![cover.run go](https://cover.run/go/github.com/mbndr/gopresskit.svg)](https://cover.run/go/github.com/mbndr/gopresskit)

 > Go implementation of [presskit()](https://github.com/ramiismail/dopresskit).

## Why
I like the idea of the original presskit() but there are a view things I don't like.

##### Deployment
If you want to deploy presskit() you need a PHP server which you have to set up and which is always a potential secure gap. Although it's no web application where are often changes or need user interaction therefore static files would do their job and be a lot faster.

##### Simplicity
I'm a friend of single binaries and very simple configuration for the user. So I like the idea of a simple folder structure for the data, a single config file for page meta data and a single executable for easy generation of static html files.

## Compatibility
I try to be most compatible with the original presskit() but there are a few concepts I don't like so there are a few incompatibilities.

For easy conversion there might be a future command to migrate presskit() data to the structure of this project.

## Example
Currently I'm using most images and videos from [subset games](http://subsetgames.com/presskit/sheet.php?p=into_the_breach) for example presentation. All copyright goes to them.

## Testing
Currently the project has very low coverage. I want to write more unit tests to increase this.  
This process takes some time because of some rewrites of the code base for easier testing.
