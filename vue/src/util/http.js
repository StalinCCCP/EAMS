import axios from 'axios'
import router from "@/router/index.js";

export const http = axios.create({
    baseURL: 'http://localhost:8080'
})

http.interceptors.request.use(config => {
    if (localStorage.getItem('Authorization')) {
        config.headers.Authorization = localStorage.getItem('Authorization')
    }
    return config
}, error => Promise.reject(error))

http.interceptors.response.use(response => {
    return response
}, error => {
    if(error.response.status===401){
        router.push('/login').then(r=>r)
    }
})

http.interceptors.response.use(response => {
    if (response.data.msg === "Not initialized") {
        router.push('/init').then(r => r)
    }
    return response
}, error => Promise.reject(error))
export default http