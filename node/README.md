# Node Greeting Services

## Notes

- Frontend, message, and name services use beelines
- Year service uses vanilla OTel

## How to run

**Required**: set environment variables

- HONEYCOMB_API_KEY
- HONEYCOMB_DATASET

### Install Node
If you're using option 1 or 3 listed below, you will have to install node on your machine. The preferred way to install node is with [asdf](https://asdf-vm.com/guide/getting-started.html#_1-install-dependencies).

Once you have asdf set up, run the following to install node. Check the version needed in the `.tool-versions` file in the root of this repo.

```
asdf plugin add nodejs https://github.com/asdf-vm/asdf-nodejs.git
```

```
asdf install nodejs <insert-version-here>
```

### Option 1: One service at a time

- In each individual service directory, run `npm start`

### Option 2: All Node Services via Docker

- In Node directory, run `docker-compose up --build`

### Option 3: All Node Services via Tilt

In top-level directory run `tilt up node`

## See it in action

`curl localhost:7000/greeting` for greeting

`curl localhost:9000/message` for message only

`curl localhost:8000/name` for name only

`curl localhost:6001` for year only
