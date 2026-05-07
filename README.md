# mustag

Minimalistic and fast CLI utility for viewing and editing ID3v2 tags in audio files.

Written in Go using Cobra and id3v2 library.

![Go Version](https://img.shields.io/github/go-mod/go-version/d1manpro/mustag)
![License](https://img.shields.io/github/license/d1manpro/mustag)
![Release](https://img.shields.io/github/v/release/d1manpro/mustag)

## Features

* View ID3v2 metadata
* Edit standard tags:
  * title
  * artist
  * album
  * album artist
  * genre
  * year
  * track number
  * disk number
* Add lyrics from file
* Embed cover art
* Add custom ID3 frames
* Show all raw ID3 frames
* Script-friendly output

---

# Installation

## Download Binary

Prebuilt binaries are available on the [github releases page](https://github.com/d1manpro/mustag/releases).

### Linux amd64

```bash
curl -LO https://github.com/d1manpro/mustag/releases/latest/download/mustag-linux-amd64.tar.gz

tar -xzf mustag-linux-amd64.tar.gz

chmod +x mustag

sudo install -m755 mustag /usr/local/bin/mustag
```

Verify installation:

```bash
mustag --version
```

---

## Build From Source

Requirements:
- Go 1.26+

```bash
git clone https://github.com/d1manpro/mustag.git
cd mustag

go build -o mustag .
```

Optional install:

```bash
sudo install -m755 mustag /usr/local/bin/mustag
```
---

# Usage

```bash
mustag [command]
```

Available commands:

| Command | Description     |
| ------- | --------------- |
| `get`   | Show metadata   |
| `set`   | Update metadata |

---

## Get Tags

Show all basic metadata:

```bash
mustag get song.mp3
```

Example output:

```text
=== song.mp3 ===
title:  Numb
artist: Linkin Park
album:  Meteora
year:   2003
track:  13
cover:  1 image(s): image (image/jpeg, 89452 bytes)
lyrics: 1 item(s), 43 lines
```

---

### Get Specific Fields

Single field:

```bash
mustag get song.mp3 title
```

Output:

```text
Numb
```

Multiple fields:

```bash
mustag get song.mp3 title artist album
```

Output:

```text
title:  Numb
artist: Linkin Park
album:  Meteora
```

---

### Get All Fields

```bash
mustag get song.mp3 --full
```

---

## Set Tags

Set basic tags:

```bash
mustag set song.mp3 \
  -t "Numb" \
  -a "Linkin Park" \
  -A "Meteora" \
  -y 2003
```


### Set Album Artist

```bash
mustag set song.mp3 --album-artist "Various Artists"
```


### Set Track and Disk Numbers

```bash
mustag set song.mp3 -n 13 -d 1
```


### Add Lyrics

```bash
mustag set song.mp3 --lyrics lyrics.txt
```


### Add Cover Art

```bash
mustag set song.mp3 --cover cover.jpg
```


### Add Custom Frames

```bash
mustag set song.mp3 --custom TXXX:MyValue
```

Multiple custom frames:

```bash
mustag set song.mp3 \
  --custom TXXX:Value1 \
  --custom COMM:Hello
```

---

## Flags

### `get`

| Flag         | Description             |
| ------------ | ----------------------- |
| `-f, --full` | Show all raw ID3 frames |

### `set`

| Flag             | Description           |
| ---------------- | --------------------- |
| `-t, --title`    | Set title             |
| `-a, --artist`   | Set artist            |
| `-A, --album`    | Set album             |
| `--album-artist` | Set album artist      |
| `-g, --genre`    | Set genre             |
| `-y, --year`     | Set year              |
| `-n, --number`   | Set track number      |
| `-d, --disk`     | Set disk number       |
| `-l, --lyrics`   | Load lyrics from file |
| `-c, --cover`    | Load cover image      |
| `--custom`       | Add custom ID3 frame  |
