
<template>
    <n-form>
      <n-form-item label="LabID">
        <n-input :disabled="!isadmin" v-model:value="LabID" type="text" placeholder="Input Lab Name" />
      </n-form-item>
      <!-- <n-form-item label="Category">
        <n-input :disabled="!isadmin" v-model:value="Category" type="text" placeholder="Input Category" />
      </n-form-item> -->
      <n-form-item>
      <n-input
      :disabled="!isadmin"
        v-model:value="IssueDescription"
        type="textarea"
        placeholder="Input Issue Description"
      />
      </n-form-item>
      <n-form-item>
      <n-input
      :disabled="!isadmin"
        v-model:value="SolutionDescription"
        type="textarea"
        placeholder="Input Solution Description"
      />
      </n-form-item>
      <n-form-item label="Status">
        <n-select
        :disabled="!isadmin"
          v-model:value="Status"
          placeholder="Select"
          :options="statusOpt"
        />
      </n-form-item>
      <n-form-item label="Cost">
        <n-input :disabled="!isadmin" v-model:value="Cost" type="text" placeholder="Input Lab Name" />
      </n-form-item>
      <n-form-item label="MaintenanceDate">
        <n-input :disabled="!isadmin" v-model:value="MaintenanceDate" type="text" placeholder="Input Lab Name" />
      </n-form-item>
      <div style="display: flex; justify-content: flex-end">
        <n-button round type="primary" @click="post">
          Update
        </n-button>
      </div>
      </n-form>
    <!-- </n-form>
      <n-space vertical :size="12">
      <n-data-table
        :bordered="false"
        :single-line="false"
        :columns="columns"
        :data="data.Maintinfo"
      />
    </n-space> -->
  
  
  
  </template>
  <script setup>
  import { onMounted } from 'vue';
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
  import {useRoute} from 'vue-router'
  import { createToast } from 'mosha-vue-toastify';
  import http from '@/util/http';
  import {ref} from 'vue'
//   import createColumns from '@/models/SoftwareMaintenanceDetailTable';
//   import router from '@/router/index'
  const LabID=ref()
  const IssueDescription=ref()
  const SolutionDescription=ref()
  const MaintenanceDate=ref()
  const Status=ref()
  const Cost=ref()
  //const Location=ref()
  const route=useRoute()
  const { query:{MaintenanceProcessID} } = route
  const req={
    MaintenanceProcessID:+MaintenanceProcessID
  }
  const statusOpt=['已完成','待处理'].map(
          (v) => ({
            label: v,
            value: v
          })
        )
  const data=ref([])
  const isadmin=ref(false)
  const postadmin = ()=>{
    http
      .get("/p/admin/isadmin", {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          isadmin.value=true
        }
  })
  }
  const postq = () => {
    http
      .post("/p/user/lmdq", req, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          data.value = r.data.data;
          LabID.value=data.value.LabID
          IssueDescription.value=data.value.IssueDescription
          SolutionDescription.value=data.value.SolutionDescription
          MaintenanceDate.value=data.value.MaintenanceDate
          Status.value=data.value.Status
          Cost.value=data.value.Cost
          if (data.value === null) {
            data.value = [];
          }
  
      }else{
          data.value={}
          createToast("Failed to fetch data: "+r.data.msg)
      }
  })
  }
  const post= () => {
    data.value.LabID=LabID.value
    data.value.IssueDescription=IssueDescription.value
    data.value.SolutionDescription=SolutionDescription.value
    data.value.Status=Status.value
    data.value.Location=Cost.value
    data.value.MaintenanceDate=MaintenanceDate.value
    http
      .post("/p/admin/lmupd", data.value, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          createToast("Information updated!")
      }else{
          createToast("Failed to update data: "+r.data.msg)
      }
  })
  }
  onMounted(()=>{
      //const HardwareID=route.params?.HardwareID
      if(MaintenanceProcessID){
          console.log(MaintenanceProcessID)
      }else{
          createToast("No selected MaintenanceProcessID!")
      }
      postq()
      postadmin()
  })
//   const columns=createColumns()
  </script>
  
  <!-- <style lang="scss">
  .centerform {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .n-input:not(.n-input--autosize) {
      width: 500px;
  }
  .n-select:not(.n-select--autosize) {
      width: 500px;
  }
  </style> -->
  