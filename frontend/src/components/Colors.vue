<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
<div class="colors">
  <h2>Color Codes</h2>
  <span v-for="(color, index) of colors" v-bind:style="{background: color}">{{index}}</span>
</div>
</template>

<script>
  export default {
    name: "Colors",
    props: {
      colors: Array
    }
  }
</script>

<style scoped>
  span {
    height: 90px;
    width: 90px;
    margin: 20px;
    border-radius: 220px;
    color: white;
    display: inline-block;
    font-size: 30px;
    line-height: 90px;
    background-color: dimgrey;
  }
  .colors {
    max-width: 85%;
    margin: 50px 0;
  }
</style>