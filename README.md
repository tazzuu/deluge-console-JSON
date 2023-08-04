# dcJSON

Deluge console JSON output formatter.

A JSON serializer for command line output from `deluge-console`, so you can more easily parse the status of your Deluge torrents from the cli. Designed for use with [`jq`](https://jqlang.github.io/jq/) to allow for easier filtering and formatting.

Built with Go (Golang); https://go.dev/doc/install

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

Quick notes for setting up and handling torrent downloads from the command line with Deluge.

Download all .torrent files and load them

```
$ wget -nd -r -P . -A torrent "https://old-releases.ubuntu.com/releases/

$ wget -nd -r -P . -A torrent "https://torrent.ubuntu.com/xubuntu/releases/"

$ wget -nd -r -P . -A torrent "https://cdimage.ubuntu.com/lubuntu/releases/"

$ for i in *.torrent; do deluge-console "add $i"; done

```

save text output

```
deluge-console "info -d -v" > deluge-console-v-d.$(date '+%s').txt
```

move a torrent

```
$ deluge-console "move ubuntu-23.04-desktop-amd64.iso /home/username/Downloads"
```
