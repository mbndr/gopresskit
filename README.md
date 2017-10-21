# Gopresskit

[![Build Status](https://travis-ci.org/mbndr/logo.svg?branch=master)](https://travis-ci.org/mbndr/logo)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/gopresskit)](https://goreportcard.com/report/github.com/mbndr/gopresskit)
[![cover.run go](https://cover.run/go/github.com/mbndr/gopresskit.svg)](https://cover.run/go/github.com/mbndr/gopresskit)

 > Go implementation of [presskit()](https://github.com/ramiismail/dopresskit).


## Generation
To generate your [set up](#installation) presskit you can generate the output files with
```bash
gopresskit -input presskit-data/ -output press/
```
Run `gopresskit -help` for usage information.

## Why
I like the idea of the original presskit() but there are a view things I don't like.

##### Deployment
If you want to deploy presskit() you need a PHP server which you have to set up and which is always a potential secure gap. Although it's no web application where are often changes or need user interaction therefore static files would do their job and be a lot faster.

##### Simplicity
I'm a friend of single binaries and very simple configuration for the user. So I like the idea of a simple folder structure for the data, a single config file for page meta data and a single executable for easy generation of static html files.

## Compatibility
I try to be most compatible with the original presskit() but there are a few concepts I don't like so there are a few incompatibilities.

For easy conversion there might be a future command to migrate presskit() data to the structure of this project.

## Setup
If you want to generate your presskit you have to create following directory structure in a folder of your choice:
```
games/
    my-game-name/
        images/
            _header.png
            _icon.png
            _logo.png
            menu.png
            ingame.png
            ...
        videos/
            tailer.mp4
        game.xml
images/
    _header.png
    _icon.png
    _logo.png
    office_1.jpg
    ...
videos/
    short-movie.mp4
company.xml
favicon.ico
```
#### Optional
All `images/` and `videos/` folders and the `favicon.ico` file are optional and only have to be created if needed.

#### Images
The three images with an underscore prefix are special images (header, icon, logo) that are displayed in their sections. All other images in an `image/` folder are displayed in the ***Images*** section of the presskit.

#### Videos
There are two ways of adding videos to your presskit, [using externally hosted ones](#video) or adding your self hosted `mp4` files (other formats can be implemented, feel free to open an [issue](https://github.com/mbndr/gopresskit/issues)). The videos in this folder are like the images automatically added to your presskit's ***Videos*** section.


#### xml files
The `*.xml` files hold the data for your presskit on company and game level. See [configuration](CONFIG.md) for explanation.

## Testing
Currently the project has very low coverage. I want to write more unit tests to increase this.  
This process takes some time because of some rewrites of the code base for easier testing.
