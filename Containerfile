FROM scratch

COPY main.wasm /app/bin/server.wasm

CMD ["/app/bin/server.wasm"]
