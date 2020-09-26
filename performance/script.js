import http from 'k6/http';
import { check, sleep } from 'k6';

// 1. init code
export let options = {
  vus: 1000,
  stages: [
    { duration: '50s', target: 60 },
    { duration: '10s', target: 60 },
    { duration: '30s', target: 100 },
    { duration: '20s', target: 100 },
    { duration: '30s', target: 60 },
    { duration: '10s', target: 60 },
    { duration: '50s', target: 0 },
  ],
  thresholds: {
    errors: ['rate<0.1'],
    http_req_duration: ['p(95)<1000','p(99)<5000'], // 99% of requests must complete below 1.5s
  },
};

export function setup() {
  // 2. setup code
}

export default function () {
  let res = http.get('https://httpbin.org/');
  check(res, { 'status was 200': r => r.status == 200 });
  sleep(1);
}

export function teardown(data) {
  // 4. teardown code
}