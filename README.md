# m3u8 downloader ![Go](https://github.com/q-meet/m3u8Tool/workflows/Go/badge.svg?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/q-meet/m3u8Tool)](https://goreportcard.com/report/github.com/q-meet/m3u8Tool)
A M3U8 downloader (CLI) written in Golang to download HTTP Live Streaming videos, forked from [oopsguy/m3u8](https://github.com/oopsguy/m3u8) with continuous updates.

You only need to specify the flags(`u`, `o`, `c`) to run, downloader will automatically download all TS files and consolidate them into a single TS file.

## Features

- Download and parse M3U8（VOD）
- Retry on download TS failure
- Parse Master playlist
- Decrypt TS
- Merge TS

## Usage

### Build

```bash
git clone https://github.com/q-meet/m3u8Tool.git
cd github.com/q-meet/m3u8Tool && go build
```

### Download

- [Build Artifacts](https://github.com/q-meet/m3u8Tool/actions)
- [Release Packages](https://github.com/q-meet/m3u8Tool/releases)

### Execution:

Linux

```
./github.com/q-meet/m3u8Tool-linux -u=http://example.com/index.m3u8 -o=/data/example
```

Windows PowerShell

```
.\github.com/q-meet/m3u8Tool-win.exe -u="http://example.com/index.m3u8" -o="D:\data\example"
```

MacOS

```
./github.com/q-meet/m3u8Tool-mac -u=http://example.com/index.m3u8 -o=/data/example
```

## References

- [HTTP Live Streaming draft-pantos-http-live-streaming-23](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23#section-4.3.4.2)
- [MPEG transport stream - Wikipedia](https://en.wikipedia.org/wiki/MPEG_transport_stream)

## License

[MIT License](./LICENSE)
