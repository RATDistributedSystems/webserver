FROM scratch

COPY frontend server tserver.json Caddyfile /app/
WORKDIR "/app"
EXPOSE 44440
CMD ["./server"]
