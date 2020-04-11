<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div>
    <p>{{$t("repo_tip")}}

      <a class="social-icon" href="https://github.com/nanmu42/orly" target="_blank">
        <img src="../assets/githubIcon.svg">
      </a>
    </p>
    <p class="bold">{{$t("mit_license")}}</p>
    <p>Copyright © {{year}} LI Zhennan</p>
  </div>
</template>

<script>
  export default {
    name: "Footer",
    computed: {
      year() {
        return (new Date()).getFullYear()
      },
    },
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
  .social-icon {
    display: inline-block;
    width: 3em;
  }
  .bold {
    font-weight: bold;
  }
  a {
    color: white;
    text-decoration: underline dotted;
    transition-duration: 300ms;
    transition-timing-function: ease-in-out;
  }
  a:hover {
    transform: scale(1.5);
  }
</style>