<template>

    <el-container>
        <el-header>
            <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
                <el-menu-item index="1">单个项目处理</el-menu-item>
                <el-menu-item index="10">批量项目处理</el-menu-item>
                <el-menu-item index="3-1">项目配置</el-menu-item>
                <el-menu-item index="3-2">系统配置</el-menu-item>
                <!--                <el-submenu index="3">-->
                <!--                    <el-menu-item index="3-1">项目配置</el-menu-item>-->
                <!--                    <el-menu-item index="3-2">系统配置</el-menu-item>-->
                <!--                    <el-menu-item index="3-3">选项3</el-menu-item>-->
                <!--                    <el-submenu index="3-4">-->
                <!--                        <template slot="title">选项4</template>-->
                <!--                        <el-menu-item index="3-4-1">选项1</el-menu-item>-->
                <!--                        <el-menu-item index="3-4-2">选项2</el-menu-item>-->
                <!--                        <el-menu-item index="3-4-3">选项3</el-menu-item>-->
                <!--                    </el-submenu>-->
                <!--                </el-submenu>-->
                <el-menu-item index="4">帮助中心</el-menu-item>
                <el-menu-item index="5">联系反馈</el-menu-item>
            </el-menu>
        </el-header>
        <el-main style="margin-top: 1px;">
            <div v-if="tab1" style="margin-top: 2px">
                <el-form :inline="true" :model="formInline" class="demo-form-inline">
                    <el-form-item label="项目信息">
                        <el-select v-model="formInline.project" placeholder="项目信息">
                            <el-option v-for="item in options" :key="item.value" :label="item.label"
                                       :value="item.value">
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
                    <el-form-item>
                        <el-button type="primary" @click="datas()">打包更新文件</el-button>
                    </el-form-item>
                </el-form>
                <!-- <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button> -->

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
                            width="350">
                        <template slot-scope="scope">{{ scope.row.time }}</template>
                    </el-table-column>
                    <el-table-column
                            prop="path"
                            label="文件地址"
                            show-overflow-tooltip>
                    </el-table-column>
                </el-table>

                <!--                <el-button @click="">打包更新文件（包含jar包）</el-button>-->
                <!--                <el-button @click="">整包获取</el-button>-->
                <el-drawer
                        title="编译情况"
                        :visible.sync="drawer"
                        :direction="direction"
                        :before-close="handleClose" size="99%">
                    <div style="height: 800px ">
                        <!-- 注意需要给 el-scrollbar 设置高度，判断是否滚动是看它的height判断的 -->
                        <el-scrollbar ref="myScrollbar" style="height: 100%;"> <!-- 滚动条 -->
                            <p style="white-space: pre-line;color: #303133;text-align: left;">{{ message }}</p>
                        </el-scrollbar><!-- /滚动条 -->
                    </div>
                </el-drawer>
            </div>
            <div v-if="tab10" style="margin-top: 2px">
                <p style="white-space: pre-line;color: #303133;text-align: left;">批量处理</p>
                <el-form ref="form" label-width="80px">
                    <el-form-item label="开始时间">
                        <el-input v-model="kssj_pl" placeholder="开始时间"></el-input>
                    </el-form-item>
                    <el-form-item label="结束时间">
                        <el-input v-model="jssj_pl" placeholder="结束时间"></el-input>
                    </el-form-item>
                    <el-form-item label="操作人">
                        <el-input v-model="czr_pl" placeholder="操作人"></el-input>
                    </el-form-item>
                    <el-form-item label="是否编译">
                        <el-switch
                                v-model="buildOrNot"
                                active-text="编译"
                                inactive-text="不编译">
                        </el-switch>
                    </el-form-item>
                    <el-form-item label="项目选择">
                        <el-checkbox-group v-model="checks" class="checkGroup" size="900">
                            <el-checkbox v-for="(item) in checkboxs" :label="item.label" :key="item.value" border>{{item.value}}
                            </el-checkbox>
                        </el-checkbox-group>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="Alldatas()">批量打包文件</el-button>
                    </el-form-item>
                </el-form>
                <el-drawer
                        title="批量编译情况"
                        :visible.sync="drawer_pl"
                        :direction="direction_pl"
                        :before-close="handleClose" size="99%">
                    <div style="height: 800px ">
                        <!-- 注意需要给 el-scrollbar 设置高度，判断是否滚动是看它的height判断的 -->
                        <el-scrollbar ref="myScrollbar_pl" style="height: 100%;"> <!-- 滚动条 -->
                            <p style="white-space: pre-line;color: #303133;text-align: left;">{{ message_pl }}</p>
                        </el-scrollbar><!-- /滚动条 -->
                    </div>
                </el-drawer>
            </div>
            <div v-if="tab3_1" style="margin-top: 2px">
                <p style="white-space: pre-line;color: #303133;text-align: left;">项目信息</p>
                <el-button type="primary" @click="projectAdd()">新增空配置</el-button>
                <el-carousel trigger="click" height="500px" :interval="4000000" arrow="always" style="margin-top: 20px" ref="carousel" >
                    <el-carousel-item  v-for="(item,i) in projectConfs" :key="i" :name="i">

                        <el-form label-width="180px">
                            <el-input type="text" v-model="projectConfs[i].uid" hidden></el-input>
                            <el-form-item label="项目名称">
                                <el-input type="text" v-model="projectConfs[i].name"></el-input>
                            </el-form-item>
                            <el-form-item label="svn地址">
                                <el-input type="text" v-model="projectConfs[i].svnpath"></el-input>
                            </el-form-item>
                            <el-form-item label="更新提交地址">
                                <el-input type="text" v-model="projectConfs[i].subpath"></el-input>
                            </el-form-item>
                            <el-form-item label="更新输出地址">
                                <el-input type="text" v-model="projectConfs[i].outpath"></el-input>
                            </el-form-item>
                            <el-form-item label="工程物理文件夹">
                                <el-input type="text" v-model="projectConfs[i].dirpath"></el-input>
                            </el-form-item>

                            <el-button type="primary" @click="projectSubmit(projectConfs[i],i)">保存</el-button>
                            <el-button type="primary" @click="projectCopy(projectConfs[i],i)">复制</el-button>
                            <el-button type="primary" @click="projectDel(projectConfs[i],i)">删除</el-button>

                        </el-form>
                    </el-carousel-item>
                </el-carousel>

            </div>
            <div v-if="tab3_2" style="margin-top: 2px">
                <p style="white-space: pre-line;color: #303133;text-align: left;">系统信息</p>
                <el-form ref="form" :model="sysform" label-width="180px">
                    <el-form-item label="用户名称中文">
                        <el-input v-model="sysform.cname"></el-input>
                    </el-form-item>
                    <el-form-item label="用户名称英文">
                        <el-input v-model="sysform.ename"></el-input>
                    </el-form-item>
                    <el-form-item label="批量输出文件夹位置">
                        <el-input v-model="sysform.plOuPath"></el-input>
                    </el-form-item>
                    <el-form-item label="jar包通用名称">
                        <el-input
                                type="textarea"
                                :rows="10"
                                placeholder="jar包名称，去掉版本号部分。"
                                v-model="sysform.textarea">
                        </el-input>
                    </el-form-item>
                    <el-button type="primary" @click="sysOnSubmit">保存信息</el-button>
                </el-form>


            </div>
            <div v-if="tab4" style="margin-top: 2px">
                <div style="height: 700px">
                    <el-scrollbar style="height: 100%;">
                        <p style="white-space: pre-line;color: #303133;text-align: left;">帮助中心</p>
                        <p style="white-space: pre-line;color: #303133;text-align: left;">
                            1.批量输出地址没有配置，则使用找到的第一个项目输出地址


                        </p>
                        <p style="white-space: pre-line;color: #303133;text-align: left;">已知问题</p>
                        <p style="white-space: pre-line;color: #303133;text-align: left;">
                            * 日志记录问题没有实现
                            * 整包没有实现
                            * 配置信息修改后没有进行参数刷新，需要重启
                            * 进入后默认帮助页面和单项页面切换
                        </p>
                    </el-scrollbar>
                </div>

            </div>
            <div v-if="tab5" style="margin-top: 2px">
                <p style="white-space: pre-line;color: #303133;text-align: left;">联系反馈</p>
                <p style="white-space: pre-line;color: #303133;text-align: left;">有问题请联系：zhangyiyang@gtmap.cn</p>
            </div>
        </el-main>
    </el-container>


</template>

<script>
    export default {
        data() {
            return {
                sysform: {
                    cname: '',
                    ename: '',
                    plOuPath: '',
                    textarea: ''
                },
                projectForm: {},
                checkboxs: [],
                checks: [],
                tab1: true,
                tab10: false,
                tab3_1: false,
                tab3_2: false,
                tab4: false,
                tab5: false,
                activeIndex: '1',
                message: '',
                drawer: false,
                //direction: 'rtl',//从右往左开
                //direction: 'ltr',//从左往右开
                //direction: 'ttb',//从上往下开
                direction: 'btt',//从下往上开
                message_pl: '',
                drawer_pl: false,
                direction_pl: 'ttb',
                loading: false,
                formInline: {
                    user: '',
                    project: '',
                    kssj: this.formartDate("min", "first"),
                    jssj: this.formartDate("min", "end"),
                    czr: ''
                },
                kssj_pl: this.formartDate("day", "first"),
                jssj_pl: this.formartDate("day", "end"),
                czr_pl: '',
                buildOrNot: true,
                options: [],
                tableData: [],
                multipleSelection: [],
                projectConfs: [],
                //新增配置或复制配置时的提示信息
                tips:''
            }
        },
        beforeCreate() {

        },
        created() {
            this.project()
            this.projectName()
            this.getEName()
            this.getCName()
            this.getPlOuPath()
            this.getJarNames()
        },
        updated: function () {
            this.scrollDown()
        },
        mounted: function () {
            //绑定事件可以使用了
            window.wails.Events.On('cpu_usage', cpu_usage => {
                if (cpu_usage) {
                    //console.log("sss"+cpu_usage.str);
                    this.message = this.message + "\n" + cpu_usage.str;
                }
            });
            //绑定事件可以使用了
            window.wails.Events.On('builds_pl', s => {
                if (s) {
                    //console.log("sss"+cpu_usage.str);
                    this.message_pl = this.message_pl + "\n" + s.str;
                }
            });
        },
        methods: {
            projectCopy(project,i){
                var item = {
                    uid: 'project-'+new Date().getTime(),
                    name: project.name+new Date().getTime(),
                    svnpath: project.svnpath,
                    dirpath: project.dirpath,
                    subpath: project.subpath,
                    outpath: project.outpath
                };
                this.projectConfs.splice(0, 0, item)
                window.backend.ThisCopy.SaveOrUpdatePorjectConfig(JSON.stringify(item)).then(result => {
                    if (result == '操作成功') {
                        this.showMessages('操作成功！' + '请重启。', 'info')
                    } else {
                        this.showMessages(result, 'warning')
                    }
                });
            },
            projectDel(project,i){
                if(this.projectConfs.length <= 1){
                    this.showMessages("最后一个配置了，别删了。", 'warning')
                    return
                }
                this.projectConfs.splice(i, 1)
                window.backend.ThisCopy.DelProjectConfig(JSON.stringify(project)).then(result => {
                    if (result == '操作成功') {
                        this.showMessages('操作成功！' + '请重启。', 'info')
                    } else {
                        this.showMessages(result, 'warning')
                    }
                });

            },
            //新增一个新项目的配置
            projectAdd(){
                var item = {
                    uid: 'project-'+new Date().getTime(),
                    name: 'project-'+new Date().getTime(),
                    svnpath: '',
                    dirpath: '',
                    subpath: '',
                    outpath: ''
                };
                this.projectConfs.splice(0, 0, item)
                window.backend.ThisCopy.SaveOrUpdatePorjectConfig(JSON.stringify(item)).then(result => {
                    if (result == '操作成功') {
                        this.showMessages('操作成功！' + '请重启。', 'info')
                    } else {
                        this.showMessages(result, 'warning')
                    }
                });
            },
            projectSubmit(project,i) {
                window.backend.ThisCopy.SaveOrUpdatePorjectConfig(JSON.stringify(project)).then(result => {
                    if (result == '操作成功') {
                        this.showMessages('操作成功！' + '请重启。', 'info')
                    } else {
                        this.showMessages(result, 'warning')
                    }
                });
            },
            sysOnSubmit() {
                window.backend.ThisCopy.UpSysConfig(this.sysform.cname, this.sysform.ename, this.sysform.plOuPath, this.sysform.textarea).then(result => {
                    if (result == '操作成功') {
                        this.showMessages('操作成功！' + '请重启。', 'info')
                    } else {
                        this.showMessages("操作失败了，请重新操作", 'warning')
                    }
                });
            },
            //获取项目列表
            projectName() {
                window.backend.ThisCopy.GetProjectName().then(result => {
                    this.options = eval(result);
                    this.formInline.project = this.options[0].value;
                    this.checkboxs = this.options;
                });

            },
            //获取所有已经配置的项目
            project() {
                window.backend.ThisCopy.GetProjectConfs().then(result => {
                    if (result) {
                        var _array = eval(result);
                        _array.forEach((s, i) => {
                            this.projectConfs.push(s);
                        })
                    }
                });
            },
            //获取查询姓名信息
            getEName() {
                window.backend.ThisCopy.GetEName().then(result => {
                    this.sysform.ename = result;
                    this.formInline.czr = result;
                    this.czr_pl = result;
                });
            },
            //获取中文姓名信息
            getCName() {
                window.backend.ThisCopy.GetCName().then(result => {
                    this.sysform.cname = result;
                });
            },
            //获取批量导出文件夹
            getPlOuPath() {
                window.backend.ThisCopy.GetPlOuPath().then(result => {
                    this.sysform.plOuPath = result;
                });
            },
            //jar包名称
            getJarNames() {
                window.backend.ThisCopy.GetJarNames().then(result => {
                    if (result && result.length > 0) {
                        result.forEach(k => {
                            this.sysform.textarea += k + "\n"
                        })
                    }
                });
            },
            //提交查询信息填写列表数据
            onSubmit() {
                this.loading = true;
                window.backend.ThisCopy.GetSubmitedLogInfo(this.formInline.project, this.formInline.kssj, this.formInline.jssj, this.formInline.czr).then(result => {
                    var arry = eval(result);
                    this.tableData = [];
                    this.tableData = arry;
                    this.multipleSelection = [];
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
            Alldatas: function () {
                this.drawer_pl = true
                if (this.checks.length > 0) {
                    var k = ""
                    this.checks.forEach(s => {
                        k += s + "^";
                    })
                    window.backend.ThisCopy.GetAllProject(k, this.kssj_pl, this.jssj_pl, this.czr_pl, this.buildOrNot).then()
                } else {
                    this.showMessages("没有选中打包项目。", 'warning')
                }
            },
            //编译操作
            datas: function () {
                //window.backend.getPaths().then()
                if (this.multipleSelection.length > 0) {
                    this.$confirm('此次操作是否需要编译?', '确认信息', {
                        confirmButtonText: '编译',
                        cancelButtonText: '不编译',
                        type: 'info'
                    }).then(() => {
                        //打开弹出框
                        this.drawer = true
                        //拼接参数字符串
                        var strs = "[{}";
                        this.multipleSelection.forEach(s => {
                            strs += ',{"name":"' + s.name + '","path":"' + s.path + '"}';
                        })
                        strs += "]";
                        //调用后台方法
                        window.backend.ThisCopy.GetCom(this.formInline.project, strs, true).then()
                        // .catch(error => {
                        //     //console.log(error.message);
                        // });
                    }).catch(() => {
                        //打开弹出框
                        this.drawer = true
                        //拼接参数字符串
                        var strs = "[{}";
                        this.multipleSelection.forEach(s => {
                            strs += ',{"name":"' + s.name + '","path":"' + s.path + '"}';
                        })
                        strs += "]";
                        //调用后台方法
                        window.backend.ThisCopy.GetCom(this.formInline.project, strs, false).then()
                        // .catch(error => {
                        //     //console.log(error.message);
                        // });
                    });

                } else {
                    this.showMessages("需要选中打包的文件。", 'warning')
                }
            },
            handleClose(done) {
                this.$confirm('关闭此页面无法停止已经运行的编译操作，请慎重')
                    // eslint-disable-next-line no-unused-vars
                    .then(_ => {
                        done();
                        this.message = "";
                    })
                    // eslint-disable-next-line no-unused-vars
                    .catch(_ => {
                    });
            },
            scrollDown() {
                //滚动条处于置底
                try {
                    this.$refs['myScrollbar'].wrap.scrollTop = this.$refs['myScrollbar'].wrap.scrollHeight
                    // eslint-disable-next-line no-empty
                } catch (e) {
                }
                try {
                    this.$refs['myScrollbar_pl'].wrap.scrollTop = this.$refs['myScrollbar_pl'].wrap.scrollHeight
                    // eslint-disable-next-line no-empty
                } catch (e) {
                }
            },
            // eslint-disable-next-line no-unused-vars
            handleSelect(key, keyPath) {
                //console.log(key, keyPath);
                if (key == "1") {
                    this.tab1 = true;
                    this.tab10 = false;
                    this.tab3_1 = false;
                    this.tab3_2 = false;
                    this.tab4 = false;
                    this.tab5 = false;
                }
                if (key == "10") {
                    this.tab1 = false;
                    this.tab10 = true;
                    this.tab3_1 = false;
                    this.tab3_2 = false;
                    this.tab4 = false;
                    this.tab5 = false;
                }
                if (key == "3-1") {
                    this.tab1 = false;
                    this.tab10 = false;
                    this.tab3_1 = true;
                    this.tab3_2 = false;
                    this.tab4 = false;
                    this.tab5 = false;
                }
                if (key == "3-2") {
                    this.tab1 = false;
                    this.tab10 = false;
                    this.tab3_1 = false;
                    this.tab3_2 = true;
                    this.tab4 = false;
                    this.tab5 = false;
                }
                if (key == "4") {
                    this.tab1 = false;
                    this.tab10 = false;
                    this.tab3_1 = false;
                    this.tab3_2 = false;
                    this.tab4 = true;
                    this.tab5 = false;
                }
                if (key == "5") {

                    this.tab1 = false;
                    this.tab10 = false;
                    this.tab3_1 = false;
                    this.tab3_2 = false;
                    this.tab4 = false;
                    this.tab5 = true;
                }
            },
            formartDate(type, order) {
                var time = new Date();
                if (type != "day") {
                    if (order == "first") {
                        time = new Date(time.getTime() - 10 * 60 * 1000);
                    }
                    if (order == "end") {
                        time = new Date(time.getTime() + 10 * 60 * 1000);
                    }
                }
                var y = time.getFullYear();
                var m = time.getMonth() + 1;
                var d = time.getDate();
                var h = time.getHours();
                var mm = time.getMinutes();
                var s = time.getSeconds();
                if (type == "day") {
                    if (order == "first") {
                        return y + '-' + this.add0(m) + '-' + this.add0(d) + ' 00:00:00';
                    } else if (order == "end") {
                        return y + '-' + this.add0(m) + '-' + this.add0(d) + ' 23:59:59';
                    }
                }
                return y + '-' + this.add0(m) + '-' + this.add0(d) + ' ' + this.add0(h) + ':' + this.add0(mm) + ':' + this.add0(s);
            },
            add0(m) {
                return m < 10 ? '0' + m : m;
            },
            showMessages(msg, type) {
                this.$message({
                    message: msg,
                    type: type
                });
                // this.$message({
                //     message: '警告哦，这是一条警告消息',
                //     type: 'warning'
                // });
            },
        }
    }

</script>
<style>
    .el-carousel__item h3 {
        color: #475669;
        font-size: 14px;
        opacity: 0.75;
        line-height: 200px;
        margin: 0;
    }

    .el-carousel__item:nth-child(2n) {
        background-color: #99a9bf;
    }

    .el-carousel__item:nth-child(2n+1) {
        background-color: #d3dce6;
    }
</style>