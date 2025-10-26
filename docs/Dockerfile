FROM ghcr.io/astral-sh/uv:alpine

COPY . .

RUN uv sync --locked

CMD ["uv", "run", "mkdocs", "serve", "-a", "0.0.0.0:8000"]