<template>
  <div id="app" @contextmenu.prevent="rightClick()">

    <LogsTable/>
  </div>
</template>

<script>
import LogsTable from "./components/LogsTable.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  components: {

    LogsTable
  },
    method:{

    }
};
</script>
