FROM scratch
COPY hello /hello
EVN HELLO_VERSION 1.0

CMD ["/hello"]
