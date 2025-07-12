import axios from "./request"

function register(account, password){
    return axios({
        url: "/api/v1/reg",
        method: 'post',
        data: {
            "account": account,
            "password": password
        }
    })
}

function login(account, password){
    return axios({
        url: "/api/v1/sayhi",
        method: 'post',
        data: {
            "account": account,
            "password": password
        }
    })
}

export {register, login}