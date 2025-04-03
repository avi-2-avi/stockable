import { vi } from "vitest";

let mockCookies: Record<string, string> = {};

const axiosMock = {
  post: vi.fn((url) => {
    if (url.includes("/api/login")) {
      mockCookies["session"] = "mockSessionToken"; 
      return Promise.resolve({ data: { success: true } });
    }
    return Promise.resolve({ data: {} });
  }),
  get: vi.fn((url, config) => {
    if (config?.withCredentials && !mockCookies["session"]) {
      return Promise.reject(new Error("Unauthorized: No session cookie"));
    }
    return Promise.resolve({ data: {} });
  }),
};

export default axiosMock;
