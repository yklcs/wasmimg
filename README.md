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

Even if a particular language isn't supported, Cram's .wasm files can be used with any WASM runtime supporting the below runtime requirements.

## Runtime requirements

`wasi_snapshot_preview1` support is required.

Libraries are compiled with `ALLOW_MEMORY_GROWTH=1` via Emscripten.
This flag requires `emscripten_notify_memory_growth` to be defined. Most popular WASM runtimes support this - [otherwise, you can easily define your own.](https://github.com/zeux/meshoptimizer/blob/bdc3006532dd29b03d83dc819e5fa7683815b88e/js/meshopt_decoder.js#L10)

```wat
(import "env" "emscripten_notify_memory_growth" (func (;3;) (type 0)))
```
