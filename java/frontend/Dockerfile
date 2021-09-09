FROM gradle:7.0.1-jdk11 AS build
WORKDIR /home/gradle/
RUN mkdir ./libs
ENV OTEL_AGENT_VER=1.5.3
ENV JAVA_AGENT=opentelemetry-javaagent-all.jar
ADD https://github.com/open-telemetry/opentelemetry-java-instrumentation/releases/download/v${OTEL_AGENT_VER}/${JAVA_AGENT} ./libs
COPY --chown=gradle:gradle *.gradle ./
COPY --chown=gradle:gradle ./src ./src
RUN gradle bootJar --no-daemon

FROM openjdk:13-jdk-alpine
VOLUME /tmp
ENV OTEL_AGENT_VER=1.5.3
ENV JAVA_AGENT=opentelemetry-javaagent-all.jar
ENV JAVA_TOOL_OPTIONS=-javaagent:${JAVA_AGENT}
COPY --from=build /home/gradle/libs/${JAVA_AGENT} ./${JAVA_AGENT}
COPY --from=build /home/gradle/build/libs/*.jar app.jar
CMD ["java", "-jar", "app.jar"]
