# cram 🗜️

> ⚠️ Work in progress

Cram provides cross-platform image compression, allowing you to cram as many pixels as possible into a smaller number of bytes. Mature image libraries are compiled to WASM, allowing them to be used in any platform supporting WASM (+ WASI).

Cram focuses on compression (for now), leaving decoding and image transformations to the user. Most standard libraries have better decoding support than encoding.

Features:

- Language agnostic (powered by WASM)
- Simple high level API
- Opinionated, minimal configuration

## Support matrix

|            | mozjpeg | webp |
| :--------: | :-----: | :--: |
|     Go     |   ⚠️    |  ❌  |
| JavaScript |   ❌    |  ❌  |

## Development

Use good old `make`.
