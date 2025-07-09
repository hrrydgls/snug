import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  vus: 1000,
  duration: "10s",
};

export default function () {
  const randomId = Math.floor(Math.random() * 1000000);
  const email = `user${randomId}@example.com`;

  const url = "http://localhost:8888/auth/email";

  // manually create form body
  const payload = `email=${encodeURIComponent(email)}`;

  const headers = {
    "Content-Type": "application/x-www-form-urlencoded",
  };

  const res = http.post(url, payload, { headers });

  check(res, {
    "status is expected": () =>
      res.status === 200 ||
      res.status === 201 ||
      res.status === 422 ||
      res.status === 500,
  });

  sleep(0.1);
}
