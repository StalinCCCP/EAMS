<script setup>
import http from "@/util/http.js";
import {reactive} from "vue";
import {createToast} from "mosha-vue-toastify";

const data = reactive({
  form: {
    Username: '',
    Pwd: '',
    //re_enter_password: ''
  },
  re_enter_password:''
})

const register = () => {
  if (data.form.Pwd === data.re_enter_password) {
    http.post("/p/register", data.form,{validateStatus: function (status) {
    return true; 
  }}).then(r => {
      if (r.status===200) {
        createToast("Registered successfully")
        localStorage.setItem("Authorization", r.data.data.token)
        router.back()
      }
      else
        createToast("Register failed:"+r.data.msg)
    })
  } else {
    createToast("Passwords doesn't match")
  }
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
        <div class="mb-3">
          <label class="form-label">Re-enter the password</label>
          <input type="password" class="form-control" v-model="data.re_enter_password">
        </div>
        <p class="text-muted">
          Already have an
          <router-link to="/login" class="text-reset">account?</router-link>
        </p>
        <div class="d-grid gap-2">
          <button type="button" class="btn btn-primary d-grid gap-2" @click="register">Register</button>
        </div>
      </form>
    </div>
  </div>
</template>