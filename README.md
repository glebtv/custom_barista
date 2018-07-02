## Custom Barista config

This is my custom config for [barista](https://github.com/soumya92/barista) i3status replacement.

Includes custom keyboard layout module

## Installation

With a working GO env:

    go get github.com/glebtv/custom_barista/

Add fonts from fonts dir to ~/.fonts

# Usage

in ~/.i3/config:

```
bar {
  position top
  status_command exec $GOPATH/bin/custom_barista
  font pango:PragmataPro Mono 11
}
```
