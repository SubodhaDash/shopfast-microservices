from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List
import os

app = FastAPI(title="Product Service", version="1.0.0")

# In-memory product store (real app would use a DB)
PRODUCTS = [
    {"id": 1, "name": "Laptop", "price": 999.99, "stock": 50},
    {"id": 2, "name": "Mouse", "price": 29.99, "stock": 200},
    {"id": 3, "name": "Keyboard", "price": 79.99, "stock": 150},
]

class Product(BaseModel):
    id: int
    name: str
    price: float
    stock: int

@app.get("/health")
def health_check():
    return {
        "status": "healthy",
        "service": "product-service",
    }

@app.get("/products", response_model=List[Product])
def get_products():
    return PRODUCTS

@app.get("/products/{product_id}", response_model=Product)
def get_product(product_id: int):
    product = next((p for p in PRODUCTS if p["id"] == product_id), None)
    if not product:
        raise HTTPException(status_code=404, detail="Product not found")
    return product