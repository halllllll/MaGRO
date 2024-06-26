/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as SelectunitImport } from './routes/select_unit'
import { Route as RedirectImport } from './routes/redirect'
import { Route as LoginImport } from './routes/login'
import { Route as IndexImport } from './routes/index'
import { Route as UserIndexImport } from './routes/user/index'

// Create Virtual Routes

const AboutLazyImport = createFileRoute('/about')()

// Create/Update Routes

const AboutLazyRoute = AboutLazyImport.update({
  path: '/about',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/about.lazy').then((d) => d.Route))

const SelectunitRoute = SelectunitImport.update({
  path: '/select_unit',
  getParentRoute: () => rootRoute,
} as any)

const RedirectRoute = RedirectImport.update({
  path: '/redirect',
  getParentRoute: () => rootRoute,
} as any)

const LoginRoute = LoginImport.update({
  path: '/login',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const UserIndexRoute = UserIndexImport.update({
  path: '/user/',
  getParentRoute: () => rootRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/login': {
      id: '/login'
      path: '/login'
      fullPath: '/login'
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    '/redirect': {
      id: '/redirect'
      path: '/redirect'
      fullPath: '/redirect'
      preLoaderRoute: typeof RedirectImport
      parentRoute: typeof rootRoute
    }
    '/select_unit': {
      id: '/select_unit'
      path: '/select_unit'
      fullPath: '/select_unit'
      preLoaderRoute: typeof SelectunitImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutLazyImport
      parentRoute: typeof rootRoute
    }
    '/user/': {
      id: '/user/'
      path: '/user'
      fullPath: '/user'
      preLoaderRoute: typeof UserIndexImport
      parentRoute: typeof rootRoute
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  IndexRoute,
  LoginRoute,
  RedirectRoute,
  SelectunitRoute,
  AboutLazyRoute,
  UserIndexRoute,
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/login",
        "/redirect",
        "/select_unit",
        "/about",
        "/user/"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/login": {
      "filePath": "login.tsx"
    },
    "/redirect": {
      "filePath": "redirect.tsx"
    },
    "/select_unit": {
      "filePath": "select_unit.tsx"
    },
    "/about": {
      "filePath": "about.lazy.tsx"
    },
    "/user/": {
      "filePath": "user/index.tsx"
    }
  }
}
ROUTE_MANIFEST_END */
