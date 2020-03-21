<template>
  <div id="app">
    <HelloWorld/>
    <LogsTable/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import LogsTable from "./components/LogsTable.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {
    HelloWorld,
    LogsTable
  }
};
</script>
