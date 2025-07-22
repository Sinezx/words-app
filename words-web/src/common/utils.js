import { ElMessage } from 'element-plus'

function isBlank(str){
    return str == null || str == undefined || str.trim() == '';
}

function Message(str, msType){
    ElMessage({
        message: str,
        type: msType,
    });
}

export {isBlank, Message}