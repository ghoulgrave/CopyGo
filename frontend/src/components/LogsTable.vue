<template>
    <div style="margin-top: 20px">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
            <el-form-item label="项目信息">
                <el-select v-model="formInline.project" placeholder="项目信息">
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
        <!-- <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button> -->
        <el-button @click="datas()">come on baby ！</el-button>
                <el-table
                        ref="multipleTable"
                        :data="tableData"
                        tooltip-effect="dark"
                        style="width: 100%"
                        height="600"
                        @selection-change="handleSelectionChange">
                    <el-table-column
                            type="selection"
                            width="55">
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="姓名"
                            width="120">
                    </el-table-column>
                    <el-table-column
                            prop="version"
                            label="提交版本号"
                            width="120">
                    </el-table-column>
                    <el-table-column
                            label="日期"
                            width="300">
                        <template slot-scope="scope">{{ scope.row.time }}</template>
                    </el-table-column>
                    <el-table-column
                            prop="path"
                            label="文件地址"
                            show-overflow-tooltip>
                    </el-table-column>
                </el-table>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                formInline: {
                    user: '',
                    project: '',
                    kssj: '2020-03-18 08:50:39',
                    jssj: '2020-03-22 08:50:39',
                    czr: ''
                },
                options: [],
                tableData: [],
                multipleSelection: []
            }
        },
        created() {
            this.projectName();
        },
        methods: {
            //获取项目列表
            projectName() {
                window.backend.projectName().then(result => {
                    this.options = eval(result);
                    this.formInline.project = this.options[0].value;
                });
            },
            //提交查询信息填写列表数据
            onSubmit() {
                window.backend.sub(this.formInline.project, this.formInline.kssj, this.formInline.jssj, this.formInline.czr).then(result => {
                    var arry = eval(result);
                    this.tableData = [];
                    this.tableData = arry;
                    this.multipleSelection=[];
                });
            },
            //复选
            toggleSelection(rows) {
                if (rows) {
                    rows.forEach(row => {
                        this.$refs.multipleTable.toggleRowSelection(row);
                    });
                } else {
                    this.$refs.multipleTable.clearSelection();
                }
            },
            //单选
            handleSelectionChange(val) {
                this.multipleSelection = val;
            },
            datas: function () {
                var self = this;
                window.backend.getPaths().then(result => {
                    var arry = eval(result)
                    // arry = eval("[{\"name\":\"wangzijie\",\"time\":\"2020-03-18 11:25:27 +0800 (三, 18  3 2020)\",\"version\":\"r250954\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/service/impl/ExternalInfServiceImpl.java\",\"sublogs\":\"\"},{\"name\":\"wangzijie\",\"time\":\"2020-03-18 11:25:27 +0800 (三, 18  3 2020)\",\"version\":\"r250954\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/service/impl/TurnProjectBgdjServiceImpl.java\",\"sublogs\":\"\"},{\"name\":\"wangzijie\",\"time\":\"2020-03-18 11:25:27 +0800 (三, 18  3 2020)\",\"version\":\"r250954\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/web/query/JzReadHtxxController.java\",\"sublogs\":\"\"},{\"name\":\"wangzijie\",\"time\":\"2020-03-18 11:25:27 +0800 (三, 18  3 2020)\",\"version\":\"r250954\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/web/WEB-INF/views/query/jzReadHtxx.ftl\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/core/service/impl/BdcBankServiceImpl.java\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/web/main/WfProjectController.java\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/java/cn/gtmap/bdcdj/web/query/BdcZhInfoController.java\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/resources/META-INF/conf/bdcdj/application.properties\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/resources/conf/bdcdj-mybatis/BdcZhInfoQuery.xml\",\"sublogs\":\"\"},{\"name\":\"shaoliyao\",\"time\":\"2020-03-18 12:50:30 +0800 (三, 18  3 2020)\",\"version\":\"r250969\",\"path\":\" /bdcdj/branches/bdcdj_dbqy/src/main/resources/conf/spring/bdcdj-config-cas.xml\",\"sublogs\":\"\"}]");
                    self.tableData = [];
                    self.tableData = arry;
                    // for(var i =0 ;i<arry.length ;i++){
                    //     self.tableData[i]=arry[i];
                    // }

                });
            }
        }
    }
</script>