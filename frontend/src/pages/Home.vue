<template>
  <v-container fluid>
    <Profile :user="user" @follow="follow" @unfollow="unfollow"/>
  </v-container>
</template>

<script>
import Profile from "../components/Profile";
import axios from "@/plugins/axios";

export default {
  components: {Profile},
  data() {
    return {
      user: {
        following: [],
        followers: []
      }
    }
  },

  methods: {
    async follow(userID) {
      const token = this.$store.getters.token
      const config = {
        headers: {Authorization: `Bearer ${token}`}
      };
      const bodyParameters = {};
      await axios.post(`/restricted/friends/${userID}`, bodyParameters, config)
    },
    async unfollow(userID) {
      const token = this.$store.getters.token
      const config = {
        headers: {Authorization: `Bearer ${token}`}
      };
      await axios.delete(`/restricted/friends/${userID}`, config)
    },

    async fetchUser() {
      try {
        let userID = this.$route.params.id
        if (userID === undefined) {
          userID = this.$store.getters.userID
        }
        const response = await axios.get(`/users/${userID}`)
        this.user = response.data
      } catch (e) {
        console.log("Ошибка")
      }
    },
  },
  created() {
    this.fetchUser()
  },
  async beforeRouteUpdate(from, to, next) {
    console.log(to)
    const response = await axios.get(`/users/${to.params.id}`)
    this.user = response.data
    next()
  },

  watch: {
    '$route.params.id': function (id) {
      console.log(id)
      this.fetchUser()
    }
  }
}

</script>

<style>

</style>
