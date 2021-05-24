FROM gradle:7.0.1-jdk11 AS build
WORKDIR /home/gradle/
RUN mkdir ./libs
ENV HONEY_OTEL_VER=0.1.1
ENV HONEY_JAVA_AGENT=honeycomb-opentelemetry-javaagent-${HONEY_OTEL_VER}-all.jar
ADD https://github.com/honeycombio/honeycomb-opentelemetry-java/releases/download/v${HONEY_OTEL_VER}/${HONEY_JAVA_AGENT} ./libs
COPY --chown=gradle:gradle *.gradle ./
COPY --chown=gradle:gradle ./src ./src
RUN gradle bootJar --no-daemon

FROM openjdk:13-jdk-alpine
VOLUME /tmp
ENV HONEY_OTEL_VER=0.1.1
ENV HONEY_JAVA_AGENT=honeycomb-opentelemetry-javaagent-${HONEY_OTEL_VER}-all.jar
ENV JAVA_TOOL_OPTIONS=-javaagent:${HONEY_JAVA_AGENT}
COPY --from=build /home/gradle/libs/${HONEY_JAVA_AGENT} ./${HONEY_JAVA_AGENT}
COPY --from=build /home/gradle/build/libs/*.jar app.jar
CMD ["java", "-jar", "app.jar"]
