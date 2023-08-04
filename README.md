# dcJSON

## Usage

start Deluge daemon process

```
$ deluged
```

check Deluge console

```
$ deluge-console "info -d -v"
```

pipe it to `dcJSON`

```
$ deluge-console "info -d -v" | go run main.go
```

Get path to all torrent files

```
$ deluge-console "info -d -v" | go run main.go | jq -r '.[] | .downloadFolder + "/" + .name + "\t" + .state '
```

# Notes

Download all .torrent files and load them

```
$ wget -nd -r -P . -A torrent "https://old-releases.ubuntu.com/releases/

$ wget -nd -r -P . -A torrent "https://torrent.ubuntu.com/xubuntu/releases/"

$ wget -nd -r -P . -A torrent "https://cdimage.ubuntu.com/lubuntu/releases/"

$ for i in *.torrent; do deluge-console "add $i"; done

```

move a torrent

```
$ deluge-console "move ubuntu-23.04-desktop-amd64.iso /home/username/Downloads"


```