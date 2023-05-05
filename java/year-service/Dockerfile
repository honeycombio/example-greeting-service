FROM gradle:7.0.1-jdk11 AS build
WORKDIR /home/gradle/
RUN mkdir ./libs
COPY --chown=gradle:gradle *.gradle ./
COPY --chown=gradle:gradle ./src ./src
RUN gradle bootJar --no-daemon

FROM eclipse-temurin:17-jdk
VOLUME /tmp
COPY --from=build /home/gradle/build/libs/*.jar app.jar
CMD ["java", "-jar", "app.jar"]
