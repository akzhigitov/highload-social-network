<template>
  <div>
    <template v-if="isLoggedIn">
      <v-app id="inspire">
        <div class="app-container">
          <toolbar @toggleNavigationBar="drawer = !drawer"/>
          <navigation :toggle="drawer"/>
          <v-main>
            <router-view/>
          </v-main>
        </div>
      </v-app>
    </template>
    <template v-else>
      <router-view/>
    </template>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      drawer: true
    }
  },
  computed: {
    isLoggedIn: function () {
      return this.$store.getters.isAuthenticated
    }
  },
}
</script>

<style>
.v-btn:hover:before {
  color: transparent !important;
}
</style>
