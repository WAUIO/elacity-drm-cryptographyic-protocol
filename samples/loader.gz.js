// eslint-disable-next-line func-names
(async function (wasmURL, w) {
  const fetcher = async (url) => {
    const response = await fetch(url);
    let buf = await response.arrayBuffer();

    buf = pako.ungzip(buf);
    if (buf[0] === 0x1f && buf[1] === 0x8b) {
      console.info('2nd ungzip')
      buf = pako.ungzip(buf);
    }

    return buf;
  }

  const go = new w.Go();
  fetcher(wasmURL).then((buf) => {
    return WebAssembly.instantiate(buf, go.importObject);
  }).then(
    (result) => {
      return go.run(result.instance);
    }
  ).catch(console.error);
}('/debug.crypto.protocol.wasm.gz', window));
