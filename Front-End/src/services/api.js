const baseURL = "http://localhost:9000/api";
const defaultHeaders = {
  "Content-Type": "application/json",
};

async function request(path, options = {}) {
  const response = await fetch(`${baseURL}${path}`, {
    headers: defaultHeaders,
    ...options,
    body: options.body ? JSON.stringify(options.body) : undefined,
  });

  if (!response.ok) {
    const errorText = await response.text();
    throw new Error(`HTTP ${response.status}: ${errorText}`);
  }

  return response.json();
}

const api = {
  get(path) {
    return request(path, { method: "GET" });
  },
  post(path, body) {
    return request(path, { method: "POST", body });
  },
  put(path, body) {
    return request(path, { method: "PUT", body });
  },
  delete(path) {
    return request(path, { method: "DELETE" });
  },
};

export default api;
