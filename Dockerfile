FROM scratch
COPY hello /hello
ENV HELLO_VERSION 2.0
EXPOSE 8080

CMD ["/hello"]
