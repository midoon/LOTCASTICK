const baseURL = import.meta.env.VITE_BE_BASEURL;
const defaultHeaders = {
  "Content-Type": "application/json",
};

async function refreshAccessToken() {
  const refreshToken = localStorage.getItem("refreshToken");
  if (!refreshToken) {
    return null;
  }

  const response = await fetch(`${baseURL}/auth/refresh`, {
    method: "POST",
    headers: defaultHeaders,
    body: JSON.stringify({ refreshToken }),
  });

  if (!response.ok) {
    return null;
  }

  const data = await response.json();
  if (data.accessToken) {
    localStorage.setItem("accessToken", data.accessToken);
  }

  return data.accessToken || null;
}

async function request(path, options = {}, retry = true) {
  const headers = {
    ...defaultHeaders,
    ...(options.headers || {}),
  };

  const accessToken = localStorage.getItem("accessToken");
  if (accessToken) {
    headers.Authorization = `Bearer ${accessToken}`;
  }

  const response = await fetch(`${baseURL}${path}`, {
    ...options,
    headers,
    body: options.body != null ? JSON.stringify(options.body) : undefined,
  });

  if (response.status === 401 && retry) {
    const newToken = await refreshAccessToken();
    if (newToken) {
      return request(path, options, false);
    }
  }

  if (!response.ok) {
    const errorData = await response.json();

    throw {
      status: response.status,
      ...errorData,
    };
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
