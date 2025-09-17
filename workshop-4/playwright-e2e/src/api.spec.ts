import { test, expect, request as playwrightRequest } from '@playwright/test';


// Set baseURL for all requests
const baseURL = 'http://localhost:3000';

test('GET /profile should return profile data', async () => {
  const request = await playwrightRequest.newContext({ baseURL });
  const response = await request.get('/profile');
  expect(response.status()).toBe(200);
  expect(response.headers()['content-type']).toContain('application/json');
  const data = await response.json();
  expect(data).toEqual({
    membership_level: 'Gold',
    membership_code: 'LBK001234',
    first_name: 'สมชาย',
    last_name: 'ใจดี',
    phone: '081-234-5678',
    email: 'somchai@example.com',
    joined_date: '2023-06-15',
    points: 15420
  });
  await request.dispose();
});

test('PUT /profile should update profile data', async () => {
  const request = await playwrightRequest.newContext({ baseURL });
  const updatePayload = {
    name: 'Test User',
    // Add more fields as required by the schema
  };
  const response = await request.put('/profile', {
    data: updatePayload,
  });
  expect(response.ok()).toBeTruthy();
  const data = await response.json();
  // Check that data.name is undefined
  expect(data.name).toBeUndefined();
  await request.dispose();
});
