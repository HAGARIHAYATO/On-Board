import Vue from 'vue'
import Router from 'vue-router'
import { interopDefault } from './utils'
import scrollBehavior from './router.scrollBehavior.js'

const _cde162a8 = () => interopDefault(import('../pages/config.vue' /* webpackChunkName: "pages/config" */))
const _10ba8d22 = () => interopDefault(import('../pages/login.vue' /* webpackChunkName: "pages/login" */))
const _490e6182 = () => interopDefault(import('../pages/signup.vue' /* webpackChunkName: "pages/signup" */))
const _c91b56de = () => interopDefault(import('../pages/users/index.vue' /* webpackChunkName: "pages/users/index" */))
const _a3045cea = () => interopDefault(import('../pages/works/index.vue' /* webpackChunkName: "pages/works/index" */))
const _b4e2e0ce = () => interopDefault(import('../pages/works/new.vue' /* webpackChunkName: "pages/works/new" */))
const _1c03718e = () => interopDefault(import('../pages/users/_id.vue' /* webpackChunkName: "pages/users/_id" */))
const _27ad2ef6 = () => interopDefault(import('../pages/works/_id/index.vue' /* webpackChunkName: "pages/works/_id/index" */))
const _b75b84d4 = () => interopDefault(import('../pages/works/_id/edit.vue' /* webpackChunkName: "pages/works/_id/edit" */))
const _425d414c = () => interopDefault(import('../pages/works/_id/edit_item.vue' /* webpackChunkName: "pages/works/_id/edit_item" */))
const _2dfb1658 = () => interopDefault(import('../pages/index.vue' /* webpackChunkName: "pages/index" */))

// TODO: remove in Nuxt 3
const emptyFn = () => {}
const originalPush = Router.prototype.push
Router.prototype.push = function push (location, onComplete = emptyFn, onAbort) {
  return originalPush.call(this, location, onComplete, onAbort)
}

Vue.use(Router)

export const routerOptions = {
  mode: 'history',
  base: decodeURI('/'),
  linkActiveClass: 'nuxt-link-active',
  linkExactActiveClass: 'nuxt-link-exact-active',
  scrollBehavior,

  routes: [{
    path: "/config",
    component: _cde162a8,
    name: "config"
  }, {
    path: "/login",
    component: _10ba8d22,
    name: "login"
  }, {
    path: "/signup",
    component: _490e6182,
    name: "signup"
  }, {
    path: "/users",
    component: _c91b56de,
    name: "users"
  }, {
    path: "/works",
    component: _a3045cea,
    name: "works"
  }, {
    path: "/works/new",
    component: _b4e2e0ce,
    name: "works-new"
  }, {
    path: "/users/:id",
    component: _1c03718e,
    name: "users-id"
  }, {
    path: "/works/:id",
    component: _27ad2ef6,
    name: "works-id"
  }, {
    path: "/works/:id/edit",
    component: _b75b84d4,
    name: "works-id-edit"
  }, {
    path: "/works/:id/edit_item",
    component: _425d414c,
    name: "works-id-edit_item"
  }, {
    path: "/",
    component: _2dfb1658,
    name: "index"
  }],

  fallback: false
}

export function createRouter () {
  return new Router(routerOptions)
}
