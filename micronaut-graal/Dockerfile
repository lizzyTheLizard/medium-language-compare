FROM oracle/graalvm-ce:20.2.0-java11 AS graalvm
RUN gu install native-image

WORKDIR /home/app
COPY build/layers/libs /home/app/libs
COPY build/layers/resources /home/app/resources
COPY build/layers/application.jar /home/app/application.jar
RUN native-image -H:Class=lizzy.medium.compare.micronaut.Application -H:Name=application --no-fallback -cp /home/app/libs/*.jar:/home/app/resources:/home/app/application.jar

FROM frolvlad/alpine-glibc
RUN apk update && apk add libstdc++
EXPOSE 8080
COPY --from=graalvm /home/app/application /app/application
ENTRYPOINT ["/app/application"]
