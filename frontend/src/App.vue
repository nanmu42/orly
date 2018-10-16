<template>
  <div id="app">
    <TitleBar class="w100" title-msg="O'RLY Cover Generator"/>
    <Inputer class="w50" />
    <Result class="w50 gap" />
    <Colors class="w100" v-bind:colors="colors"/>
  </div>
</template>

<script>
  import TitleBar from "./components/TitleBar"
  import Inputer from "./components/Inputer"
  import Result from "./components/Result"
  import Colors from "./components/Colors"

  export default {
    name: 'app',
    components: {
      Colors,
      Result,
      Inputer,
      TitleBar,
    },
    data: function () {
      return {
        colors: [
          "#61005e",
          "#70706d",
          "#890029",
          "#c4000e",
          "#6d001d",
          "#6a00bd",
          "#f10000",
          "#0071b1",
          "#f9bc00",
          "#2c0077",
          "#ba009a",
          "#009047",
          "#009d9e",
          "#222e85",
          "#bd002e",
          "#009d1a",
          "#75a500",
        ]
      }
    },
  }
</script>

<style>
  body {
    background-color: #fafafa;
  }
  h1 {
    display: block;
    font-size: 2em;
    margin: 2em 0 1em 0;
    font-weight: bold;
  }
  h2 {
    display: block;
    font-size: 1.6em;
    margin: 1.6em 0 1em 0;
    font-weight: bold;
  }
  #app {
    box-sizing: border-box;
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    text-align: center;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
  #app * {
    box-sizing: inherit;
    font-family: inherit;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
  .w100 {
    flex: 0 0 100%;
  }
  .w50 {
    flex: 1 0 auto;
  }
</style>
