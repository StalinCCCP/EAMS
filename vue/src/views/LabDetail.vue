<template>
    <n-form>
      <n-form-item label="Lab Name">
        <n-input :disabled="!isadmin" v-model:value="LabName" type="text" placeholder="Input Lab Name" />
      </n-form-item>
      <!-- <n-form-item label="Category">
        <n-input :disabled="!isadmin" v-model:value="Category" type="text" placeholder="Input Category" />
      </n-form-item> -->
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
  import createColumns from '@/models/LabDetailTable';
  import router from '@/router/index'
  const LabName=ref()
  const Description=ref()
//   const Category=ref()
  const Status=ref()
  const Location=ref()
  const route=useRoute()
  const { query:{LabID} } = route
  const req={
      LabeID:+LabID
  }
  const statusOpt=[ '正常', '占用', '非正常'].map(
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
      .post("/p/user/ldq", req, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          data.value = r.data.data;
          LabName.value=data.value.Linfo.LabName
        //   Category.value=data.value.Linfo.Category
          Description.value=data.value.Linfo.Description
          Status.value=data.value.Linfo.Status
          Location.value=data.value.Linfo.Location
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
    data.value.Linfo.LabName=LabName.value
    // data.value.Hinfo.Category=Category.value
    data.value.Linfo.Description=Description.value
    data.value.Linfo.Status=Status.value
    data.value.Linfo.Location=Location.value
    http
      .post("/p/admin/lupd", data.value.Linfo, {
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
      if(LabID){
          console.log(LabID)
      }else{
          createToast("No selected LabID!")
      }
      postq()
      postadmin()
  })
  const columns=createColumns({click(row){
      const url = router.resolve({
          name:'Lab Maintenance Detail',
          query:{
              MaintenanceProcessID:row.MaintenanceProcessID
          }
      })
      window.open(url.href,"_blank")
  }})
  </script>