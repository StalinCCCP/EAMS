<template>
  <n-space vertical :size="12">
    <n-form class="centerform" ref="formRef" inline :label-width="80" :model="req" :size="size">
      <n-form-item>
        <n-input size="large" v-model:value="req.HardwareName" placeholder="Find..." />
      </n-form-item>
      <n-form-item>
        <n-button attr-type="button" @click="find"> Go! </n-button>
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
import { h, defineComponent, onMounted } from "vue";
import {
  NButton,
  useMessage,
  NDataTable,
  NSpace,
  NForm,
  NFormItem,
  NInput,
} from "naive-ui";
// import {DataTableColumns} from 'naive-ui'
import http from "@/util/http";
import { createToast } from "mosha-vue-toastify";
import createColumns from "@/models/HardwareTable";
import { ref } from "vue";
const req = ref({
  HardwareName: "",
  Category: "",
  Status: "",
  Location: "",
});
const data = ref([]);
const formRef = ref(null);
const size = ref("medium");

const post = () => {
  http
    .post("/p/user/hlq", req.value, {
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
      } else {
        createToast("Failed to fetch data: " + r.data.msg);
      }
    });
};
onMounted(post);
const find = (e) => {
  e.preventDefault();
  post();
};
const columns = createColumns();
const pagination = {
  pagiSize: 20,
};
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
