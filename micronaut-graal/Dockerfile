FROM oracle/graalvm-ce:20.1.0-java11 as graalvm
COPY . /home/app/micronaut
WORKDIR /home/app/micronaut
RUN gu install native-image -v --debug
RUN native-image --no-server -cp build/libs/micronaut-*-all.jar

FROM frolvlad/alpine-glibc
RUN apk update && apk add libstdc++
EXPOSE 8080
COPY --from=graalvm /home/app/micronaut/micronaut /app/micronaut
ENTRYPOINT ["/app/micronaut"]
