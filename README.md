# Self-host Planning Poker

A hassle-free Planning Poker application to deploy everywhere.

This is sort of a fork from [axeleroy](https://github.com/axeleroy/self-host-planning-poker)'s planning poker, but slightly adjusted to my likings.

[![GitHub last commit](https://img.shields.io/github/last-commit/the-technat/self-host-planning-poker?logo=github&logoColor=959DA5)](https://github.com/the-technat/self-host-planning-poker/commits/main)
[![License](https://img.shields.io/github/license/axeleroy/self-host-planning-poker?logo=github&logoColor=959DA5)](https://github.com/axeleroy/self-host-planning-poker/blob/main/LICENSE)
[![Tests](https://github.com/axeleroy/self-host-planning-poker/actions/workflows/tests.yml/badge.svg)](https://github.com/axeleroy/self-host-planning-poker/actions/workflows/tests.yml)
[![Docker build](https://github.com/the-technat/self-host-planning-poker/actions/workflows/publish.yml/badge.svg)](https://github.com/the-technat/self-host-planning-poker/actions/workflows/publish.yml)

## What is it?

This application is intended as a simplified and self-hostable alternative to
[Planning Poker Online](https://planningpokeronline.com/).

It features:

  * Multiple deck types: Fibonacci, modified Fibonacci, T-Shirt sizes, powers of 2 and trust vote (0 to 5)
  * Spectator mode
  * Responsive layout
  * Vote summary
  * Translations _(English, French, German, Italian and Polish.)_
 
It does not have fancy features like issues management, Jira integration or timers.

## Screenshots
<a href="https://github.com/the-technat/self-host-planning-poker/blob/main/assets/screenshot.png"><img alt="Application screenshot with cards face down" src="https://github.com/the-technat/self-host-planning-poker/blob/main/assets/screenshot.png" width="412px"></a>
<a href="https://github.com/the-technat/self-host-planning-poker/blob/main/assets/screenshot.png"><img alt="Application screenshot with cards revealed" src="https://github.com/the-technat/self-host-planning-poker/blob/main/assets/screenshot-revealed.png" width="412px"></a>

## Deployment

Deploying the application is easy as it's self-contained in a single container.
All you need is to create a volume to persist the games settings (ID, name and deck).

### Docker
```bash
docker run \
  -v planning-poker-data:/data \
  -p 8000:8000 \
  ghcr.io/the-technat/self-host-planning-poker:latest
```

### docker-compose
```yml
version: "3"
services:
  planning-poker:
    image: ghcr.io/the-technat/self-host-planning-poker:latest
    ports:
      - 8000:8000
    volumes:
      - planning-poker-data:/data
volumes:
  planning-poker-data: {}
```

### Helm chart

There is a helm chart available at [charts/self-host-planning-poker](./charts/self-host-planning-poker).

## Development

The app consists of two parts:

* a [back-end](flask/) written in Python with [Flask](https://flask.palletsprojects.com/), [Flask-SocketIO](https://flask-socketio.readthedocs.io/en/latest/index.html) and [peewee](http://docs.peewee-orm.com/en/latest/).
* a [front-end](angular/) written with [Angular](https://angular.io) and [Socket.IO](https://socket.io/).

### Back-end development

You must first initialise a virtual environment and install the dependencies

```sh
# Run the following commands in the flask/ folder
python3 -m venv env
source env/bin/activate
pip3 install -r requirements.txt
```

Then launching the development server is as easy as that:
```bash
FLASK_DEBUG=1 python app.py
```

#### Run unit tests

After initializing the virtual environment, run this command in the `flask/` directory:
```sh
python -m unittest
```

### Front-end development

> <details>
> <summary>
> <b>Note:</b> You might want to test the front-end against a back-end. You can either follow the instructions in the
> previous section to install and run it locally or use the following command to run it in a Docker container:
> </summary>
>
> ```bash
> docker run --rm -it \
>   -v $(pwd)/flask:/app \
>   -p 5000:5000 \
>   python:3.11-slim \
>   bash -c "cd /app; pip install -r requirements.txt; FLASK_DEBUG=1 gunicorn --worker-class eventlet -w 1 app:app --bind 0.0.0.0:5000"
> ```
> </details>

First make sure that [Node.js](https://nodejs.org/en/) (preferably LTS) is installed.
Then, install dependencies and launch the development server

```sh
# Run the following commands in the angular/ folder
npm install
npm start
```

### Building Docker image

```sh
# After checking out the project
docker build . -t the-technat/self-host-planning-poker:custom
# Alternatively, if you don't want to checkout the project
docker build https://github.com/the-technat/self-host-planning-poker -t the-technat/self-host-planning-poker:custom
```
