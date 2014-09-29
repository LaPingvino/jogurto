jogurto
=======

A pacman/yaourt inspired tool for Debian package management.

This is a go-program. You can build it with `go build`, or if you don't like to install go on your system,
use [![Gobuild Download](https://img.shields.io/badge/gobuild-download-green.svg?style=flat)](http://gobuild.io/github.com/lapingvino/jogurto).

    Usage: jogurto option [packages]
    
    Where options is one of:
        --help, -h: show this help
        -S: install packages
        -Sy: update package database [and install]
        -Syu: update && upgrade && install
        -Syyu: update && dist-upgrade && install
        -Ss: search in installable packages
        -R: remove packages
        -Rn: purge packages
        -Rs: remove packages, then autoremove dependencies
        -Q: show all installed packages

I tried to make the code really easy to add to, and your contributions are very welcome!
