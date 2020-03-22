<template>
    <div style="margin-top: 20px">
        <div>
dddd
        </div>
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
        methods: {
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