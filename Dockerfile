FROM scratch
EXPOSE 8080
ENTRYPOINT ["/awsappmesh-meetup"]
COPY ./bin/ /