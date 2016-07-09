FROM scratch
COPY hello /hello
ENV HELLO_VERSION 1.0

CMD ["/hello"]
