# Build Stage
FROM python:3.12-slim AS builder
WORKDIR /app
COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

# Production Stage
FROM python:3.12-slim
WORKDIR /app

RUN addgroup --system appgroup && adduser --system --group appuser

COPY  --from=builder /usr/local /usr/local
COPY . /app/

RUN chown -R appuser:appgroup /app
USER appuser

ENV PYTHONUNBUFFERED=1

EXPOSE 5000

HEALTHCHECK --interval=30s --timeout=30s CMD python -c "import urllib.request; urllib.request.urlopen('http://localhost:5000/health')" || exit 1

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]