# EpiCRM Pluggable Module Demo

[![CI](https://github.com/epixelnan/epicrm-module-demo/actions/workflows/main.yml/badge.svg)](https://github.com/epixelnan/epicrm-module-demo/actions/workflows/main.yml)

This repo intends to be an example pluggable module for EpiCRM.

NOTE: Both the API service source code and the manifesto files (manifesto.json,
docker-compose.yml, etc.) are kept in the same repo in this example. But this
is for convenience only. The frameworks expects modules as pre-built container
images, and `epicrmanage.py` requires only the manifesto files in the source
form.

## How to Load

Create a directory named `xmodules` (if it doesn't exist already) under your
epicrm project directory. Enter the directory and clone this repo, so that the
directory structure looks something like this:

```
├── config.json
...
├── modconf.d
...
└── xmodules
    └── epicrm-module-demo
        ├── apps
        │   └── helloapi
...
```
Now enable the module:

```
epicrm-bootstrap/epicrmanage.py --projdir $PROJDIR --module hello enmod
epicrm-bootstrap/epicrmanage.py --projdir $PROJDIR init-fix-missing
epicrm-bootstrap/epicrmanage.py --projdir $PROJDIR up
epicrm-bootstrap/epicrmanage.py --projdir $PROJDIR reload-gateway
```
