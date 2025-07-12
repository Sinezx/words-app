<script setup>
import { Plus } from '@element-plus/icons-vue'
import { ElMessageBox } from 'element-plus'
import { ref } from 'vue';
import { queryWord, updateWord, addWord } from '@/api/word';
import { FIFOList } from '@/common/FIFOList';

var wordBuf = new FIFOList()

var word = ref({
    id:0,
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
    word.value.id = targetWord.id
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
    var id = word.value.id
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
    addWord(addWordObj.value.sourceText, addWordObj.value.targetText).then((resp)=>{
        console.log(resp.data.id)
    })
    addWordObj.value.sourceText = ""
    addWordObj.value.targetText = ""
    dialogVisible.value = false
}

function addWordCancel(){
    addWordObj.value.sourceText = ""
    addWordObj.value.targetText = ""
    dialogVisible.value = false
}

// init word buffer
refreshWordBuf()

</script>

<template>
  <div id="word-card">
    <div id="opt-head">
        <el-button type="primary" :icon="Plus" circle @click="dialogVisible = true"></el-button>
    </div>
    <div id="word-info">
        <div id="word-source">
            {{ word.sourceText }}
        </div>
        <div id="word-target">
            {{ word.targetText }}
        </div>
    </div>
    <div id="button-containor">
        <el-button type="success" round @click="knowClick">
            know
        </el-button>
        <el-button type="danger" round @click="unknowClick">
            unknow
        </el-button>
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
                <el-input v-model="addWordObj.sourceText"></el-input>
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
#word-card{
    margin-left: 40%;
    margin-right: 40%;
    min-width: 20%;
    margin-top: 10%;
    margin-bottom: 10%;
    min-height: 80%;
    border-radius: 20px;
    background-color: rgb(19, 71, 134);
    padding: 1%;
    box-shadow: var(--el-box-shadow-dark)
}

#opt-head{
    text-align: right;
    margin-top: 2%;
    margin-bottom: 2%;
}

#word-info{
    color: white;
    text-align: center;
    margin-top: 2%;
    margin-bottom: 2%;
}

#button-containor{
    text-align: center;
    margin-top: 2%;
    margin-bottom: 2%;
}
</style>
