# wasmimg 🗜️

> ⚠️ Work in progress

wasmimg provides image processing in native Go. Mature image libraries are compiled to WASM, then run with [wazero](https://wazero.io) allowing them to be used without CGo.

wasmimg focuses on encoding and decoding (for now), leaving image transformations to the user.

Features:

- No CGo
- Simple high level API
- Minimal configuration

## Support matrix

|            | mozjpeg | webp |
| :--------: | :-----: | :--: |
| encode     |   ✅    |  ❌  |
| decode     |   ✅    |  ❌  |

Even if a particular language isn't supported, wasmimg's .wasm files can be used with any WASM runtime supporting the below runtime requirements.
