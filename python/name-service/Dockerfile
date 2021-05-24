FROM python:alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev python3-dev libffi-dev openssl-dev cargo

ENV PYTHONFAULTHANDLER=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONHASHSEED=random \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PIP_DEFAULT_TIMEOUT=100 \
  POETRY_VERSION=1.0.0

RUN pip install "poetry==$POETRY_VERSION"

COPY poetry.lock pyproject.toml ./
RUN poetry config virtualenvs.create false \
  && poetry install --no-root --no-interaction --no-ansi

COPY ./.flaskenv ./
COPY ./name_service ./name_service

EXPOSE 8000
CMD ["poetry", "run", "flask", "run"]