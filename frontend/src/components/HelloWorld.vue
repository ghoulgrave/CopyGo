<template>
  <div class="container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="审批人">
        <el-input v-model="formInline.user" placeholder="审批人"></el-input>
      </el-form-item>
      <el-form-item label="活动区域">
        <el-select v-model="formInline.region" placeholder="活动区域">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
    <h1>{{message}}</h1>
    <a @click="getMessage">Press Me!</a>
  </div>
</template>

<script>
export default {
  data() {
    var ff ="";


    return {
      message: ff,
      formInline: {
        user: '',
        region: '选项3'
      },
    options: [],
    };
  },
  created() {
    this.logselect();
  },
  methods: {
    getMessage: function() {
      var self = this;
      window.backend.basic().then(result => {
        self.message = result;
      });
    },
    onSubmit() {
      // eslint-disable-next-line no-console
      console.log(this.region.valueOf())

    },
    logselect(){
      window.backend.logselect().then(result =>{
        this.options = eval(result);
      });
    }


  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
