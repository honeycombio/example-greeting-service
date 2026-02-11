FROM gradle:8.1.0-jdk17 AS build
WORKDIR /home/gradle/
RUN mkdir ./libs
COPY --chown=gradle:gradle *.gradle ./
COPY --chown=gradle:gradle ./src ./src
RUN gradle bootJar --no-daemon

FROM eclipse-temurin:17-jdk
VOLUME /tmp
ENV JAVA_AGENT=otel-javaagent.jar
ENV JAVA_TOOL_OPTIONS=-javaagent:${JAVA_AGENT}
ENV OTEL_EXPERIMENTAL_CONFIG_FILE=/otelconf.yaml
COPY --from=build /home/gradle/agent/${JAVA_AGENT} ./${JAVA_AGENT}
COPY --from=build /home/gradle/build/libs/*.jar app.jar
CMD ["java", "-jar", "app.jar"]
