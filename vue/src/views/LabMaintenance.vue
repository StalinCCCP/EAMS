<template>
    <n-space vertical :size="12">
      <n-form class="centerform" ref="formRef" inline :label-width="80" :model="req" :size="size">
      <!-- <n-form-item>
          <n-select
              v-model:value="req.Category"
              placeholder="Select"
              :options="Category"
          />
      </n-form-item>
      <n-form-item>
          <n-select
              v-model:value="req.Status"
              placeholder="Select"
              :options="statusOpt"
          />
       </n-form-item>
       <n-form-item>
          <n-select
              v-model:value="req.Location"
              placeholder="Select"
              :options="Location"
          />
       </n-form-item> -->
        <n-form-item>
          <n-input size="large" v-model:value="req.LabName" placeholder="Find..." />
        </n-form-item>
        <n-form-item>
          <n-button attr-type="button" @click="find"> Go! </n-button>
        </n-form-item>
        <n-form-item>
          <n-button :disabled="!isadmin" attr-type="button" @click="create"> + </n-button>
        </n-form-item>
      </n-form>
    </n-space>
    <n-space vertical :size="12">
      <n-data-table
        :bordered="false"
        :single-line="false"
        :columns="columns"
        :data="data"
        :pagination="pagination"
      />
    </n-space>
  </template>
  <script setup>
  import { h, defineComponent, onBeforeMount, onMounted } from "vue";
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
  // import {DataTableColumns} from 'naive-ui'
  
  import http from '@/util/http';
  import { createToast } from 'mosha-vue-toastify';
  import createColumns from '@/models/LabMaintenanceTable'
  import {ref} from 'vue'
  import router from '@/router/index'
  const req=ref({
    LabName:'',
  })
//   const statusOpt=ref(['保留', '正常', '占用', '非正常'].map(
//           (v) => ({
//             label: v,
//             value: v
//           })
//         ))
//   statusOpt.value.push({label:'(Empty)',value:''})
  const data=ref([])
  const formRef = ref(null)
  const size =ref("medium")
  const isadmin=ref(false)
//   const Category=ref()
//   const Location=ref()
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
  const create = ()=>{
      router.push('/createlabmaintenance')
  }
  const post = () => {
    http
      .post("/p/user/lmlq", req.value, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          data.value = r.data.data;
          if (data.value === null) {
            data.value = [];
          }
  
      }else{
          createToast("Failed to fetch data: "+r.data.msg)
      }
  })
  }
//   const getCategory = () => {
//     http
//       .get("/p/user/hcq", {
//         validateStatus: function (status) {
//           return true;
//         },
//       })
//       .then((r) => {
//         if (r.status === 200) {
          
//           Category.value=r.data.data.filter(item=>item!==null).map(
//           (v) => ({
//             label: v,
//             value: v
//           })
//         )
//         Category.value.push({label:'(Empty)',value:''})
  
//           if (Category.value === null) {
//               Category.value = [];
//           }
  
//       }else{
//           createToast("Failed to fetch data: "+r.data.msg)
//       }
//   })
//   }
//   const getLocation = () => {
//     http
//       .get("/p/user/hlocq", {
//         validateStatus: function (status) {
//           return true;
//         },
//       })
//       .then((r) => {
//         if (r.status === 200) {
//           Location.value=r.data.data.filter(item=>item!==null).map(
//           (v) => ({
//             label: v,
//             value: v
//           })
//         )
//           Location.value.push({label:'(Empty)',value:''})
//           // if (Location.value === null) {
//           //     Location.value = [];
//           // }
  
//       }else{
//           createToast("Failed to fetch data: "+r.data.msg)
//       }
//   })
//   }
  onMounted(()=>{
      post()
      postadmin()
    //   getCategory()
    //   getLocation()
  })
  const find = (e)=>{
      e.preventDefault()
      post()
  }
  const columns=createColumns({click(row){
      const url = router.resolve({
          name:'Lab Maintenance Detail',
          query:{
              MaintenanceProcessID:row.MaintenanceProcessID
          }
      })
      window.open(url.href,"_blank")
  },del(row){
      http
      .delete("/p/admin/lmdlt", {data:{MaintenanceProcessID:row.MaintenanceProcessID}}, {
        validateStatus: function (status) {
          return true;
        },
      })
      .then((r) => {
        if (r.status === 200) {
          createToast("Delete Successfully ")
      }else{
          createToast("Failed to delete")
      }
  })
  }})
      //window.open(url.href,"_blank")
  
  const pagination={
      pagiSize:20,
  }
  
  // defineComponent({
  //         setup(){
  //             const message = useMessage()
  //             return {
  //                 data,
  //                 columns:createColumns(),
  //                 pagination:{
  //                     pageSize: 20
  //                 }
  //             }
  //         }
  //     })
  </script>
  
  <style lang="scss">
  .centerform {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  
  .n-input:not(.n-input--autosize) {
      width: 500px;
  }
  </style>
  