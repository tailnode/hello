FROM scratch
COPY hello /hello
ENV HELLO_VERSION 2.0

CMD ["/hello"]
