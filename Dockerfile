FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY lorem-go /
CMD ["/lorem-go"]