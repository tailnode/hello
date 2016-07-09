FROM scratch
COPY hello /hello
ENV HELLO_VERSION 1.0
EXPOSE 8080

CMD ["/hello"]
