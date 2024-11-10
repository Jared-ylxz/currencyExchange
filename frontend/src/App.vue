<template>
    <el-container>
      <el-header>
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal"  :ellipsis="true" @select="handleSelect">
          <el-menu-item index="home">首页</el-menu-item>
          <el-menu-item index="currencyExchange">兑换货币</el-menu-item>
          <el-menu-item index="news">查看新闻</el-menu-item>
          <el-menu-item index="login">登录</el-menu-item>
          <el-menu-item index="register">注册</el-menu-item>
          <el-menu-item index="logout">退出</el-menu-item>
        </el-menu>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex = ref('home')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style>
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style>