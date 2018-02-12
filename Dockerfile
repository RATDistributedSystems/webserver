FROM scratch

COPY frontend server config.json /app/
WORKDIR "/app"
EXPOSE 44440
CMD ["./webserver"]
