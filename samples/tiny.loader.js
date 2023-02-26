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
}('/tiny.crypto.protocol.wasm', window));
