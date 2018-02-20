FROM scratch

COPY frontend webserver config.json /app/
WORKDIR "/app"
EXPOSE 44440
CMD ["./webserver"]
