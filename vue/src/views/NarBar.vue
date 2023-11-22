<template>
  <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions" />
</template>

<script setup>
import { h, ref } from "vue";
import{
  NMenu
} from 'naive-ui'
const activeKey = ref(null);
const menuOptions = [
  {
    label:()=>h(
      "a",
      {
        href: "",
        target: "",
        // rel: "noopenner noreferrer"
      },
      "且听风吟"
    ),
    key: "",
  },
  {
    label:()=>h(
      "a",
      {
        href: "",
        target: "",
        // rel: "noopenner noreferrer"
      },
      "且听风吟"
    ),
    key: "",
  },
  {
    label:()=>h(
      "a",
      {
        href: "",
        target: "",
        // rel: "noopenner noreferrer"
      },
      "且听风吟"
    ),
    key: "",
  }
]
</script>