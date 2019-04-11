<template>
  <div>
    <h2>OK</h2>
    <button @click="click">22333</button>
  </div>
</template>

<script>

  export default {
    name: "ChatRoom",
    data() {
      messages: []
      ws: null
    },
    methods: {
      click() {
        this.ws.send("ojbk")
      },
      init() {
        this.ws = new WebSocket("ws://localhost:8080/room?name=" + this.$store.getters.name);
        this.ws.onopen = function () {
          this.ws.send("ok");
          alert("数据发送中");
        };

        this.ws.onmessage = function (event) {
          let received_msg = event.data;
          alert(received_msg);
        };

        this.ws.onclose = function () {
          alert("关闭！")
        };
      }
    },
    created() {
      this.init()
    }
  }

</script>

<style scoped>

</style>
