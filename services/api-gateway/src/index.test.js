const request = require('supertest');
const app = require('./index');

describe('API Gateway', () => {
  it('GET /health should return 200', async () => {
    const res = await request(app).get('/health');
    expect(res.statusCode).toBe(200);
    expect(res.body.status).toBe('healthy');
    expect(res.body.service).toBe('api-gateway');
  });
});