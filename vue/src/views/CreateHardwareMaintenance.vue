<template>
    <n-form>
      <n-form-item label="HardwareID">
      <n-input v-model:value="req.HardwareID" />
      </n-form-item>
      <n-form-item label="IssueDescription">
      <n-input v-model:value="req.IssueDescription" />
      </n-form-item>
      <n-form-item label="SolutionDescription">
      <n-input v-model:value="req.SolutionDescription" />
      </n-form-item>
      <n-form-item label="Status">
      <n-select v-model:value="req.Status" placeholder="Select" :options="statusOpt"/>
      </n-form-item>
      <n-form-item label="MaintenanceDate" >
      <n-input v-model:value="req.MaintenanceDate" />
      </n-form-item>
      <n-form-item label="Cost" >
      <n-input v-model:value="req.Cost" />
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
    HardwareID:'',
    IssueDescription:'',
    SolutionDescription:'',
    MaintenanceDate:'',
    Cost:0,
    Status:''
})
const statusOpt=['保留', '正常', '占用', '非正常'].map(
        (v) => ({
          label: v,
          value: v
        })
      )
const put = () => {
  http
    .put("/p/admin/hmc", req.value, {
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
