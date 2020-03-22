<template>
  <div class="container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="项目信息">
        <el-select v-model="formInline.region" placeholder="项目信息">
          <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="开始时间">
        <el-input v-model="formInline.kssj" placeholder="开始时间"></el-input>
      </el-form-item>
      <el-form-item label="结束时间">
        <el-input v-model="formInline.jssj" placeholder="结束时间"></el-input>
      </el-form-item>
      <el-form-item label="操作人">
        <el-input v-model="formInline.czr" placeholder="操作人"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: "hello",
      formInline: {
        user: '',
        region: '',
        kssj:'',
        jssj:'',
        czr:''
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
      window.backend.sub(this.formInline.czr).then(result =>{
        this.message = result;
      });
    },
    logselect(){
      window.backend.logselect().then(result =>{
        this.options = eval(result);
        this.formInline.region = this.options[0].value;
      });
    }


  }
};
</script>
