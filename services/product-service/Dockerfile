# Build Stage
FROM python:3.12-slim AS builder
WORKDIR /app
COPY requirements.txt ./

RUN pip install --no-cache-dir --user -r requirements.txt

# Production Stage
FROM python:3.12-slim
WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY  --from=builder /root/.local /root/.local
COPY /app /app/

USER appuser

ENV PATH=/home/appuser/.local/bin:$PATH
ENV PYTHONUNBUFFERED=1

EXPOSE 5000

HEALTHCHECK --interval=30s --timeout=30s CMD python -c "import urllib.request; urllib.request.urlopen('http://localhost:5000/health')" || exit 1

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]