## Custom Barista config

This was my custom config for [barista](https://github.com/soumya92/barista) i3status replacement.

**Note:** The barista project is now archived. This config has been migrated to [i3status-rust](https://github.com/greshake/i3status-rust).

## Migration to i3status-rust

Since the barista library is archived, the custom modules have been replaced with i3status-rust configuration files.

### Installation

Install i3status-rust via your distribution's package manager:

```bash
# Arch Linux
pacman -S i3status-rust

# Debian/Ubuntu
apt install i3status-rust

# Fedora
dnf install i3status-rust

# Other distributions: build from source via cargo
cargo install i3status-rust
```

### Font Requirements

This configuration uses [Nerd Fonts](https://www.nerdfonts.com/) for icons. You must install a Nerd Font for icons to display correctly.

**Arch Linux:**
```bash
pacman -S ttf-inconsolata-nerd
```

See the [Arch Nerd Fonts package group](https://archlinux.org/groups/any/nerd-fonts/) for all available options.

**Other distributions:** Download from [Nerd Fonts GitHub](https://github.com/ryanoasis/nerd-fonts) or [nerdfonts.com](https://www.nerdfonts.com/)

### Configuration Files

Two configuration variants are provided in the `docs/` directory:

| File | Description |
|------|-------------|
| `config-minimal.toml` | Minimal working config with essential modules |
| `config-modern.toml` | Full-featured config with all common modules enabled (auto-detects hardware)

### Usage

1. Copy the desired config to your i3status-rust config directory:

```bash
cp docs/config-modern.toml ~/.config/i3status-rust/config.toml
```

2. Configure your `~/.i3/config`:

```
bar {
  position top
  status_command i3status-rs
  font pango:Inconsolata Nerd Font 11
}
```

### Calendar Integration

The time block is configured to open `gsimplecal` when clicked. Install it:

```bash
# Arch Linux
pacman -S gsimplecal

# Other distributions: build from https://github.com/dmedvinsky/gsimplecal
```

### Module Mapping (barista → i3status-rust)

| barista Module | i3status-rust Block |
|----------------|---------------------|
| kbdlayout | keyboard_layout |
| batt | battery |
| dsk (diskspace) | disk_space |
| dsk (diskio) | disk_iostats |
| load | load |
| mem | memory |
| netm | net |
| music | music |
| ltime | time |
| temp | temperature |

### Caps Lock / Num Lock Indicators

**i3status-rust does not have built-in caps/num lock indicators.** This would require modifying the keyboard_layout block source code to access XKB modifier states.

**Workaround using a custom block:**

Add this to your i3status-rust config:
```toml
[[block]]
block = "custom"
command = "cat /sys/class/leds/input*::capslock/brightness 2>/dev/null | grep -q 1 && echo ' CAPS ' || echo ''"
interval = 1
```

And in your `~/.i3/config`:
```
bindsym --release Caps_Lock exec pkill -SIGRTMIN+11 i3status-rs
```

For a proper implementation with event-driven updates (like the original barista config), the i3status-rust keyboard_layout block would need to be modified to expose XKB LED states as placeholders.

### Further Customization

See the [i3status-rust documentation](https://man.archlinux.org/man/extra/i3status-rust/i3status-rs.1.en) for all available blocks and configuration options.

Icon sets can be changed in the `[icons]` section:
- `awesome6` (default, FontAwesome 6)
- `awesome5` (FontAwesome 5)
- `material` (Material Design icons)
- `emoji` (emoji icons)

Themes can be changed in the `[theme]` section:
- `solarized-dark` (default)
- `nord-dark`
- `dracula`
- `gruvbox-dark`
- And many more...

---

## Historical Installation (barista)

With a working GO env:

    go get github.com/glebtv/custom_barista/

Add fonts from fonts dir to ~/.fonts

in ~/.i3/config:

```
bar {
  position top
  status_command exec $GOPATH/bin/custom_barista
  font pango:PragmataPro Mono 11
}
```
