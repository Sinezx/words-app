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

function getWordAudio(word){
    return axios({
        url: "/dictvoice?audio=" + word + "&type=2",
        method: 'get',
        responseType: 'blob'
    })
}

function uploadWordVoice(formData){
    return axios({
        url: "/api/v1/word/uploadwordvoice",
        method: 'post',
        headers: {
            "Content-Type": "multipart/form-data"
        },
        data: formData
    })
}

export {getUserInfo, queryWord, updateWord, addWord, getWordAudio, uploadWordVoice}