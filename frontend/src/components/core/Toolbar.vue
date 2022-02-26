<template>
  <v-app-bar
      app
      dark>
    <v-toolbar-title>
      <v-app-bar-nav-icon @click="toggleNavigationBar"></v-app-bar-nav-icon>
    </v-toolbar-title>

    <v-spacer></v-spacer>

    <v-snackbar
        v-model="showResult"
        :timeout="2000"
        top>
      {{ result }}
    </v-snackbar>
    <v-menu :nudge-bottom="10" class="toolbar-menu-item" offset-y origin="center center" transition="scale-transition">
      <template v-slot:activator="{ on }">
        <v-btn :ripple="false" icon large text v-on="on">
          <v-avatar size="42px">
            <img
                src="https://avataaars.io/?avatarStyle=Circle&topType=ShortHairShortFlat&accessoriesType=Sunglasses&hairColor=Black&facialHairType=Blank&clotheType=CollarSweater&clotheColor=Black&eyeType=Default&eyebrowType=Default&mouthType=Smile&skinColor=Light"/>
          </v-avatar>
        </v-btn>
      </template>
      <v-list>
        <v-list-item
            @click="logOut">
          <v-list-item-content>
            <v-list-item-title>LogOut</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>
<script>

export default {
  data() {
    return {
      error: false,
      showResult: false,
      result: '',
    }
  },

  methods: {
    toggleNavigationBar() {
      const vm = this;
      vm.$emit('toggleNavigationBar');
    },
    async logOut() {
      try {
        await this.$store.dispatch("LogOut");
        await this.$router.push({name: 'Login'});
      } catch (e) {
        console.log(e)
      }
    }
  }
};
</script>
<style>
.toolbar-menu-item {
  padding-left: 5px;
}

</style>
