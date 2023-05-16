# lfm2ytm: YouTube (Music) playlist generator with Last.fm hot tracks

lfm2ytm generates a YouTube (Music) playlist consisting of hot tracks (i.e., most played songs) in [Last.fm](https://www.last.fm/).

## Features

- Generate a YouTube (Music) playlist automatically based on playcounts in Last.fm
- Try to search [art tracks](https://support.google.com/youtube/answer/6007071) in YouTube
- Only print hot tracks in Last.fm running with `--dryrun`

## Example

### My Last.fm hot tracks in 11 July 2017 - 18 July 2017

![lastfm](https://user-images.githubusercontent.com/7788821/235303633-b332787b-4171-49c3-a3ab-8cf6b61b4a1b.png)

### Run generator

```bash
./lfm2ytm \
  -secretjson=./client_secret.json \
  -fromdate=2017-07-11T00:00:00+09:00 \
  -todate=2017-07-18T23:59:59+09:00 \
  -region=JP \
  -lfmmax=10 \
  tsubasaogawa
```

Results

```text
✔ Start A Fire - John Legend - Topic (3 times; id=dwTRnPtpSSk)
✔ Herman’s Habit - Justin Hurwitz - Topic (2 times; id=zwiy2EfS0GY)
✔ Planetarium - Justin Hurwitz - Topic (2 times; id=qV1KxmGDkk8)
✔ City Of Stars (Humming) - Justin Hurwitz - Topic (2 times; id=AR1YYN4-qYE)
✔ &#39;City of Stars&#39; (Duet ft. Ryan Gosling, Emma Stone) - La La Land Original Motion Picture Soundtrack - Interscope Records (2 times; id=GTWqwSNQCcg)
✔ 2 Unlimited - No Limit (RM Radio Edit) - Dove White (1 times; id=q-6EgZZHuds)
✔ Chumbawamba - Tubthumping - ChumbawambaVEVO (1 times; id=2H5uWRjFsGc)
✔ CUBIC 22 / NIGHT IN MOTION (K-Groove Mix) - DancemaniaMania (1 times; id=DN8_52mDDw4)
✔ Dj Miko - What&#39;s Up 2000 - Saifam Music (1 times; id=lG86oPotqRU)
✔ Green Flower Street - Donald Fagen - Topic (1 times; id=SoHyj8i8Ywo)
```

### Generated YouTube Music playlist

![ytmusic](https://user-images.githubusercontent.com/7788821/235303635-479550fd-a866-4816-9252-949cd91da8df.png)

## How to use

### 1. Enable YouTube Data API v3

https://developers.google.com/youtube/v3/getting-started

### 2. Generate OAuth 2.0 client secret file

https://developers.google.com/youtube/v3/guides/auth/installed-apps

- Choose "Desktop Application"

### 3. Run

See the example.

## Options

See [command_help.md](command_help.md)

## Development

1. Fork this repository.
2. Chdir and `go mod tidy`
3. Do implements.
4. Run `LASTFM_API_KEY=<your lastfm api key> go run *.go [options]`
5. Create a pull-request.
