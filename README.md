# elacity-drm-cryptographyic-protocol
Cryptographic Protocol to manage key exchanges with smart contracts

### Development notes
What we are building here is a wasm-based component, debugging it or even testing as wasm compiled program is quite tricky, so for develoment we will use HTTP and keep all the core decoupled from its output (normal binary for HTTP and wasm for the final output)

Though you can find samples in `./samples` and serve the folder as static webserver with 
```shell
python -m SimpleHTTPServer 8000
```

Then go to [http://localhost:8000](http://localhost:8000)

```shell
cd samples & mkdir -p video

# not make sure to add a video in video folder
# fragment it
mp4fragment video/<source> video/frag.sample.mp4

# then encrypt

mp4dash \
-v -d --force \
--use-segment-list \
--use-segment-timeline \
--subtitles --clearkey \
--encryption-cenc-scheme=cenc \
--encryption-args="--global-option mpeg-cenc.eme-pssh:true --pssh bf8ef85d2c54475d8c1ee27db60332a2:pssh.json" \
--encryption-key=1064dbc48cc35c7fa724f6e723e73e49:933ea3404d78463c86fcbd1913323389 \
video/frag.sample.mp4 -o video/output
```

Not here we have a new system Id `bf8ef85d-2c54-475d-8c1e-e27db60332a2` where we set PSSH information. This key system is not natively suported by any browser but the dash manifest and media controller can handle it through EME `encrypted` event as a Common Encryption derivated. The most important is that we can register it during encryption and set it in media header as [PSSH Box](https://www.w3.org/TR/eme-initdata-cenc/#common-system), we can trigger our own flow from it and use thes specific data from the box to secure the license exchange flow with the wasm component we are building here.

```json
{
  "systemId": "bf8ef85d-2c54-475d-8c1e-e27db60332a2",
  "keySystem": "eth.web3.clearkey",
  "data": {
    "authority": "0x44bC6b5B64E78Aa14a2f82E2D97026444557F746",
    "chainId": 1337,
    "rpc": "http://localhost:8545"
  }
}
```