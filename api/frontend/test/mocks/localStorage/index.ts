interface StorageMap {
  [key: string]: string;
}

interface LocalStorageMock {
  getItem: (key: string) => string | null;
  setItem: (key: string, value: string) => void;
  removeItem: (key: string) => void;
  clear: () => void;
}

const localStorageMock: LocalStorageMock = (() => {
  let store: StorageMap = {};
  return {
    // if key == 'access_token' return a dummy token for testing purposes
    getItem: (key: string): string | null =>
      key === "access_token" ? "dummy-token" : store[key] || null,
    setItem: (key: string, value: string): void => {
      store[key] = value.toString();
    },
    removeItem: (key: string): void => {
      delete store[key];
    },
    clear: (): void => {
      store = {};
    },
  };
})();

Object.defineProperty(window, "localStorage", {
  value: localStorageMock,
});
