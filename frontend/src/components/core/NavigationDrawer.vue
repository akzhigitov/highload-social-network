<template>
  <v-navigation-drawer
      v-model="toggle"
      app fixed>
    <v-toolbar class="toolbar" dark flat>
      <router-link :to="{ name: 'Home' }">
        <!--        <img src="logo.png" width="36px"/>-->
      </router-link>
      <router-link :to="{ name: 'Home' }" class="text">
        HighLoad Social Network
      </router-link>
    </v-toolbar>
    <v-list>
      <v-list-item @click="changeRoute('Home', 1)">
        <v-list-item-action>
          <v-icon>home</v-icon>
        </v-list-item-action>
        <v-list-item-title :class="[{'active': selectedIndex === 1}, 'item-title' ]">Home</v-list-item-title>
      </v-list-item>

      <v-list-item @click="changeRoute('Users', 2)">
        <v-list-item-action>
          <v-icon>group</v-icon>
        </v-list-item-action>
        <v-list-item-title :class="[{'active': selectedIndex === 2}, 'item-title' ]">Users</v-list-item-title>
      </v-list-item>


    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  props: {
    toggle: {
      type: Boolean,
      required: false,
      default: true
    }
  },

  data() {
    return {
      selectedIndex: 1
    }
  },

  methods: {
    changeRoute(routeName, selectedIndex) {
      let vm = this;

      vm.selectedIndex = selectedIndex;

      return vm.$router.push({name: routeName});
    }
  }
}
</script>

<style>
.toolbar {
  font-weight: bold;
  font-size: 18px;
}

.toolbar .text {
  padding-left: 15px;
  color: white;
  text-decoration: none;
}

.item-title {
  font-size: 17px;
  font-weight: 500;
}

.item-sub-title {
  font-size: 15px;
  font-weight: 500;
}

.active {
  font-weight: bold;
}
</style>
