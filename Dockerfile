FROM scratch

COPY frontend webserver /app/
WORKDIR "/app"
EXPOSE 44440
CMD ["./webserver"]
