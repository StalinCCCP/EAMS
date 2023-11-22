import axios from 'axios'

function get() {

    // 定义 data 数据对象
    const data = {
        name: "wangwu",
        age: 30
    }
    // 格式化成 json 字符串
    const _data = JSON.stringify(data)

    // 使用 axios 发送请求
    resp = axios({
        method: "get",
        url: "http://127.0.0.1:8088/get/id12312312?money=300",
        data: _data,
    })
    // 输出结果， 成功获取
    // { ID: 'id12312312', Money: 300, Data: { name: 'wangwu', age: 30 } }
    //console.log(resp.data); 

}

// node main.js
get() 