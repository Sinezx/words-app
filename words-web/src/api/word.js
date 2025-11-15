import axios from "./request"

function getUserInfo(data){
    return axios({
        url: "/api/userinfo",
        method: 'post',
        data
    })
}

function queryWord(page, pageSize){
    return axios({
        url: "/api/v1/word/queryword",
        method: 'post',
        data: {
            "page": page, 
            "pagesize": pageSize
        }
    })
}

function updateWord(id){
    return axios({
        url: "/api/v1/word/updateword",
        method: 'post',
        data: {
            "id": id,
            "status": 0
        }
    })
}

function addWord(formData){
    return axios({
        url: "/api/v1/word/addword",
        method: 'post',
        headers: {
            "Content-Type": "multipart/form-data"
        },
        data: formData
    })
}

function getWordAudio(wordId){
    return axios({
        url: "/api/v1/word/getwordvoice/" + wordId,
        method: 'get',
        responseType: 'blob'
    })
}

export {getUserInfo, queryWord, updateWord, addWord, getWordAudio}