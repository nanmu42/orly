<template>
  <div id="app">
    <TitleBar class="w100" title-msg="O'RLY Cover Generator"/>
    <Inputer class="w50" />
    <Result class="w50 gap" />
  </div>
</template>

<script>
  import TitleBar from "./components/TitleBar"
  import Inputer from "./components/Inputer"
  import Result from "./components/Result"

  export default {
    name: 'app',
    components: {
      Result,
      Inputer,
      TitleBar,
    }
  }
</script>

<style>
  body {
    background-color: #fafafa;
  }
  #app {
    box-sizing: border-box;
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    text-align: center;
    margin-top: 60px;
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
