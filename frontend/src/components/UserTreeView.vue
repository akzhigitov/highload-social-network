<template>
  <v-list>
    <v-list-item
        v-for="user in users"
        :key="user.id"
        @click="$router.push(`/users/${user.id}`);"
    >
      <v-list-item-content>
        <v-list-item-title v-text="user.firstName + ' '+ user.lastName"></v-list-item-title>
      </v-list-item-content>

    </v-list-item>
  </v-list>
</template>

<script>

export default {
  props: {
    users: Array
  },
}
</script>

<style>

</style>
