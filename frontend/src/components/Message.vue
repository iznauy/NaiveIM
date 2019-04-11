<template>
  <el-card>
    <div slot="header" style="text-align: left; height: 5px">
      <span style="font-size: 10px">{{time}}</span>
    </div>
    <span>{{message}}</span>
  </el-card>
</template>

<script>
  export default {
    name: "Message",
    props: {
      Type: Number,
      User: String,
      TimeStamp: Number,
      Content: String
    },
    computed: {
      message: function () {
        if (this.Type === 1) { // someone login
          return this.User + "加入了聊天室";
        } else if (this.Type === 2) {
          return this.User + "离开了聊天室";
        } else if (this.Type === 4) {
          return this.User + ": " + this.Content
        }
        return "";
      },
      time() {
        return new Date(parseInt(this.TimeStamp) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
      }
    }
  }
</script>

<style scoped>

</style>
