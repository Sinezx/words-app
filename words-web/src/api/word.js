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

function addWord(sourceText, targetText){
    return axios({
        url: "/api/v1/word/addword",
        method: 'post',
        data:{
            "source_text": sourceText,
            "target_text": targetText
        }
    })
}

export {getUserInfo, queryWord, updateWord, addWord}