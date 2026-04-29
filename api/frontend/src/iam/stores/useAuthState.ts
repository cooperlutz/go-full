/**
 * Tracks whether the user has authenticated in this session.
 * This is a UX hint only — actual security is enforced by the httpOnly cookie.
 * If tampered with, API calls will return 401 and the middleware will redirect to /login.
 */
export function useAuthState() {
  const isAuthenticated = () => sessionStorage.getItem('auth') === '1'
  const setAuthenticated = () => sessionStorage.setItem('auth', '1')
  const clearAuthenticated = () => sessionStorage.removeItem('auth')

  return { isAuthenticated, setAuthenticated, clearAuthenticated }
}
