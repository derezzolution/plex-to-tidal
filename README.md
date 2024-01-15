# plex-to-tidal

![Under Construction](https://github.com/derezzolution/plex-to-tidal/blob/main/docs/under-construction.gif)

## Notes

* Need to figure out the best primary key between Plex Media Server responses and Tidal. This will likely be a problem
  since there doesn't seem to be any overlap.

* Tidal has support for ISRC (International Standard Recording Code) but it doesn't look like Plex does. :(

## Usage

1. Copy `config.json.template` to `config.json` and populate with your specifics

1. Run with `go run .`
