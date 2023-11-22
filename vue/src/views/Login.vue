<script setup>
import http from '@/util/http';
import router from "@/router/index.js";
import {reactive} from "vue";
import {createToast} from "mosha-vue-toastify";

// export default {
//   name: 'Login',
//   components: {
//     Login
//   }
// }

const data = reactive({
  form: {
    Username: '',
    Pwd: ''
  }
})

const login = () => {
  http.post("/p/login", data.form,{validateStatus: function (status) {
    return true; 
  }}).then(r => {
    if (r.status === 200) {
      localStorage.setItem("Authorization", r.data.data.token)
      createToast("Login successfully")
      router.back()
    } else {
        createToast("Login failed:"+r.data.msg)
    }
  });
}
</script>
<template>
  <div class="vh-100 d-flex justify-content-center align-items-center">
    <div class="shadow-sm p-3 mb-5 bg-body rounded">
      <form>
        <div class="mb-3">
          <label class="form-label">Username</label>
          <input type="email" class="form-control" v-model="data.form.Username">
        </div>
        <div class="mb-3">
          <label class="form-label">Password</label>
          <input type="password" class="form-control" v-model="data.form.Pwd">
        </div>
        <p class="text-muted">
          Don't have an
          <router-link to="/register" class="text-reset">account</router-link>
          yet?
        </p>
        <div class="d-grid gap-2">
          <button type="button" class="btn btn-primary d-grid gap-2" @click="login">Login</button>
        </div>
      </form>
    </div>
  </div>
</template>