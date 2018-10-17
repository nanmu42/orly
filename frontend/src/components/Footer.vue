<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div>
    <p>
      <a class="social-icon" href="https://github.com/nanmu42/orly" target="_blank">
        <img src="../assets/githubIcon.svg">
      </a>
    </p>
    <p class="bold">Released under the MIT License</p>
    <p>Copyright © 2018 LI Zhennan</p>
  </div>
</template>

<script>
  export default {
    name: "Footer"
  }
</script>

<style scoped>
  div {
    margin-top: 80px;
    padding: 30px 0 20px 0;
    width: 100%;
    background: lightseagreen;
    color: white;
  }
  p {
    text-align: center;
    white-space: pre-line;
  }
  .social-icon{
    display: inline-block;
    width: 2.4em;
  }
  .bold {
    font-weight: bold;
  }
</style>