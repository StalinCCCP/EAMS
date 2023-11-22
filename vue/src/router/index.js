import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: function () {
      return import(/* webpackChunkName: "about" */ '../views/HomeView.vue')
    }
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: function () {
      return import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/Login.vue')
    }
  },
  {
    path: '/register',
    name: 'Register',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/Register.vue')
    }
  },
  {
    path: '/init',
    name: 'Init',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/Init.vue')
    }
  },
  {
    path: '/hardware',
    name: 'Hardware',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/Hardware.vue')
    }
  },
  {
    path: '/hardware-detail',
    name: 'Hardware Detail',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/HardwareDetail.vue')
    }
  },
  {
    path: '/createhardware',
    name: 'CreateHardware',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/CreateHardware.vue')
    }
  },
  {
    path: '/hardware-maintenance-detail',
    name: 'Hardware Maintenance Detail',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/HardwareMaintenanceDetail.vue')
    }
  },
  {
    path: '/hardware-maintenance',
    name: 'Hardware Maintenance',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/HardwareMaintenance.vue')
    }
  },
  {
    path: '/createhardwaremaintenance',
    name: 'Create Hardware Maintenance',
    component: function (){
      return import(/* webpackChunkName: "about" */ '../views/CreateHardwareMaintenance.vue')
    }
  }
  
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
