<script setup>
import { onMounted } from 'vue';
import {useRoute} from 'vue-router'
import { createToast } from 'mosha-vue-toastify';
const route=useRoute()
const HardwareID=route.params?.HardwareID
onMounted(()=>{
    const HardwareID=route.params?.HardwareID
    if(HardwareID){
        console.log(HardwareID)
    }else{
        createToast("No selected HardwareID!")
    }
})
</script>