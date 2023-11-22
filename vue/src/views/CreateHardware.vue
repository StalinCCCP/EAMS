

<template>
    <n-form>
      <n-form-item label="Hardware Name">
      <n-input v-model:value="req.HardwareName" />
      </n-form-item>
      <n-form-item label="Category">
      <n-input v-model:value="req.Category" />
      </n-form-item>
      <n-form-item label="Description">
      <n-input v-model:value="req.Description" />
      </n-form-item>
      <n-form-item label="Status">
      <n-input v-model:value="req.Status"  />
      </n-form-item>
      <n-form-item label="Status" >
      <n-input v-model:value="req.Location" />
      </n-form-item>
    </n-form>
    <n-form>
      <n-button type="primary" @click="put"> Submit </n-button>
    </n-form>
</template>
<script setup>
import { ref } from 'vue'
import http from '@/util/http';
import router from '@/router/index'
import {
  NButton,
  useMessage,
  NDataTable,
  NSpace,
  NForm,
  NFormItem,
  NInput,
} from "naive-ui";
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
        createToast("Successfully created!")
        router.back();
      }
    });
};

</script>
<style lang="scss"></style>
