

<template>
  <n-space vertical :size="12">
    <n-form>
      <n-input v-model:value="req.HardwareName" label="Hardware Name" />
      <n-input v-model:value="req.Category" label="Category" />
      <n-input v-model:value="req.Description" label="Description" />
      <n-input v-model:value="req.Status" label="Status" />
      <n-input v-model:value="req.Location" label="Status" />
    </n-form>
    <n-form>
      <n-button type="primary" @click="put"> Submit </n-button>
    </n-form>
  </n-space>
</template>
<script setup>
import { ref } from 'vue'
import http from '@/util/http';
import router from '@/router/index'
const req=ref({
    HardwareName:'',
    Category:'',
    Description:'',
    Status:'',
    Location:''
})
const put = () => {
  http
    .put("/p/admin/hc", req.value, {
      validateStatus: function (status) {
        return true;
      },
    })
    .then((r) => {
      if (r.status === 200) {
        console.log('成功!'+r.data)
        router.back();
      }
    });
};

</script>
<style lang="scss"></style>
