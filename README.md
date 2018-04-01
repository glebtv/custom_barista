## Custom Barista config

# NOT WORKING YET / DO NOT USE

This is my custom config for [barista](https://github.com/soumya92/barista) i3bar replacement.

## Installation

With a working GO env:

    go get github.com/glebtv/custom_barista/

# Usage

in ~/.i3/config:

```
bar {
  position top
  status_command exec $GOPATH/bin/custom_barista
  font pango:PragmataPro Mono 11
}
```
