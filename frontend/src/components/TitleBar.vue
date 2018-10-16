<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <h1>{{titleMsg}}</h1>
</template>

<script>
  export default {
    name: "TitleBar",
    props: {
      titleMsg: String
    }
  }
</script>

<style scoped>
  h1 {
    display: block;
    font-size: 2em;
    margin: 0.67em 0;
    font-weight: bold;
  }
</style>