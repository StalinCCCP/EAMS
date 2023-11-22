<template>
    <n-form>
      <n-form-item label="Software Name">
        <n-input :disabled="!isadmin" v-model:value="SoftwareName" type="text" placeholder="Input Software Name" />
      </n-form-item>
      <n-form-item label="Category">
        <n-input :disabled="!isadmin" v-model:value="Category" type="text" placeholder="Input Category" />
      </n-form-item>
      <n-form-item>
      <n-input
      :disabled="!isadmin"
        v-model:value="Description"
        type="textarea"
        placeholder="Input Description"
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
      <n-form-item label="Location">
        <n-input :disabled="!isadmin" v-model:value="Location" type="text" placeholder="Input Location" />
      </n-form-item>
      <div style="display: flex; justify-content: flex-end">
        <n-button round type="primary" @click="post">
          Update
        </n-button>
      </div>
    </n-form>
      <n-space vertical :size="12">
      <n-data-table
        :bordered="false"
        :single-line="false"
        :columns="columns"
        :data="data.Maintinfo"
      />
    </n-space>
  
  
  
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
  import createColumns from '@/models/SoftwareDetailTable';
  import router from '@/router/index'
  const SoftwareName=ref()
  const Description=ref()
  const Version=ref()
  const Status=ref()
  const Location=ref()
  const route=useRoute()
  const { query:{SoftwareID} } = route
  const req={
      SoftwareID:+SoftwareID
  }
  const statusOpt=['保留', '正常', '占用', '非正常'].map(
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
      .post("/p/user/sdq", req, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          data.value = r.data.data;
          SoftwareName.value=data.value.Sinfo.SoftwareName
          Version.value=data.value.Sinfo.Version
          Description.value=data.value.Sinfo.Description
          Status.value=data.value.Sinfo.Status
          Location.value=data.value.Sinfo.Location
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
    data.value.Sinfo.SoftwareName=SoftwareName.value
    data.value.Sinfo.Version=Version.value
    data.value.Sinfo.Description=Description.value
    data.value.Sinfo.Status=Status.value
    data.value.Sinfo.Location=Location.value
    http
      .post("/p/admin/supd", data.value.Sinfo, {
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
      if(SoftwareID){
          console.log(SoftwareID)
      }else{
          createToast("No selected SoftwareID!")
      }
      postq()
      postadmin()
  })
  const columns=createColumns({click(row){
      const url = router.resolve({
          name:'Software Maintenance Detail',
          query:{
              MaintenanceProcessID:row.MaintenanceProcessID
          }
      })
      window.open(url.href,"_blank")
  }})
  </script>