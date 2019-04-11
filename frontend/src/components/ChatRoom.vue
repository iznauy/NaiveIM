<template>
  <div>
    <el-form :inline="true">
      <el-form-item label-width="0">
        <el-input v-model="text" size="medium"></el-input>
      </el-form-item>
      <el-form-item label-width="0">
        <el-button type="primary" @click="click" size="small">发送</el-button>
      </el-form-item>
    </el-form>
    <div style="width: 400px; margin: 10px auto 10px">
      <message v-for="(data, index) in messages" v-bind="data" :key="messages.length - index"></message>
    </div>
  </div>
</template>

<script>

  import Message from "@/components/Message"

  export default {
    name: "ChatRoom",
    components: {
      Message
    },
    data() {
      return {
        messages: [],
        ws: null,
        text: ""
      }
    },
    methods: {
      click() {
        this.ws.send(this.text);
        this.text = "";
      },
      init() {
        console.log(this.messages);
        this.ws = new WebSocket("ws://localhost:8080/room?name=" + this.$store.getters.name);

        this.ws.onmessage = e => {
          this.messages.unshift(JSON.parse(e.data))
        };



        this.ws.onclose = e => {
          this.$alert("服务器出现异常", "", {
            confirmButtonText: '确定'});
          this.$router.push("/")
        }
      }
    },
    created() {
      this.messages = []
      this.init()
    }
  }

</script>

<style scoped>

</style>
