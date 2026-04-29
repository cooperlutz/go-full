import { authRefreshMiddleware } from '~/app/utils/middleware'

import { Configuration } from '../services/runtime'

export const BackendConfig = new Configuration({
  basePath: '',
  credentials: 'include',
  middleware: [authRefreshMiddleware],
})
