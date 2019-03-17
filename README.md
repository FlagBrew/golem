# Golem
[![Go Report Card](https://goreportcard.com/badge/github.com/FlagBrew/golem)](https://goreportcard.com/report/github.com/FlagBrew/golem)
[![GoDoc](https://godoc.org/github.com/FlagBrew/golem?status.png)](https://godoc.org/github.com/FlagBrew/golem)
Golem is an API wrapper library for the [PokeApi](https://pokeapi.co/) V2 API, built in Golang


# Notice

This was built on Linux, so knowing my luck, it probably won't behave nicely on Windows.

If you do find an issue with using it on Windows or MacOS please let me know in the issues page.

# Using this library

First you'll want to do a go get

``go get github.com/flagbrew/golem``

Then you'll to import it into your application, once there the syntax is

golem.Get(Object)()

To see a list of available objects, you can find the documentation for the API [here](https://pokeapi.co/docs/v2.html/)


# Defaults

The default location for the cached data is $HomeDirectory/.golem/ you are able to change this with the SetCachePath() command


# Reporting bugs

I've done my best to test everything, but of course something probably will have slipped by.

If you find any bugs with this library, please report it on the issues page.

When reporting a bug, try to be as descriptive as possible, thank you.


# Copyright & trademarks

"Golem" is a Pokemon, Pokemon is owned by Gamefreak & Nintendo and licensed by Nintendo
All rights for the name and any reference to Pokemon is owned by them.

I obviously don't own the rights to the name or any name in relation to Pokemon, 
the only thing I own is the time I put into building this API wrapper.

# License
MIT License

Copyright (c) 2019 FlagBrew & Allen (FM1337/FMCore)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
