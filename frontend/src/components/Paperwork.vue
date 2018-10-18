<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div>
    <h2>{{$t("paperwork")}}</h2>
    <p v-html="$t('paperwork_detail')"></p>
  </div>
</template>

<script>
  export default {
    name: "Paperwork"
  }
</script>

<style scoped>
  p {
    font-size: 0.9em;
    margin: auto;
    max-width: 90%;
    white-space: pre-line;
  }
</style>