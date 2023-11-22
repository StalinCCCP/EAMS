<template>
  <!-- <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div> -->
</template>

<script setup>
import router from '@/router';
import { onMounted } from 'vue';

// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
onMounted(()=>{
  router.push('/hardware')
})
// export default {
//   name: 'HomeView',
//   components: {
//     HelloWorld
//   }

</script>
