export function useLocalTokenStore() {
  const getAccessToken = () => {
    return localStorage.getItem("access_token") || "";
  };

  const setAccessToken = (token: string) => {
    localStorage.setItem("access_token", token);
  };

  const clearAccessToken = () => {
    localStorage.removeItem("access_token");
  };

  const getRefreshToken = () => {
    return localStorage.getItem("refresh_token") || "";
  };

  const setRefreshToken = (token: string) => {
    localStorage.setItem("refresh_token", token);
  };

  const clearRefreshToken = () => {
    localStorage.removeItem("refresh_token");
  };

  const clear = () => {
    localStorage.clear();
  };

  return {
    getAccessToken,
    setAccessToken,
    clearAccessToken,
    getRefreshToken,
    setRefreshToken,
    clearRefreshToken,
    clear,
  };
}
