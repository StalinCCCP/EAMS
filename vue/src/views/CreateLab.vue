

<template>
    <n-form>
      <n-form-item label="Lab Name">
      <n-input v-model:value="req.LabName" />
      </n-form-item>
      <!-- <n-form-item label="Category">
      <n-input v-model:value="req.Category" />
      </n-form-item> -->
      <n-form-item label="Description">
      <n-input v-model:value="req.Description" />
      </n-form-item>
      <n-form-item label="Status">
      <n-select v-model:value="req.Status" placeholder="Select" :options="statusOpt"/>
      </n-form-item>
      <n-form-item label="Location" >
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
  NSelect
} from "naive-ui";
import { createToast } from 'mosha-vue-toastify';

const req=ref({
    LabName:'',
    // Category:'',
    Description:'',
    Status:'',
    Location:''
})
const statusOpt=[ '正常', '占用', '非正常'].map(
        (v) => ({
          label: v,
          value: v
        })
      )
const put = () => {
  http
    .put("/p/admin/lc", req.value, {
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
