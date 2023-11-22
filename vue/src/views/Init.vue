<script setup>
import http from '@/util/http.js';
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
    Server: '',
    DBName: '',
    User: '',
    Password: '',
    Username: '',
    Pwd: '',
    Full_name: '',
    Email: '',
    Phone_number: '',
  }
})



const init = () => {
  http.post("/init", data.form,{validateStatus: function (status) {
    return true; 
  }}).then(r => {
    if (r.status === 200) {
    createToast("Initialized successfully")
      localStorage.setItem("Authorization", r.data.data.token)
      router.back()
    } else {
        createToast("Init failed:"+r.data.msg)
    }
  });
}
const countDownChanged =(dismissCountDown) => {
        this.dismissCountDown = dismissCountDown
    }
const showAlert = ()=> {
    this.dismissCountDown = this.dismissSecs
}
</script>
<template>
  <div class="d-flex justify-content-center align-items-center">
    <div class="shadow-sm p-3 mb-5 bg-body rounded">
      <form>
        <div class="mb-3">
          <label class="form-label">Server</label>
          <input type="text" class="form-control" v-model="data.form.Server">
        </div>
        <div class="mb-3">
          <label class="form-label">DBName</label>
          <input type="text" class="form-control" v-model="data.form.DBName">
        </div>
        <div class="mb-3">
          <label class="form-label">MySQL Username</label>
          <input type="text" class="form-control" v-model="data.form.User">
        </div>
        <div class="mb-3">
          <label class="form-label">MySQL Password</label>
          <input type="password" class="form-control" v-model="data.form.Password">
        </div>
        <div class="mb-3">
          <label class="form-label">Supervisor Username</label>
          <input type="text" class="form-control" v-model="data.form.Username">
        </div>
        <div class="mb-3">
          <label class="form-label">Supervisor Password</label>
          <input type="password" class="form-control" v-model="data.form.Pwd">
        </div>
        <div class="mb-3">
          <label class="form-label">Full Name</label>
          <input type="text" class="form-control" v-model="data.form.Full_name">
        </div>
        <div class="mb-3">
          <label class="form-label">Email</label>
          <input type="text" class="form-control" v-model="data.form.Email">
        </div>
        <div class="mb-3">
          <label class="form-label">Phone Number</label>
          <input type="text" class="form-control" v-model="data.form.Phone_number">
        </div>
        <div class="d-grid gap-2">
          <button type="button" class="btn btn-primary d-grid gap-2" @click="init">Set up the system!</button>
        </div>
      </form>
    </div>
  </div>
</template>