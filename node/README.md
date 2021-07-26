# Node Greeting Services

## Notes

- Node services still a WIP
- Year service is on its own for now (is not yet called in any way by greeting service)
- Frontend and Message services use beelines
- Year service uses Open Telemetry

## How to run

**Required**: set environment variables

- HONEYCOMB_API
- HONEYCOMB_API_KEY
- HONEYCOMB_DATASET
- OTLP_ENDPOINT

### Option 1: One service at a time

- In each individual service directory, run `npm start`

### Option 2: All Node Services via Docker

- In Node directory, run `docker-compose up --build`

### Option 3: All Node Services via Tilt

- In Tiltfile, uncomment Node services (and comment out go services)
- In top-level directory run `tilt up`

## See it in action

`curl localhost:7000/greeting` for greeting

`curl localhost:6001` for year only

`curl localhost:9000/message` for message only
