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
                        v-loading="loading"
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
        <el-drawer
                title="我是标题"
                :visible.sync="drawer"
                :direction="direction"
                :before-close="handleClose">
            <span>我来啦!</span>
        </el-drawer>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                drawer: false,
                //direction: 'rtl',//从右往左开
                //direction: 'ltr',//从左往右开
                //direction: 'ttb',//从上往下开
                direction: 'btt',//从下往上开
                loading: false,
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
                this.loading = true;
                window.backend.sub(this.formInline.project, this.formInline.kssj, this.formInline.jssj, this.formInline.czr).then(result => {
                    var arry = eval(result);
                    this.tableData = [];
                    this.tableData = arry;
                    this.multipleSelection=[];
                    this.loading = false;
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
                this.drawer = true

            },
            handleClose(done) {
                this.$confirm('确认关闭？')
                    .then(_ => {
                        done();
                    })
                    .catch(_ => {});
            }
        }
    }
</script>