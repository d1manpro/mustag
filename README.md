# mustag

Minimalistic and fast CLI utility for viewing and editing ID3v2 tags in audio files.

Written in Go using Cobra and id3v2 library.

![Go Version](https://img.shields.io/github/go-mod/go-version/d1manpro/mustag)
![License](https://img.shields.io/github/license/d1manpro/mustag)
![Release](https://img.shields.io/github/v/release/d1manpro/mustag)

---

## Features

- Read ID3v2 metadata from audio files
- Edit standard ID3 tags:
  - title
  - artist
  - album
  - album artist
  - genre
  - year
  - track number
  - disk number
- Add and edit embedded lyrics
- Embed and replace cover art
- Manage custom ID3 frames (TXXX, COMM, etc.)
- View raw ID3 frames
- Remove specific or all tags from files
- CLI-friendly output for scripting and automation

---

## Installation

### Download binary

Prebuilt binaries are available on the [github releases page](https://github.com/d1manpro/mustag/releases).

#### Linux amd64 example:

```bash
curl -LO https://github.com/d1manpro/mustag/releases/latest/download/mustag-linux-amd64.tar.gz

tar -xzf mustag-linux-amd64.tar.gz
chmod +x mustag

sudo install -m755 mustag /usr/local/bin/mustag
````

Verify:

```bash
mustag --version
```

---

### Build from source

Requirements:

* Go 1.26+

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

## Usage

```bash
mustag <command>
```

### Commands

| Command          | Description     |
| ---------------- | --------------- |
| get (g)          | Show metadata   |
| set (s)          | Update metadata |
| remove (rm)      | Remove metadata |
| lyrics (l, lyr)  | Edit lyrics     |

---

## Get tags

Show metadata:

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

### Specific fields

```bash
mustag get song.mp3 title artist album
```

### Full output

```bash
mustag get song.mp3 --full
```

---

## Set tags

```bash
mustag set song.mp3 \
  -t "Numb" \
  -a "Linkin Park" \
  -A "Meteora" \
  -y 2003
```

### Options

* album artist

```bash
mustag set song.mp3 --album-artist "Various Artists"
```

* track and disk

```bash
mustag set song.mp3 -n 13 -d 1
```

* lyrics

```bash
mustag set song.mp3 --lyrics lyrics.txt
```

* cover

```bash
mustag set song.mp3 --cover cover.jpg
```

* custom frames

```bash
mustag set song.mp3 --custom TXXX:MyValue --custom COMM:Hello
```

---

## Remove tags

```bash
mustag remove song.mp3 -t -g
```

### Remove all

```bash
mustag remove song.mp3 --all
```

### Specific fields

```bash
mustag remove song.mp3 \
  --title \
  --artist \
  --album \
  --lyrics \
  --cover
```

### Custom frames

```bash
mustag remove song.mp3 --custom TXXX --custom COMM
```

---

## Lyrics editor

```bash
mustag lyrics song.mp3
```

Use custom editor:

```bash
mustag lyrics -e nvim song.mp3
```

or:

```bash
EDITOR=vim mustag lyrics song.mp3
```

### Behavior

* Opens lyrics in temp file
* Uses `$EDITOR` or fallback `vi`
* Empty content → remove tag
* Unchanged → no update
* Changed → save back

---

## Flags

### get

| Flag       | Description     |
| ---------- | --------------- |
| -f, --full | Show raw frames |

---

### set

| Flag           | Description      |
| -------------- | ---------------- |
| -t, --title    | Set title        |
| -a, --artist   | Set artist       |
| -A, --album    | Set album        |
| --album-artist | Set album artist |
| -g, --genre    | Set genre        |
| -y, --year     | Set year         |
| -n, --number   | Track number     |
| -d, --disk     | Disk number      |
| -l, --lyrics   | Lyrics file      |
| -c, --cover    | Cover image      |
| --custom       | Custom frame     |

---

### remove

| Flag           | Description         |
| -------------- | ------------------- |
| -t, --title    | Remove title        |
| -a, --artist   | Remove artist       |
| -A, --album    | Remove album        |
| --album-artist | Remove album artist |
| -g, --genre    | Remove genre        |
| -y, --year     | Remove year         |
| -n, --number   | Remove track        |
| -d, --disk     | Remove disk         |
| -l, --lyrics   | Remove lyrics       |
| -c, --cover    | Remove cover        |
| --custom       | Remove frame        |
| --all          | Remove all          |

---

### lyrics

| Flag         | Description     |
| ------------ | --------------- |
| -e, --editor | External editor |
