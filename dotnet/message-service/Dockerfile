# https://hub.docker.com/_/microsoft-dotnet
FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /source

# copy csproj and restore as distinct layers
COPY *.cs .
COPY *.csproj .
COPY *.json .
COPY ./Properties ./Properties
COPY ./Controllers ./Controllers
RUN dotnet restore

# copy everything else and build app
RUN dotnet publish -c release -o /app --no-restore

# final stage/image
FROM mcr.microsoft.com/dotnet/aspnet:5.0
WORKDIR /app
COPY --from=build /app ./

ENV ASPNETCORE_URLS=http://+:9000
ENV ASPNETCORE_ENVIRONMENT=”development”
EXPOSE 9000
ENTRYPOINT ["dotnet", "message-service.dll"]
