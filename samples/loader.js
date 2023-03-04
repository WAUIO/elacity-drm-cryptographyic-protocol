// eslint-disable-next-line func-names
(async function (wasmURL, w) {
  if (WebAssembly) {
    if (WebAssembly && !WebAssembly.instantiateStreaming) { // polyfill
      WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return WebAssembly.instantiate(source, importObject);
      };
    }
  }

  const go = new w.Go();
  WebAssembly.instantiateStreaming(fetch(wasmURL), go.importObject).then((result) => {
    return go.run(result.instance);
  }).catch(console.error);
}('https://cdn.jsdelivr.net/gh/WAUIO/elacity-drm-cryptographyic-protocol/releases/download/1.0.0-4-g8f6e7d7/crypto.protocol-1.0.0-4-g8f6e7d7.wasm', window));
