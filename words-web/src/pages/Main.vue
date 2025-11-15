<script setup>
import { Headset, Plus } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue';
import { queryWord, updateWord, addWord, getWordAudio } from '@/api/word';
import { FIFOList } from '@/common/FIFOList';
import { isBlank, Message} from '@/common/utils';

var wordBuf = new FIFOList()

var word = ref({
    id:0,
    uwid:0,
    sourceText:"",
    targetText:"",
    rate:0
})

const dialogVisible = ref(false)

var addWordObj = ref({
    sourceText: "",
    targetText: ""
})

const dialogClose = function(done){
  ElMessageBox.confirm('Are you sure to close this dialog?')
    .then(() => {
        addWordObj.value.sourceText = ""
        addWordObj.value.targetText = ""
        done()
    })
    .catch(() => {
      // catch error
    })
}

function refreshWordView(targetWord){
    if(targetWord == null || targetWord == undefined){
        return
    }
    word.value.id = targetWord.wordid
    word.value.uwid = targetWord.id
    word.value.sourceText = targetWord.source_text
    word.value.targetText = targetWord.target_text
    word.value.rate = targetWord.rate
}

async function refreshWordBuf(){
    if(wordBuf.isEmpty()){
        await queryWord(1, 5).then((resp)=>{
            console.log(resp.data)
            if(resp.data.total > 0){
                for(var i = 0; i < resp.data.total; i++){
                    wordBuf.add(resp.data.words[i])
                }
            }
        })
    }
    refreshWordView(wordBuf.get())
}

function knowClick(){
    var id = word.value.uwid
    //update current word's rate
    updateWord(id).then((resp)=>{
        console.log(resp.data.message)
    })
    refreshWordBuf()
}

function unknowClick(){
    refreshWordBuf()
}

function addWordSubmit(){
    if(isBlank(addWordObj.value.sourceText) || isBlank(addWordObj.value.targetText)){
        Message("input is blank", "error")
    }else{
        const formData = new FormData()
        formData.append('source_text', addWordObj.value.sourceText)
        formData.append('target_text', addWordObj.value.targetText)
        addWord(formData).then((resp) => {
            console.log(resp.data.id)
        }).finally(()=>{
            addWordObj.value.sourceText = ""
            addWordObj.value.targetText = ""
            dialogVisible.value = false
        })
    }
}

function addWordCancel(){
    addWordObj.value.sourceText = ""
    addWordObj.value.targetText = ""
    dialogVisible.value = false
}

const audio = ref(new Audio())

function playAudioClick(){

    getWordAudio(word.value.id).then((resp)=>{
        if(resp.headers.getContentType() == "audio/mpeg"){
            const audioBlob = new Blob([resp.data], {type: "audio/mpeg"})
            const audiourl = URL.createObjectURL(audioBlob)
            audio.value.src = audiourl
            audio.value.play().then(()=>{
                URL.revokeObjectURL(audiourl)
            })
        }else{
            const reader = new FileReader()
            reader.readAsText(resp.data, 'utf-8')
            reader.onload = ()=>{
                const {message} = JSON.parse(reader.result)
                ElMessage(message)
            }
        }
    })
}

// init word buffer
refreshWordBuf()

</script>

<template>
    <div class="page-container">
        <div class="word-card">
            <div class="opt-head">
                <el-button type="primary" :icon="Plus" circle @click="dialogVisible = true"></el-button>
            </div>
            <div class="word-info">
                <div>
                    {{ word.sourceText }}
                    <el-icon>
                        <Headset @click="playAudioClick"/>
                    </el-icon>
                </div>
                <div>
                    {{ word.targetText }}
                </div>
            </div>
            <div class="opt-bottom">
                <el-button class="know-button" type="success" round @click="knowClick">
                    know
                </el-button>
                <el-button class="unknow-button" type="danger" round @click="unknowClick">
                    unknow
                </el-button>
            </div>
        </div>
    </div>

    <!-- the dialog is used to send request that add word to server -->
    <el-dialog
        v-model="dialogVisible"
        title="add word"
        :before-close="dialogClose"
    >
        <el-form :model="addWordObj">
            <el-form-item label="source">
                <el-input v-model="addWordObj.sourceText" @input="value=>addWordObj.sourceText=value.trim().toLowerCase()"></el-input>
            </el-form-item>
            <el-form-item label="target">
                <el-input v-model="addWordObj.targetText"></el-input>
            </el-form-item>
        </el-form>
        <template #footer>
        <div class="dialog-footer">
            <el-button @click="addWordCancel">Cancel</el-button>
            <el-button type="primary" @click="addWordSubmit">Confirm</el-button>
        </div>
        </template>
    </el-dialog>
</template>

<style scoped>
.page-container{
    display: grid;
    width: 100%;
    height: 100%;
    grid-template-columns: 1fr 8fr 1fr;
    grid-template-rows: 1fr 8fr 1fr;
    touch-action: manipulation;
}

.word-card{
    padding: 4%;
    grid-column: 2 / 3;
    grid-row: 2 /3;
    display: grid;
    grid-template-areas:
        'opt-head'
        'word-info'
        'opt-bottom';
    grid-template-rows: 1fr 8fr 1fr;
    border-radius: 20px;
    background-color: rgb(19, 71, 134);
    box-shadow: var(--el-box-shadow-dark)
}

.opt-head{
    grid-area: opt-head;
    text-align: right;
}

.word-info{
    grid-area: word-info;
    color: white;
    text-align: center;
}

.opt-bottom{
    grid-area: opt-bottom;
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-areas: 'kb ukb';
    column-gap: 5%;
    text-align: center;
}

.know-button{
    grid-area: kb;
    height: 100%;
    width: 100%;
    margin: 0;
}

.unknow-button{
    grid-area: ukb;
    height: 100%;
    width: 100%;
    margin: 0;
}

</style>
