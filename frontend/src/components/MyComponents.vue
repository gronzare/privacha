<script setup>
import { reactive, ref } from 'vue'
import { EncryptText, ScanNewMsg, SetPassword } from '../../wailsjs/go/main/App'
import { LogDebug } from '../../wailsjs/runtime/runtime';
import { ElMessage } from 'element-plus'
import 'element-plus/dist/index.css'

function getDate(timestamp){
  let n=new Date(timestamp)
  return n.toLocaleDateString().replace(/\//g, "-") + " " + n.toTimeString().substr(0, 8)
}

const msgList = reactive({
    "mlist": []
})

function UpdateMsgList() {
    ScanNewMsg().then(result => {
        LogDebug(JSON.stringify(result))
        if (result == null) {
            return
        }
        for (let jsonStr of result) {
            try{
                let msg = JSON.parse(jsonStr)
                msgList.mlist.push(msg)
            }catch(e){
                ElMessage('无法解析的二维码，请检查Key')
            }
            
        }
    })
}

function init() {
    setInterval(UpdateMsgList, 1000)
}
init()

const textInput = ref("")
const configDialogVisible = ref(true)
const aboutDialogVisible = ref(false)
const configForm = reactive({
    Key:"defaultpasswordpleasechange",
    Nickname:"a"
})

const clickEncrypt = () => {
    let msg = { "Sender": configForm.Nickname, "Time": Date.now(), "Content": textInput.value }
    EncryptText(JSON.stringify(msg))
    ElMessage({
        message: '已复制到剪贴板',
        type: 'success',
    })
}
const clickConfigConfirm = () => {
    configDialogVisible.value=false
    SetPassword(configForm.Key)
}

</script>

<template>
    <el-dialog v-model="configDialogVisible" title="设置" width="100%">
        <el-form :model="configForm" style="max-width: 360px">
            <el-form-item label="KEY">
                <el-input v-model="configForm.Key" />
            </el-form-item>
            <el-form-item label="您的昵称">
                <el-input v-model="configForm.Nickname" />
            </el-form-item>
        </el-form>
        <el-button @click="clickConfigConfirm">确定</el-button>
    </el-dialog>
    <el-dialog v-model="aboutDialogVisible" title="使用说明" width="100%">
        功能说明：<br/>
        (0)设置密钥(KEY)<br/>
        打开设置菜单，设置一个KEY，软件之后将使用该KEY进行AES加解密。您可以把您的KEY以及该软件分享给您的朋友。<br/>
        (1)加密<br/>
        点击加密按钮，输入的文本将会按照您设置的KEY加密并转换为二维码，并自动复制到剪贴板。（请勿超过200字，二维码太复杂会难以识别）<br/>
        (2)解密<br/>
        实时识别屏幕上显示的二维码，使用您设置的KEY解码，并转换为消息。（请让二维码完整地显示在屏幕上）<br/>
    </el-dialog>

    <div id="top">
        <el-button size="small" @click="configDialogVisible=true">设置</el-button>
        <el-button size="small" @click="aboutDialogVisible=true">使用说明</el-button>
    </div>

    <div class="msg-box-area">
        <msg-box v-for="msg in msgList.mlist">
            <div class="msg-box">
                <div class="msg-box-head"><span class="msg-box-head-sender">{{ msg.Sender }}</span> <span
                        class="msg-box-head-time">{{ getDate(msg.Time) }}</span></div>
                <div class="msg-box-content">{{ msg.Content }}</div>
            </div>
        </msg-box>
    </div>
    <!-- <button @click="UpdateMsgList">Update</button> -->
    <div class="text-input-bar">
        <el-input v-model="textInput" :rows="4" type="textarea" placeholder="请输入要加密的文本" />
        <el-button type="primary" @click="clickEncrypt">加密</el-button>
    </div>

</template>

<style scoped>
#top {
    background-color: rgb(197, 197, 197);
    height: 2em;
}

.msg-box-area {
    height: calc(100% - 11em);
    overflow: auto;
}

.text-input-bar {
    height: 9em;
    background-color: rgb(202, 202, 202);
}

.msg-box-head-sender {
    color: rgb(93, 180, 52);
}

.msg-box-head-time {
    color: rgb(138, 216, 74);
}
</style>