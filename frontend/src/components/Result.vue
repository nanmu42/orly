<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div class="result">
    <progressive-img
      class="img"
      :blur="3"
      :src="imgSrc"
      :placeholder="baseURL + 'example.gif'"
      :fallback="baseURL + 'error.gif'"
      alt="EPublication"
    ></progressive-img>
    <a :href="imgSrc" target="_blank" download>{{$t("download")}}</a>
  </div>
</template>

<script>
  export default {
    name: "Result",
    props: {
      inputSrc: String,
    },
    beforeMount: function() {
      this.imgSrc = this.inputSrc
    },
    data: function () {
      return {
        baseURL: process.env.BASE_URL,
        imgSrc: "",
      }
    },
    watch: {
      inputSrc: function (newVal) {
        if (newVal === this.imgSrc) {
          return
        }
        this.imgSrc = newVal
      }
    },
  }
</script>

<style scoped>
  .result {
    width: 800px;
  }
  @media screen and (min-width: 1024px) {
    .result {
      margin: 38px 0 0 62px;
    }
  }
  @media screen and (max-width: 1024px) {
    .result {
      margin: 38px 62px 0 54px;
    }
  }
  @media screen and (max-width: 1024px) {
    .result {
      max-width: 96%;
      margin: 38px 0 0 0;
    }
  }
  .img, img, a {
    -webkit-appearance: none;
    outline: none;
    display: block;
    width: 100%;
  }
  .img div {
    width: 100%;
  }
  a {
    white-space: nowrap;
    margin: 20px 0 0 0;
    font-size: 1.105em;
    padding: 14px;
    background: black;
    color: white;
    border: none;
    font-weight: bold;
    text-decoration: none;
  }
  a:active {
    background: white;
    color: black;
    box-shadow: 0 0 0 3px black inset;
  }
  .img {
    box-shadow: 8px 8px 30px black;
  }

</style>
