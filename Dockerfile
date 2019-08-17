FROM scratch

COPY franz /

ENTRYPOINT ["/franz"]
