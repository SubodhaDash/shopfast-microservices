from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

def test_health():
    response = client.get("/health")
    assert response.status_code == 200
    assert response.json()["status"] == "healthy"

def test_get_products():
    response = client.get("/products")
    assert response.status_code == 200
    assert len(response.json()) > 0

def test_get_product_by_id():
    response = client.get("/products/1")
    assert response.status_code == 200
    assert response.json()["id"] == 1

def test_get_product_not_found():
    response = client.get("/products/999")
    assert response.status_code == 404