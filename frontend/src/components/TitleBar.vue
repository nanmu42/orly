<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div>
    <h1>{{titleMsg}}</h1>
    <label for="lang-select">language:&nbsp;</label>
    <select id="lang-select" name="language select" v-model="langSelected">
      <option v-for="(v, k) in langSet" :value="k" :key="k">{{v}}</option>
    </select>
  </div>
</template>

<script>
  export default {
    name: "TitleBar",
    beforeMount: function () {
      this.langSelected = this.initLang
    },
    props: {
      titleMsg: String,
      initLang: String,
      langSet: Object,
    },
    data: function () {
      return {
        langSelected: "",
      }
    },
    watch: {
      "langSelected": function (newVal) {
        this.$emit("lang-changed", newVal)
      }
    }
  }
</script>

<style scoped>
  label, select {
    font-size: 0.9em;
    vertical-align: middle;
  }
  select {
    -webkit-appearance: none;
    outline: none;
    background: url('../assets/options.png') no-repeat right white;
    border: none;
    padding: 0 2em 0 1em;
  }
  div  {
    margin-bottom: 20px;
  }
</style>