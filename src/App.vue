<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <HelloWorld msg="Hello Vue 3.0 + Vite" />
  <br>
  <router-link to="/hello">Hello</router-link>
  |
  <router-link to="/About">About</router-link>
  <router-view></router-view>
</template>

<script>
export default {
  name: 'App',
}
</script>
