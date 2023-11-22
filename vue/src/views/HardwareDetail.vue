<template>
  <n-space vertical>
    <n-input v-model:value="HardwareName" type="text" placeholder="Input Hardware Name" />
    <n-input
      v-model:value="Description"
      type="textarea"
      placeholder="Input Description"
    />
  </n-space>
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
} from "naive-ui";
import {useRoute} from 'vue-router'
import { createToast } from 'mosha-vue-toastify';
import http from '@/util/http';
import {ref} from 'vue'
import createColumns from '@/models/HardwareDetailTable';
import router from '@/router/index'
const HardwareName=ref()
const Description=ref()

const route=useRoute()
const { query:{HardwareID} } = route
const req={
    HardwareID:+HardwareID
}
const data=ref([])
const post = () => {
  http
    .post("/p/user/hdq", req, {
      validateStatus: function (status) {
        return true;
      },
    })
    .then((r) => {
      if (r.status === 200) {
        data.value = r.data.data;
        HardwareName.value=data.value.Hinfo[0].HardwareName
        Description.value=data.value.Hinfo[0].Description

        if (data.value === null) {
          data.value = [];
        }

    }else{
        createToast("Failed to fetch data: "+r.data.msg)
    }
})
}
onMounted(()=>{
    //const HardwareID=route.params?.HardwareID
    if(HardwareID){
        console.log(HardwareID)
    }else{
        createToast("No selected HardwareID!")
    }
    post()
})
const columns=createColumns({click(row){
    const url = router.resolve({
        name:'Hardware Maintenance Detail',
        query:{
            MaintenanceProcessID:row.MaintenanceProcessID
        }
    })
    window.open(url.href,"_blank")
}})
</script>