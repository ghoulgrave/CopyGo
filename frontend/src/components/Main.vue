<template>
    <el-tabs v-model="activeName">
        <el-tab-pane label="获取更新" name="getUpdates">
        <LogsTable/>
        </el-tab-pane>
        <el-tab-pane label="配置项目信息" name="setProjectConfig">配置项目信息</el-tab-pane>
        <el-tab-pane label="设置系统信息" name="setSystemConfig">设置系统信息</el-tab-pane>
    </el-tabs>

</template>

<script>


    export default {
        name: "Main",
        data() {
            return {
                activeName: 'getUpdates'
            }
        },
        components: {
            //HelloWorld,
            LogsTable
        },
        methods: {}
    }
</script>

<style scoped>

</style>