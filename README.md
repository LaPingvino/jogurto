jogurto
=======

A pacman/yaourt inspired tool for Debian package management.

This is a go-program. You can build it with `go build`, or if you don't like to install go on your system,
use [![Gobuild Download](https://img.shields.io/badge/gobuild-download-green.svg?style=flat)](http://gobuild.io/github.com/lapingvino/jogurto).

It's in a very early stage, but the following options work (they are passed to the corresponding apt-commands):

- -S to install
- -Sy to update, then install
- -Syu to update, then upgrade, then install
- -Syyu to update, then dist-upgrade, then install
- -Ss to search

I tried to make the code really easy to add to, and your contributions are very welcome!
