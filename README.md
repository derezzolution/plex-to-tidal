# plex-to-tidal

![Under Construction](https://github.com/derezzolution/plex-to-tidal/blob/main/docs/under-construction.gif) ...lololololl

## Notes

* Need to figure out the best primary key between Plex Media Server responses and Tidal. This will likely be a problem
  since there doesn't seem to be any overlap.

* Tidal has support for ISRC (International Standard Recording Code) but it doesn't look like Plex does. :(

* We can query Tidal tracks by ISRC code: <https://github.com/orgs/tidal-music/discussions/26>

* Tidal github discussions (use this for support since Tidal support won't talk to you):
  <https://github.com/orgs/tidal-music/discussions>

> [!WARNING]
> If we can get a viable POC, we'll DRY up the code and wrap with a service layer. But we NEEEEED a reasonable primary
> key. I don't want to fallback to search API and take the first result nor incorporate some LLM madness. Also the
> [Tidal API](https://github.com/orgs/tidal-music/discussions/39) seems like it's read-only. Multiple blocking issues
> here.

## Usage

1. Copy `config.json.template` to `config.json` and populate with your specifics

1. Run with `go run .`
