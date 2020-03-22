<template>
    <div style="margin-top: 20px">
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
        <!-- <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button> -->
        <el-button @click="datas()">come on baby ！</el-button>
    <el-table
            ref="multipleTable"
            :data="tableData"
            tooltip-effect="dark"
            style="width: 100%"
            @selection-change="handleSelectionChange">
        <el-table-column
                type="selection"
                width="55">
        </el-table-column>
        <el-table-column
                label="日期"
                width="120">
            <template slot-scope="scope">{{ scope.row.date }}</template>
        </el-table-column>
        <el-table-column
                prop="name"
                label="姓名"
                width="120">
        </el-table-column>
        <el-table-column
                prop="address"
                label="地址"
                show-overflow-tooltip>
        </el-table-column>
    </el-table>


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
                tableData: [{
                    date: '2016-05-03',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-02',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-04',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-01',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-08',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-06',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-07',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }],
                multipleSelection: []
            }
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
            },
            toggleSelection(rows) {
                if (rows) {
                    rows.forEach(row => {
                        this.$refs.multipleTable.toggleRowSelection(row);
                    });
                } else {
                    this.$refs.multipleTable.clearSelection();
                }
            },
            handleSelectionChange(val) {
                this.multipleSelection = val;
            },
            datas: function () {
                var self = this;
                window.backend.datas().then(result => {
                    var arry  = eval(result);
                    self.tableData =[];
                    self.tableData =arry;
                    // for(var i =0 ;i<arry.length ;i++){
                    //     self.tableData[i]=arry[i];
                    // }

                });
            }
        }
    }
</script>